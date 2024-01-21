# CDK Project to deploy Go Lambda

- `npm install` install all dependencies
- `go get github.com/aws/aws-lambda-go` cd into lambda pull down dependencies
- `GOOS=linux GOARCH=amd64 go build main.go` build go code using the linux parameters
- `cdk bootstrap` cd to root bootstrap AWS credentials
- `cdk synth` build a cdk.out file
- `cdk deploy` deploy code to AWS
- test endpoint -

## GO Build

- `GOOS=linux GOARCH=amd64 go build main.go` build your Go code for the Linux environment
