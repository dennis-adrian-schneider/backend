package Webservice

import (
	"fmt"
	"log"
	"net/http"
	"sv_backend/Auth"
)

func Init() {
	checkID("/check-id/")
	http.ListenAndServe("localhost:8080", nil)

}

func checkID(path string) {
	http.HandleFunc(path, func(writer http.ResponseWriter, request *http.Request) {
		id := request.RequestURI[len(path):]
		log.Println("Requested auth for ID " + id)
		switch Auth.CheckID(id) {
		case Auth.LOW:
			fmt.Fprintf(writer, Auth.LOW)
			log.Printf("Sent response for id %s, auth level %s\n", id, Auth.LOW)
		case Auth.MEDIUM:
			fmt.Fprintf(writer, Auth.MEDIUM)
			log.Printf("Sent response for id %s, auth level %s\n", id, Auth.MEDIUM)
		case Auth.HIGH:
			fmt.Fprintf(writer, Auth.HIGH)
			log.Printf("Sent response for id %s, auth level %s\n", id, Auth.HIGH)
		default:
			fmt.Fprintf(writer, "ID %s not valid", id)
		}
	})
}
