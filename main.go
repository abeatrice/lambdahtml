package main

import (
	"bytes"
	"context"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

var err error

func main() {
	if os.Getenv("APP_ENV") == "local" {
		localHandler()
	} else {
		lambda.Start(lambdaFunctionUrlHandler)
	}
}

func lambdaFunctionUrlHandler(_ context.Context, request events.LambdaFunctionURLRequest) (events.LambdaFunctionURLResponse, error) {
	fmt.Println("request", request)
	switch request.RawPath {
	case "/":
		fmt.Println("root")
		tpl := template.Must(template.ParseFiles(
			"templates/layout.html",
			"templates/index.html",
		))
		var tplBuffer bytes.Buffer
		logErr(tpl.Execute(&tplBuffer, nil))
		return events.LambdaFunctionURLResponse{Body: tplBuffer.String(), StatusCode: 200}, nil
	case "/about":
		fmt.Println("about")
	case "/nested":
		fmt.Println("nested root")
	case "/nested/about":
		fmt.Println("nested about")
	default:
		fmt.Println("404")
	}
	return events.LambdaFunctionURLResponse{Body: request.Body, StatusCode: 200}, nil
}

func localHandler() {
	http.HandleFunc("/", index)
	http.HandleFunc("/about", about)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("APP_PORT")), nil))
}

func index(w http.ResponseWriter, r *http.Request) {
	display(w, "index", nil)
}

func about(w http.ResponseWriter, r *http.Request) {
	display(w, "about", nil)
}

func display(w http.ResponseWriter, view string, data interface{}) {
	tpl := template.Must(template.ParseFiles(
		"templates/layout.html",
		fmt.Sprintf("templates/%s.html", view),
	))
	logErr(tpl.Execute(w, data))
}

func logErr(err error) {
	if err != nil {
		log.Println(err)
	}
}

func fatalErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
