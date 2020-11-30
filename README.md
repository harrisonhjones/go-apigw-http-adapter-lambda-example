# go-apigw-http-adapter-lambda-example

Example AWS Lambda function utilizing the
[harrisonhjones.com/go-apigw-http-adapter](https://github.com/harrisonhjones/go-apigw-http-adapter)
package.

## Badges

- Build:
  ![Build](https://github.com/harrisonhjones/go-apigw-http-adapter-lambda-example/workflows/Go/badge.svg)
- Report Card:
  [![Go Report Card](https://goreportcard.com/badge/harrisonhjones.com/go-apigw-http-adapter-lambda-example)](https://goreportcard.com/report/harrisonhjones.com/go-apigw-http-adapter-lambda-example)

## Links

- [pkg.go.dev (documentation)](https://pkg.go.dev/harrisonhjones.com/go-apigw-http-adapter-lambda-example)
- [Demo](https://lzqz36hhb4.execute-api.us-east-1.amazonaws.com)

## Testing

### Installing

1. Git clone this project.
1. Run `go build ./...`. This should create a
   `go-apigw-http-adapter-lambda-example` binary.

### Running Locally

1. Run `./go-apigw-http-adapter-lambda-example`.
   1. Open http://localhost:8080 in your brower.
   1. You should see a web page served, an image, and a favicon (see the demo
      above).

### Deploying to Lambda and API Gateway

1. Run
   `zip go-apigw-http-adapter-lambda-example go-apigw-http-adapter-lambda-example`.
   This should create a `go-apigw-http-adapter-lambda-example.zip` file.
1. Create a new AWS Lambda function using the AWS Lambda console.
   1. Make sure the `Go 1.x` runtime is chosen.
1. Upload the `go-apigw-http-adapter-lambda-example.zip` file to the Lambda
   function.
1. Update the handler to be `go-apigw-http-adapter-lambda-example`.
1. Create a new **HTTP** API in the AWS API Gateway console.
1. Set the integration to the Lambda function created earlier.
   1. Make sure the version is `2.0`.
1. Create a route with path `$default`.
   1. This will set the "method" to `ANY`.
1. Leave all the other options as the defaults.
1. Create the API and test it.
