# Echoing app with Go

This project is an example about create functionalities with aws-cdk for Go. It
designed to be loosely coupled between components. I think this is the powerful
philosophy about cdk constructs.

## There are two actors

- Writer
- Storage

## Writer

- RestApiLambda
- FargateContainer

## Storage

- S3 bucket
- DynamoDB

## Workflow

![Alt text](/images/lambdas3ch.png "Lambda-S3")
![Alt text](/images/lambdadbch.png "Lambda-DynamoDB")
![Alt text](/images/fargates3ch.png "Fargate-S3")
![Alt text](/images/fargatedbch.png "Fargate-DynamoDB")

## Prerequisites

Install aws-cdk

```bash
npm install -g aws-cdk
cdk --version
```

Put a role in cdk.json with the necessary permissions

```bash
{
  "app": "go mod download && go run writer_storage_app.go",
  "profile": "<INTRODUCE-YOUR-ROLE>",
  "watch": {
    "include": ["**"],
    ...
```

## Take a look

Clone the project

```bash
 git clone https://github.com/orionverso/EchoApp
```

Go to the project directory

```bash
  cd EchoApp
```

Choose one stack in the main function

```go
func main() {
	defer jsii.Close()

	app := awscdk.NewApp(nil)
	//This is the model to deploy other stacks
	NewWriterStorageAppStackApiLambdaDB(app, "WriterStorageAppStack-Lambda-DB-", &WriterStorageAppStackProps{

		awscdk.StackProps{
			Env: env(),
		},
	})

	app.Synth(nil)
}
```

Deploy the stack

```bash
  cdk deploy
```

Push docker webserver image to ecr repo when is created (only in Fargate
workflows) ![Alt text](/images/ecr-push.png "push commands")

Expected output

```bash
   ✅  WriterStorageAppStack-Lambda-DB-

 ✨  Deployment time: 82.06s

 Outputs:
 WriterStorageAppStack-Lambda-DB-.LambdaApiWriterEndpointWriterEndpointA2088D8E =
 https://<API-ID>.execute-api.<REGION>.amazonaws.com/test/
 Stack ARN:
 arn:aws:cloudformation:<REGION>:<ACCOUNT-ID>:stack/WriterStorageAppStack-Lambda-DB-/b9ef8b30-04a6-11ee-a6a1-0ea4e49dd5fb

 ✨  Total time: 89.7s

```

Check funcionality

```bash
curl https://<API-ID>.execute-api.<REGION>.amazonaws.com/test/ \
-X POST \
-d "Hello Storage..."
```

![Alt text](/images/echo-db.png)

```bash
cdk destroy
#In S3 and fargate deployment, you must delete bucket and ecr repository manually.
```

## Conclusion

This project was a way to show me how the cdk constructs model can be a powerful
way to build in the cloud. Two simple reasons:

- The modular nature. For example, with a more different pairing of Writer and
  Storage, it increases up to nine possible stacks.
- The ability to compose. For example, these nine possible stacks may belong to
  high-level component.

Without forgetting that DevOps practices become implementable as well.

## FAQ

#### CDK guided

https://docs.aws.amazon.com/cdk/v2/guide/home.html

#### Questions about roles

https://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles.html
