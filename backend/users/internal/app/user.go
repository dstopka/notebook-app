package app

// User is the representation of an app user.
type User struct {
	UUID string

	Name      string
	Role      string
	AvatarURL string
	LastIP    string
}