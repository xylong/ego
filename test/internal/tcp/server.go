package main

import "ego/entity"

func main() {
	server := entity.NewServer("test")
	server.Run()
}
