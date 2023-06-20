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

var rec receiver.Receiver
var ctx context.Context

func main() {
	http.HandleFunc("/", postHandler)
	http.ListenAndServe(":80", nil)
	http.ListenAndServe(":443", nil)
}

func postHandler(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error Body Read", http.StatusInternalServerError)
		return
	}

	if err != nil {
		log.Println(err)
	}

	if string(body) != "" {
		fmt.Println("la string a enviar es:", string(body))
		rec.Write(ctx, string(body))
	}

	fmt.Fprintf(w, "Thanks you for take a look. I am from a Container. See you")
}

func init() {

	ctx = context.TODO()
	cfg, err := config.LoadDefaultConfig(ctx)

	if err != nil {
		// Handle error
	}
	rec, err = receiver.GetReceiver(ctx, cfg)

}

type LoadDefaultConfigProps struct {
	profile config.LoadOptionsFunc
	region  config.LoadOptionsFunc
}

// CONFIGURATION
var LoadDefaultConfigProps_IMAGE_IN_LOCAL LoadDefaultConfigProps = LoadDefaultConfigProps{
	profile: config.WithSharedConfigProfile("cdk-role"),
	region:  config.WithRegion("us-east-1"),
}

var LoadDefaultConfigProps_BINARY_IN_LOCAL LoadDefaultConfigProps = LoadDefaultConfigProps{
	profile: config.WithSharedConfigProfile("workerdev"),
	region:  config.WithRegion("us-east-1"),
}

var LoadDefaultConfigProps_IN_CLOUD LoadDefaultConfigProps = LoadDefaultConfigProps{ //Asumme Credentials from Task Role
	profile: nil,
	region:  nil,
}
