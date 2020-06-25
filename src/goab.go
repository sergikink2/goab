package main

import (
	"os"
	"strconv"
	"fmt"

)

var number int //number of petitions
var concurrency int //level of concurrency
var urlhost string //server we make petitions

func chooseOption(args []string){
	//Getting the number of petitions and level of concurrency
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
	number = 1 //default petitions
	concurrency = 1 //default concurrency
	urlhost = args[len(args)-1]
	chooseOption(args)
	//Getting the info of the requests
	var inf Info = MakeRequests(number,concurrency,urlhost)
	//Show the info of the requests
	fmt.Printf("Average latency: %0.3f\n",inf.latency)
	fmt.Printf("TPS: %.3f\n",inf.tps)
	fmt.Printf("Error rate: %.2f\n",inf.errorT)

}