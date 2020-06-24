package main

import (
    "fmt"
	"net/http"
	"time"
	"github.com/shirou/gopsutil/cpu"
)

type Info struct {
    latency   float64
    cpu		  float64
}

func MakeRequest(url string,ch chan Info){
	cpuT,_ := cpu.Percent(time.Second,true)
	//Getting start time
	start := time.Now()
	//Getting the GET request
	resp, err := http.Get(url)
	if err != nil{
		panic(err)
	}
	seconds := time.Since(start).Seconds()
	defer resp.Body.Close()
	//Getting the request time on the channel for compute the latency

	ch<-Info{seconds,cpuT[0]}
}

func Avg(nums []float64) float64{
	sum := 0.0  
	for _,num:= range nums {
		
		sum += num 

	}

	return sum/float64(len(nums))
}

func MakeRequests(n int,c int, url string){
	//Creating the channel for the go rutines comunication
	ch := make(chan Info)
	var infoT Info
	infoT.latency = 0.0
	infoT.cpu = 0.0
	var cpuMean [] float64
	for i:=0; i < n; i+=c{
		for j:=0; j < c; j++{
			go MakeRequest(url,ch)
			

		}
		inf :=<-ch 
		cpuMean = append(cpuMean,inf.cpu)
		infoT.latency+=inf.latency

	}
	fmt.Printf("Latency: %f\n",infoT.latency)
	fmt.Println("Percentage CPU (avg): ",Avg(cpuMean))

}