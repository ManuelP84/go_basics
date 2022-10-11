package user

type User interface {
	Permissions() int
	GetName() string
}
