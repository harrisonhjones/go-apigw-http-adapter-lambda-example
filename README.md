# go-apigw-http-adapter-lambda-example

Example AWS Lambda function utilizing the
[harrisonhjones.com/go-apigw-http-adapter](https://github.com/harrisonhjones/go-apigw-http-adapter)
package for both HTTP and REST APIs.

## Badges

- Build:
  ![Build](https://github.com/harrisonhjones/go-apigw-http-adapter-lambda-example/workflows/Go/badge.svg)
- Report Card:
  [![Go Report Card](https://goreportcard.com/badge/harrisonhjones.com/go-apigw-http-adapter-lambda-example)](https://goreportcard.com/report/harrisonhjones.com/go-apigw-http-adapter-lambda-example)

## Links

- [pkg.go.dev (documentation)](https://pkg.go.dev/harrisonhjones.com/go-apigw-http-adapter-lambda-example)
- [HTTP API Demo](https://nhovle879a.execute-api.us-east-1.amazonaws.com)
- [REST API Demo](https://qcgmmbbci5.execute-api.us-east-1.amazonaws.com/test/home)

## Contributing

1. Make changes.
1. Add / update tests.
1. Run `make` to fmt, vet, test, and build your changes.
1. Follow the **Testing** steps below to deploy your changed binary to Lambda
   and test it.
1. Commit your changes.
1. Submit a PR.

## Testing

### Building

1. Git clone this project.
1. Run `make`. Two binaries and zip files will be created in the `bin`
   directory.

### Running Locally

#### HTTP API

1. Run `./bin/httpapi` to
   1. Open http://localhost:8080 in your browser.
   1. You should see a web page served, an image, and a favicon (see the demo
      above).

#### REST API

1. Run `./bin/restapi` to
1. Open http://localhost:8080/home in your browser.
1. You should see a web page served, an image, and a favicon (see the demo
   above).

### Deploying to Lambda and API Gateway

#### HTTP API

1. Create a new AWS Lambda function using the AWS Lambda console.
   1. Make sure the `Go 1.x` runtime is chosen.
1. Upload the `httpapi.zip` file to the Lambda function.
1. Update the handler to be `httpapi`.
1. Create a new **HTTP** API in the AWS API Gateway console.
1. Set the integration to the Lambda function created earlier.
   1. Make sure the version is `2.0`.
1. Create a route with path `$default`.
   1. This will set the "method" to `ANY`.
1. Leave all the other options as the defaults.
1. Create the API and test it by going to `INVOKEURL`.

#### REST API

1. Create a new AWS Lambda function using the AWS Lambda console.
   1. Make sure the `Go 1.x` runtime is chosen.
1. Upload the `restapi.zip` file to the Lambda function.
1. Update the handler to be `restapi`.
1. Create a new **REST** API in the AWS API Gateway console.
1. Create a new method set to use the Lambda function created earlier.
   1. Make sure it's setup to be a proxy integration.
1. Set the `Binary Media Types` to `image/*`.
1. Deploy the method using any stage name.
1. Leave all the other options as the defaults.
1. Create the API and test it by going to `INVOKEURL/home`.
