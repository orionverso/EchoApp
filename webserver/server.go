package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go/aws"
)

type receiver interface {
	Write(string) error
}

type s3receiver struct {
	client *s3.Client
}

func (s3rv s3receiver) Write(st string) error {
	body := []byte(st)
	_, err := s3rv.client.PutObject(ctx, &s3.PutObjectInput{
		ContentType: aws.String("application/json"),
		Bucket:      aws.String(os.Getenv("DESTINATION")),
		Key:         aws.String("Example.json"),
		Body:        bytes.NewReader(body),
	})
	if err != nil {

		log.Println(err)
		return err
	}
	return nil
}

type dynamodreceiver struct {
	client *dynamodb.Client
}

func (dbrv dynamodreceiver) Write(st string) error {
	_, err := dbrv.client.PutItem(ctx, &dynamodb.PutItemInput{
		TableName: aws.String(os.Getenv("DESTINATION")),
		Item: map[string]types.AttributeValue{
			"id": &types.AttributeValueMemberS{Value: st},
		},
	})
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

var rv receiver

func main() {
	http.HandleFunc("/", postHandler)
	http.ListenAndServe(":80", nil)
}

func postHandler(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error Body Read", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Thanks you for testing. See you from a Container \n Contenido POST: %s", body)
}
