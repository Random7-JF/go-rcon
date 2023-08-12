package model

type Auth struct {
	Status  bool
	Message string
	Admin   bool
}
type Flash struct {
	Info    string
	Warning string
	Error   string
}
