package model

type Auth struct {
	Status  bool
	Message string
}

type Flash struct {
	Info    string
	Warning string
	Error   string
}
