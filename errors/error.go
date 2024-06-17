package errors

import (
	"fmt"
	"net/http"
	"text/template"

	data "test/data"
)

// The Error function takes an http.ResponseWriter and a status code to redirect the user to an error page
func Error(w http.ResponseWriter, statusCode int) {
	var msg string
	switch statusCode {
	case 404:
		msg = "Not Found"
	default:
		msg = "Internal Server Error"
	}
	t, err := template.ParseFiles("./templates/error.page.tmpl")
	if err != nil {
		fmt.Println("Error:", err)
	}

	var test data.Status
	test.Code = statusCode
	test.Msg = msg

	t.Execute(w, test)
	if err != nil {
		fmt.Println("Error:", err)
	}
}
