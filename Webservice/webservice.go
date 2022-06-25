package Webservice

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"regexp"
	"sv_backend/Auth"
)

func Init() {
	checkID("/check-id/")

	if len(os.Args) > 1 {
		_, err := regexp.MatchString("^\\d*$", os.Args[1])
		if err != nil {
			log.Println("Line Argument One failed. Check Port Settings")
			panic("Line Argument One failed. Check Port Settings")
		}
		http.ListenAndServe(":"+os.Args[1], nil)
	} else {
		http.ListenAndServe(":4500", nil)
		log.Println("Webservice started on port 4500")
	}

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
			fmt.Fprintf(writer, "ID %s invalid", id)
		}
	})
}
