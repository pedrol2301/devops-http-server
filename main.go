package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	auth "github.com/abbot/go-http-auth"
)

func Secret(user, reaml string) string {
	if user == "john" {
		// password is "hello"
		return "$1$dlPL2MqE$oQmn16q49SqdmhenQuNgs1"
	}
	return ""
}

func main() {

	if len(os.Args) != 3 {
		fmt.Println("Usage: go run main.go <dir> <port>")
		os.Exit(1)

	}

	httpDir := os.Args[1]
	port := os.Args[2]

	authenticator := auth.NewBasicAuthenticator("example.com", Secret)

	http.Handle("/", authenticator.Wrap(func(w http.ResponseWriter, ar *auth.AuthenticatedRequest) {
		http.FileServer(http.Dir(httpDir)).ServeHTTP(w, &ar.Request)
	}))

	fmt.Println("Serving files from: " + httpDir + " on port: " + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
