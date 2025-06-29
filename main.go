package main

import (
	"fmt"

	go_say_hello "github.com/nafifurqon/go-say-hello/v2"
)

func main() {
	// fmt.Println(go_say_hello.SayHello())
	fmt.Println(go_say_hello.SayHello("Nafi"))
}

/*
	notes:
	go mod init github.com/nafifurqon/app-say-hello -> for initialize and craete file go.mod
	go get github.com/nafifurqon/go-say-hello -> for download the module
*/
