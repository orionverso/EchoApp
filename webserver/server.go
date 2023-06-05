package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"server/receiver"

	"github.com/aws/aws-sdk-go-v2/config"
)

var clientReceiver receiver.Receiver
var ctx context.Context

func main() {
	http.HandleFunc("/", postHandler)
	http.ListenAndServe(":80", nil)
}

func postHandler(w http.ResponseWriter, r *http.Request) {
	log.Println(r.RemoteAddr)

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error Body Read", http.StatusInternalServerError)
		return
	}
	if string(body) != "" {

		clientReceiver.Write(ctx, string(body))
	}

	fmt.Fprintf(w, "Thanks you for testing. I am from a Container. See you")
}

func init() {
	ctx = context.TODO()
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		log.Fatal(err)
	}
	clientReceiver, err = receiver.GetReceiver(ctx, cfg)
}
