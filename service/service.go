package service

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	gsess "../game/session"

	"../proto"
	"github.com/google/uuid"
	metadata "google.golang.org/grpc/metadata"
)

var (
	activeSessions         = map[string]gsess.Session{}
	activeConnections      = map[string]proto.Session_LookAroundServer{}
	activeSessionsMutex    = sync.RWMutex{}
	activeConnectionsMutex = sync.RWMutex{}
	sUUIDMutex             = sync.RWMutex{}
)

func init() {
	//	activeSessions = make(map[string]gsess.Session)
	//	activeConnections = make(map[string]proto.Session_LookAroundServer)
}

// SessionImpl is the transport structure
type SessionImpl struct {
	uuid string
}

// NewSession grants a new session
func (c *SessionImpl) NewSession(ctx context.Context, req *proto.SessionRequest) (*proto.SessionGrant, error) {
	fmt.Printf("NewSession context: %v\n", ctx)
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
	meta, _ := metadata.FromIncomingContext(ctx)
	fmt.Printf("CtxMeta: %v\n", meta)
	answer.Uuid = uuid.String()
	sUUIDMutex.Lock()
	c.uuid = uuid.String()
	sUUIDMutex.Unlock()
	return &answer, nil
}

func sendChat(uuidto string, from string, message string) {
	activeConnectionsMutex.RLock()
	stream, ok := activeConnections[uuidto]
	if !ok {
		return
	}
	activeConnectionsMutex.RUnlock()
	attribute := proto.ObjectAttribute{Name: from, Currentvalue: message}
	attributes := []*proto.ObjectAttribute{&attribute}
	object := proto.Object{Attributes: attributes}
	objects := make([]*proto.Object, 1)
	objects[0] = &object
	packet := proto.LookAroundAnswer{Type: proto.Types_CHAT, Results: objects}
	stream.(proto.Session_LookAroundServer).Send(&packet)
}

func sendLookArounds(stream proto.Session_LookAroundServer, endstream chan bool, uuidchan chan string) {
	ID := <-uuidchan
	fmt.Printf("ID received: %v\n", ID)
	for {
		select {
		case <-endstream:
			return
		default:
		}
		activeSessionsMutex.RLock()
		if len(activeSessions) == 0 {
			activeSessionsMutex.RUnlock()
			continue
		}

		var objects []*proto.Object
		for key := range activeSessions {
			object := proto.Object{Uuid: key, X: 1, Y: 1}
			objects = append(objects, &object)
		}

		activeSessionsMutex.RUnlock()

		answer := proto.LookAroundAnswer{Results: objects}
		if err := stream.Send(&answer); err != nil {
			fmt.Printf("Error sending: %v\n", err)
			return
		}
		time.Sleep(1 * time.Second)
	}
}

func receiveLookArounds(stream proto.Session_LookAroundServer, endstream chan bool, uuidchan chan string) {
	ID := ""
	for {
		resp, err := stream.Recv()
		fmt.Printf("Client received%v: %v", resp, err)
		if err != nil {
			if ID != "" {
				unregisterNewSession(ID)
			}
			close(endstream)
			return
		}
		if ID == "" {
			ID = resp.Myuuid
			if _, ok := activeSessions[ID]; !ok {
				unregisterNewSession(ID)
				close(endstream)
				return
			}
			activeConnectionsMutex.Lock()
			activeConnections[ID] = stream
			activeConnectionsMutex.Unlock()
			uuidchan <- ID
		}
	}
}

// LookAround returns a new Session
func (s *SessionImpl) LookAround(stream proto.Session_LookAroundServer) error {
	ctx := stream.Context()
	endstream := make(chan bool)
	uuidchan := make(chan string)
	go sendLookArounds(stream, endstream, uuidchan)
	go receiveLookArounds(stream, endstream, uuidchan)

	go func() {
		for {
			time.Sleep(5 * time.Second)
			if len(activeConnections) < 1 {
				fmt.Printf("Not enough active connections.\n")
				continue
			}
			for key := range activeConnections {
				fmt.Printf("Sending boring to %v\n", key)
				sendChat(key, "Unkonwn", "Boring!")
			}
			select {
			case <-endstream:
				return
			default:
			}
		}
	}()

	for {
		select {
		case <-ctx.Done():
		case <-endstream:
			return ctx.Err()
		default:
		}
	}
}

func registerNewSession(sess *gsess.Session) error {
	activeSessionsMutex.Lock()
	fmt.Println("***\nServer speaking:\n***")
	fmt.Printf("New active session: %v\n", sess.ID)
	fmt.Println("Old active sessions:")
	for key, val := range activeSessions {
		fmt.Printf("\t%v: %v\n", key, val)
	}
	activeSessions[sess.ID] = *sess
	activeSessionsMutex.Unlock()
	return nil
}

func unregisterNewSession(ID string) {
	activeSessionsMutex.Lock()
	fmt.Printf("Unregistering %s\n", ID)
	if _, ok := activeSessions[ID]; ok {
		delete(activeSessions, ID)
	}
	fmt.Println("Remaining active sessions:")
	for key, val := range activeSessions {
		fmt.Printf("\t%v: %v\n", key, val)
	}
	activeSessionsMutex.Unlock()

	activeConnectionsMutex.Lock()
	if _, ok := activeConnections[ID]; ok {
		delete(activeConnections, ID)
	}
	fmt.Println("Remaining active connections:")
	for key, val := range activeConnections {
		fmt.Printf("\t%v: %v\n", key, val)
	}
	activeConnectionsMutex.Unlock()
}

// compile-type check that our new type provides the
// correct server interface
var _ proto.SessionServer = (*SessionImpl)(nil)
