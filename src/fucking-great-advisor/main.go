package main

import "log"
import "fucking-great-advisor/http_client"
import "fucking-great-advisor/image_service"

func main() {
	log.Printf("Starting...")
	http_client.DoSmth()
	image_service.DoSmth()
}
