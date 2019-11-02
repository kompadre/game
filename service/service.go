package service

import (
	"context"
	"fmt"
	"log"
	"time"

	gsess "../game/session"
	"../proto"
	"github.com/google/uuid"
)

var activeSessions map[string]*gsess.Session
var activeConnections map[string]*proto.Session_LookAroundServer

func init() {
	activeSessions = make(map[string]*gsess.Session)
	activeConnections = make(map[string]*proto.Session_LookAroundServer)
}

// SessionImpl is the transport structure
type SessionImpl struct {
	uuid string
}

// NewSession grants a new session
func (c *SessionImpl) NewSession(ctx context.Context, req *proto.SessionRequest) (*proto.SessionGrant, error) {
	answer := proto.SessionGrant{Uuid: "", Reason: ""}
	if req.GetUsername() != "kompadre" || req.GetPassword() != "Unlikely" {
		answer.Reason = "Unauthorized"
		return &answer, nil
	}

	uuid := uuid.New()
	newclient := gsess.NewSession(uuid.String(), "kompadre")

	if err := registerNewSession(newclient); err != nil {
		log.Fatalf("something went wrong registering a new session %v", newclient)
		return nil, err
	}
	answer.Uuid = uuid.String()
	c.uuid = uuid.String()
	return &answer, nil
}

// LookAround returns a new Session
func (s *SessionImpl) LookAround(stream proto.Session_LookAroundServer) error {
	done := make(chan bool)
	ctx := stream.Context()
	go func() {
		for {
			time.Sleep(1 * time.Second)
			objects := make([]*proto.Object, len(activeSessions))
			i := 0
			for key := range activeSessions {
				object := proto.Object{Uuid: key, X: 1, Y: 1}
				objects[i] = &object
				i++
			}
			answer := proto.LookAroundAnswer{Results: objects}
			if err := stream.Send(&answer); err != nil {
				fmt.Printf("Error sending: %v\n", err)
				done <- true
			}
		}
	}()

	for {
		select {
		case <-done:
			unregisterNewSession(s.uuid)
			return nil
		case <-ctx.Done():
			unregisterNewSession(s.uuid)
			return ctx.Err()
		default:
		}
		time.Sleep(time.Second / 60)
	}
}

func registerNewSession(sess *gsess.Session) error {
	fmt.Println("***\nServer speaking:\n***")
	fmt.Printf("New active session: %v\n", sess.ID)
	fmt.Println("Old active sessions:")
	for key := range activeSessions {
		fmt.Printf("\t%s\n", key)
	}
	activeSessions[sess.ID] = sess
	return nil
}

func unregisterNewSession(ID string) {
	fmt.Printf("Unregistering %s\n", ID)
	delete(activeConnections, ID)
	delete(activeSessions, ID)
}

// compile-type check that our new type provides the
// correct server interface
var _ proto.SessionServer = (*SessionImpl)(nil)
