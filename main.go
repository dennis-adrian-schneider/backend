package main

import (
	"sv_backend/Database"
	"sv_backend/Webservice"
)

func main() {
	Database.Init()
	Webservice.Init()
}
