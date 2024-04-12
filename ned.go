package main

import (
	"fmt"
	"os"
)

func showHelp(){
	fmt.Println("ned filename")
}

func main(){
	if len(os.Args) == 1 {
		showHelp()
		os.Exit(0)
	}
	fileName := os.Args[1]
	fmt.Println(fileName)
}
