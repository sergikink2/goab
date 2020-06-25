# goab

## Analysis of Apache Benchmark

### Results

 N | C | CPU(mean- ms) | Latency (s) 
--- | ---| --- | ---
10  |1 |139| 2.309
10  |2 |112| 1.199
10  |3 |90| 0.699
10  |4 |542| 2.419
100 |1 |93| 20.171
100 |2 |119| 11.806
100  |3 |134| 8.516
100  |4 |139| 2.309

## Observations & conclusions

We can see after testing some cases that if we increase the number of concurrency the latency decreases. Although when we reach 4 of concurrency and an n equals 10 this fenomen changes. It’s possible this could happen because the number of packages being sended are a few. We observe that with an n equal 100 turns out the same behaviour as before, so we can conclude the number of packages sended and the concurrency are strongly related. 

The CPU time is not so clear, due to the difference between n = 10 and n = 100. On one hand when  n = 10 the same patron has the latency is obtained, with the difference that this time when we have concurrency of 4 the latency improves. Meaning it is right to say that with a large number of packages the increase of concurrency does not affect the performance. But on the other hand when n = 100 is the contrary. This could occur because the distribution of packages is larger and the values are more spread.

## Goab vs Apache Benchmark

### Results of testing

 N | C | AVG Latency (s) 
--- | --- | ---
10  |1 | 0.067
10  |2 | 0.243
10  |3 | 0.312
10  |4 | 0.356
100 |1 | 0.086
100 |2 | 0.126
100 |3 | 0.169
100 |4 | 0.165

In comparison to apache benchmark the average latency is lesser with goab, this doesn't mean goab is better. 

If we analyse it we can arrive to the conclusion that this difference could be happening because of the application’s testing. The testing of the server is done locally, for that reason the time to connect with it is shorter.

In addition we can observe  that the average latency increases with the concurrency level, this succeeds due to the number of requests being packed are more than one, in conclusion the time for processing them is greater. 


## Guidelines

Goab execution example:
  - We choose the level of concurrency and number of requests, we do it with  the flags -c and -n. Example: 
      
      `./goab -c 4 -n 100 https://www.docker.com/`
     
      Be default n and c are 1

      Output: 
      
        Average latency: 0.049
        TPS: 2034.478
        Error rate: 0.00´´
      
      





