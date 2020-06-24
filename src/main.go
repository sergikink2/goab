package main

import (
	"os"
	"strconv"

)

var number int
var concurrency int
var urlhost string

func chooseOption(args []string){

	for i := 0; i < len(args)-1; i++ {

		switch args[i]{
		case "-n":
			number,_ = strconv.Atoi(args[i+1])
			break
		case "-c":
			concurrency,_ = strconv.Atoi(args[i+1])
			break

			
		}

	}

}

func main() {
	
	args := os.Args[1:]
	number = 1
	concurrency = 1
	urlhost = args[len(args)-1]
	chooseOption(args)
	MakeRequests(number,concurrency,urlhost)

}