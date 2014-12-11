package main

import "fmt"

func main(){
  primes := make(chan int)
  p := 0
  go generate(primes)
  for i:=0; i<100; i++ {
    prime := <- primes
    in := primes
    out := make(chan int)
    go devideable(in, out, prime)
    primes = out
    p = prime
  }
  fmt.Println(p)
}

func generate(out chan int){
  for i:=2;; i++{
    out <- i
  }
} 

func devideable(input chan int, output chan int, p int){
  for{
    x := <- input
    if x%p != 0{
      output <- x 
    }
  }
}
