package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"

	"github.com/aws/aws-lambda-go/lambda"
	example "harrisonhjones.com/go-apigw-http-adapter-lambda-example"
	"harrisonhjones.com/go-apigw-http-adapter/restadapter"
)

func main() {
	mux := http.NewServeMux()
	mux.Handle("/home", http.HandlerFunc(IndexHandler))
	mux.Handle("/favicon.svg", http.HandlerFunc(FaviconHandler))
	mux.Handle("/project.png", http.HandlerFunc(ImageHandler))

	if os.Getenv("AWS_LAMBDA_FUNCTION_NAME") != "" {
		lambda.Start(func(ctx context.Context, req restadapter.Request) (*restadapter.Response, error) {
			log.Printf("Request: %#v", req)

			// FYI: Request transformation.
			httpReq, err := restadapter.TransformRequest(ctx, req)
			if err != nil {
				return nil, err
			}

			// FYI: Handle Request.
			httpRec := httptest.NewRecorder()
			mux.ServeHTTP(httpRec, httpReq)

			// FYI: Response transformation.
			httpRes, err := restadapter.TransformResponse(httpRec.Result(), func(response *http.Response) bool {
				if ct := response.Header.Get("content-type"); strings.HasPrefix(ct, "image") {
					log.Printf("content type %q treated as image; encoding...", ct)
					return true
				}
				return false
			})
			if err != nil {
				return nil, err
			}

			log.Printf("Response: %#v", httpRes)
			return httpRes, nil
		})
	} else {
		addr := ":8080"
		log.Printf("Listening on %q", addr)
		if err := http.ListenAndServe(addr, mux); err != nil {
			log.Fatalf("failed to listen: %v", err)
		}
	}
}

func IndexHandler(writer http.ResponseWriter, _ *http.Request) {
	fmt.Fprint(writer, `<!DOCTYPE html>
<html>
  <head>
    <title>Hello World</title>
    <link rel="icon" type="image/svg+xml" href="favicon.svg" />
  </head>

  <body>
    <h1>Hello world!</h1>
    <p>
      This is the REST API "/" handler from
      https://github.com/harrisonhjones/go-apigw-http-adapter-lambda-example.
    </p>
    <br />
    <img src="project.png" />
    <br />
    <small>
      <p>
        The Go gopher was designed by Renee French.
        (http://reneefrench.blogspot.com/) The design is licensed under the
        Creative Commons 3.0 Attributions license. Read this article for more
        details: https://blog.golang.org/gopher.
      </p>
    </small>
  </body>
</html>
`)
}

func FaviconHandler(writer http.ResponseWriter, _ *http.Request) {
	writer.Header().Set("content-type", "image/svg+xml")
	writer.Write(example.Favicondotsvg)
}

func ImageHandler(writer http.ResponseWriter, _ *http.Request) {
	writer.Header().Set("content-type", "image/png")
	writer.Write(example.Projectdotpng)
}
