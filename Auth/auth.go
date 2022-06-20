package Auth

import (
	"gorm.io/gorm"
	"log"
	"sv_backend/Database"
)

type Auth struct {
	gorm.Model
	ID        string
	AuthLevel AuthLevel
}
type AuthLevel string

const (
	LOW    = "LOW"
	MEDIUM = "MEDIUM"
	HIGH   = "HIGH"
)

func CheckID(requestID string) string {
	var authLevel string
	Database.DB.Raw("select auth_level from auths where id = ?", requestID).Scan(&authLevel)
	log.Printf("AuthLevel of %s is %s\n", requestID, authLevel)

	return authLevel
}
func AddID(requestID string, authLevel string) {
	Database.DB.AutoMigrate(&Auth{})
	var auth Auth
	auth.ID = requestID
	auth.AuthLevel = AuthLevel(authLevel)

	Database.DB.Create(&auth)

	log.Printf("Created User %s with AuthLevel %s\n", requestID, authLevel)
}
