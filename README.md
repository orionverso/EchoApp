# Welcome to your CDK Go project!

This is a blank project for CDK development with Go.

The `cdk.json` file tells the CDK toolkit how to execute your app.

## Useful commands

- `cdk deploy` deploy this stack to your default AWS account/region
- `cdk diff` compare deployed stack with current state
- `cdk synth` emits the synthesized CloudFormation template
- `go test` run unit tests

## Compile and deploy go lambda function

GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o main main.go

zip main.zip main

rm -f main
