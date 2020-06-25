package main

import (
	"net/http"
	"time"
)

type Info struct {
    latency   float64
	tps		  float64
	errorT	  float64
}

func MakeRequest(url string,ch chan Info){
	//Getting cpu% 
	errorT:= 0.0
	//Getting start time
	start := time.Now()
	//Getting the GET request
	resp, err := http.Get(url)
	if err != nil{
		errorT++
	}
	seconds := time.Since(start).Seconds()
	defer resp.Body.Close()
	//Getting the request time on the channel for compute the latency and  usage
	ch<-Info{seconds,0.0,errorT}
}

func Avg(nums []float64) float64{
	sum := 0.0  
	for _,num:= range nums {
		
		sum += num 

	}

	return sum/float64(len(nums))
}

func MakeRequests(n int,c int, url string) Info{
	//Creating the channel for the go rutines comunication
	ch := make(chan Info)
	var infoT Info
	infoT.latency = 0.0
	infoT.errorT = 0.0
	var latencies [] float64
	for i:=0; i < n; i+=c{
		for j:=0; j < c; j++{
			go MakeRequest(url,ch)
			

		}
		inf :=<-ch 
		infoT.errorT += inf.errorT
		latencies = append(latencies,inf.latency)

	}

	infoT.latency = Avg(latencies)
	infoT.tps = float64(n)/infoT.latency
	infoT.errorT = infoT.errorT/float64(n) * 100.0
	return infoT

}