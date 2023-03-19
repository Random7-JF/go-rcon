package model

type NoReplyCommand struct {
	Error string `json:"error"`
}
type CommandResponse struct {
	Response string `json:"response"`
	Error    string `json:"error"`
}

type KickCommand struct {
	Target   string `json:"target"`
	Response string `json:"message"`
	Error    string `json:"error"`
}

type PlayersCommand struct {
	CurrentCount int       `json:"currentcount"`
	MaxCount     int       `json:"maxcount"`
	Players      []Players `json:"players"`
}

type WhitelistCommand struct {
	Count   int       `json:"count"`
	Players []Players `json:"players"`
}

type Players struct {
	Name string `json:"name"`
}

type TeleportCommand struct {
	Xcoord   string `json:"xcoord"`
	Ycoord   string `json:"ycoord"`
	Target   string `json:"target"`
	Response string `json:"response"`
}

type TempalteData struct {
	Title string
	Data  map[string]interface{}
}
