package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	goautobahn "github.com/julez-dev/go-autobahn"
)

func main() {
	// Create a custom http client with a timeout
	http := &http.Client{
		Timeout: time.Second * 15,
	}

	// Create a new api instance
	autobahn := goautobahn.New(goautobahn.WithHTTPClient(http))

	// Get all roadworks for the A1
	resp, err := autobahn.Roadworks.GetAllForRoad(context.TODO(), "A1")

	if err != nil {
		log.Fatalln(err)
	}

	// Print Roadworks as JSON
	respDbg, _ := json.MarshalIndent(resp, "", "	")
	fmt.Println(string(respDbg))
}
