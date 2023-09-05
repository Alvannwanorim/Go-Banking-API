package main

import "fmt"

func main() {
	fmt.Println(Hello)
	server := APIServer{listenAddr: ":3000"}
	server.Run()
}
