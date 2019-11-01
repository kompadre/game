package service

import (
	"context"
	"fmt"
	"io"
	"log"

	gsess "../game/session"
	"../proto"
	"github.com/google/uuid"
)

var activeSessions map[string]*gsess.Session
var activeConnections map[string]*proto.Session_NewSessionServer

func init() {
	activeSessions = make(map[string]*gsess.Session)
	activeConnections = make(map[string]*proto.Session_NewSessionServer)
}

// SessionImpl is the transport structure
type SessionImpl struct{}

// LookAround looks around
func (c *SessionImpl) LookAround(ctx context.Context, in *proto.LookAroundRequest) (*proto.LookAroundAnswer, error) {
	results := make([]*proto.Object, 3)
	for i := 3; i > 0; i-- {
		uuid := uuid.New()
		object := proto.Object{X: 0, Y: 0, Uuid: uuid.String()}
		results[i] = &object
	}
	answer := proto.LookAroundAnswer{Results: results}
	return &answer, nil
}

// NewSession returns a new Session
func (s *SessionImpl) NewSession(stream proto.Session_NewSessionServer) error {
	ctx := stream.Context()
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}
		req, err := stream.Recv()
		/*
			if req.Password == "" || req.Username == "" {
				fmt.Printf("Username or Password were null, continuing.\n")
				continue
			}
		*/
		if err == io.EOF {
			log.Printf("exit")
			return nil
		}
		if err != nil {
			log.Printf("receive error %v", err)
			return nil
		}

		if req.Username != "kompadre" || req.Password != "Unlikely" {
			resp := proto.SessionGrant{Reason: "Incorrect Credentials"}
			stream.Send(&resp)
			return nil
		}

		uuid := uuid.New()
		newclient := gsess.NewSession(uuid.String(), "kompadre")
		resp := proto.SessionGrant{Uuid: uuid.String()}
		err = registerNewSession(newclient, &stream)
		defer unregisterNewSession(uuid.String())
		if err != nil {
			return err
		}
		if stream.Send(&resp); err != nil {
			log.Printf("send error:")
		}
	}
}

func registerNewSession(sess *gsess.Session, stream *proto.Session_NewSessionServer) error {
	activeSessions[sess.ID] = sess
	activeConnections[sess.ID] = stream
	fmt.Print("Active sessions: \n")
	for key := range activeSessions {
		fmt.Printf("\t%s\n", key)
	}
	return nil
}

func unregisterNewSession(ID string) {
	fmt.Printf("Unregistering %s", ID)
	delete(activeConnections, ID)
	delete(activeSessions, ID)
}

// compile-type check that our new type provides the
// correct server interface
var _ proto.SessionServer = (*SessionImpl)(nil)
