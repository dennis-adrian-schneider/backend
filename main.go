package main

import (
	"sv_backend/Database"
	"sv_backend/Webservice"
)

func main() {
	sub := "/check-id/"
	Database.Init()
	Webservice.Run(sub)
}
