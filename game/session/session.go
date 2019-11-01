package session

// Session is the main structure that contains session data
type Session struct {
	ID       string
	Username string
}

// NewSession sort-of constructor
func NewSession(id string, username string) *Session {
	return &Session{ID: id, Username: username}
}
