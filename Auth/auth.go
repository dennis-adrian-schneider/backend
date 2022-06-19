package Auth

import (
	"gorm.io/gorm"
	"log"
	"sv_backend/Database"
)

type Auth struct {
	gorm.Model
	ID        string
	authLevel AuthLevel
}
type AuthLevel string

const (
	LOW    = "LOW"
	MEDIUM = "MEDIUM"
	HIGH   = "HIGH"
)

func CheckID(requestID string) AuthLevel {
	var level AuthLevel
	Database.DB.Select("authLevel").Where(&level, "ID = ?", requestID)
	log.Printf("AuthLevel of %s is %s\n", requestID, level)
	return level
}
