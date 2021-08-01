package http_client

import "log"

func DoSmth() {
	log.Printf("Hi there from http clent!")
}

/*
TODO:
 - Use go http client to retrieve advice from https://fucking-great-advice.ru/api/random
   Useful links:
	* https://pkg.go.dev/net/http - Go Lang built-in http client
 - Handle JSON response. Put it to Struct (do marshall).
   Useful links:
	* https://mholt.github.io/json-to-go/
	* https://gobyexample.com/json
	* https://habr.com/ru/post/502176/
*/
