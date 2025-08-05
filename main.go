package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	html := `
		<!DOCTYPE html>
		<html>
		<head>
			<title>Go App</title>
			<style>
				body {
					background-color: #1e1e1e;
					color: #ffffff;
					font-family: Arial, sans-serif;
					text-align: center;
					padding-top: 100px;
				}
				h1 {
					font-size: 48px;
					color: #00ffd5;
				}
			</style>
		</head>
		<body>
			<h1>Hello from Jenkins CI/CD Go App 🚀</h1>
			<p>Build and Deployed via Jenkins!</p>
		</body>
		</html>
	`
	fmt.Fprint(w, html)
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Starting server on port 8080...")
	http.ListenAndServe(":8080", nil)
}
