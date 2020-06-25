package main

import (
	"os"
	"strconv"
	"fmt"

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
	var inf Info = MakeRequests(number,concurrency,urlhost)
	
	fmt.Printf("Average latency: %0.3f\n",inf.latency)
	fmt.Printf("TPS: %.3f\n",inf.tps)
	fmt.Printf("Error rate: %.2f\n",inf.errorT)

}