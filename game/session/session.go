package session

import (
	"context"

	"../../proto"
)

// Session is the main structure that contains session data
type Session struct {
	ID       string
	Username string
}

// NewSession sort-of constructor
func NewSession(id string, username string) *Session {
	return &Session{ID: id, Username: username}
}

// LookAround just looks around
func LookAround(context.Context, *proto.LookAroundRequest) (*proto.LookAroundAnswer, error) {
	answer := proto.LookAroundAnswer{}
	return &answer, nil
}
