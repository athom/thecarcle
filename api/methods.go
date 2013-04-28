package thecacleapi

type PublicScope interface {
	GetSession(email string, password string) (session string, err error)
	GetAuthenticatedScope(session string) (scope *AuthenticatedScope)
	Comment(name string, email string, content string) (err error)
}

type AuthenticatedScope interface {
	Post(title string, content string) (err error)
}
