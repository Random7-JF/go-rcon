package model

import (
	"gorm.io/gorm"
)

type Users struct {
	gorm.Model
	UserName string
	Password string
	Admin    bool
}

type CommandLog struct {
	gorm.Model
	CommandType string
	Command     string
	SentBy      string
	Response    string
}

type ServerSettings struct {
	gorm.Model
	SpawnCoords string
}
