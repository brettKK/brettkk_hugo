package main

import "fmt"

func main() {
	fmt.Println(2)
}

//GOSSAFUNC=main GOOS=linux GOARCH=amd64 go build -gcflags "-S" simple.go && open ssa.html

/*
How a Go Program Compiles down to Machine Code
https://getstream.io/blog/how-a-go-program-compiles-down-to-machine-code/
*/
