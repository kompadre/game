package service

import (
	"fmt"

	gsess "../game/session"
	"../proto"
	"github.com/google/uuid"
	"golang.org/x/net/context"
)

var activeSessions map[string]*gsess.Session

func init() {
	activeSessions = make(map[string]*gsess.Session)
}

// SessionImpl is the transport structure
type SessionImpl struct{}

// NewSession returns a new Session
func (s *SessionImpl) NewSession(ctx context.Context, req *proto.SessionRequest) (*proto.SessionGrant, error) {
	uuid := uuid.New()
	newclient := gsess.NewSession(uuid.String(), "kompadre")
	err := registerNewSession(newclient)
	if err != nil {
		return nil, err
	}
	return &proto.SessionGrant{Uuid: uuid.String()}, nil
}

func registerNewSession(sess *gsess.Session) error {
	activeSessions[sess.ID] = sess
	fmt.Print("Active sessions: \n")
	for key := range activeSessions {
		fmt.Printf("\t%s\n", key)
	}
	return nil
}

// compile-type check that our new type provides the
// correct server interface
var _ proto.SessionServer = (*SessionImpl)(nil)
