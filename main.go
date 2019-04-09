package main

import (
	"fmt"
	"net/http"
)

func root(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`
	<html>
	<head>
	<title>Hello World</title>
	</head>
	<body>
	<h1>Hello world!!!!</h1>
        <h3>Hello world!!!!</h3>
	<h2>Welcome To Spinnaker!!!!!</h2>
	<h1>Howdy Y'all!!<h1>
	<h1>Pushing directly to master is awesome!</h1>
	</body>
	</html>
	`))
}

func main() {
	fmt.Println("Hello World")
	http.HandleFunc("/", root)
	http.ListenAndServe(":8080", nil)
}
