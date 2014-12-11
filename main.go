package main

import(
  "os"
  "fmt"
)

func main() {
  cmd_io()
  arrays()
  paralel()
}

func cmd_io(){
  // Definetion and initialization of strings
  s0 := "Hello "
  var s1 = "World"
  var s2 string = "!\n"

  os.Stdout.WriteString(s0+s1+s2)
}

func arrays(){
  // Definetion an array with length 6
  var a [6]int

  // Definetion and initialization of an array
  strArray := [3]string{"foo", "bar", "foobar"}
 
  // Set the 6th elemet to 6
  a[5] = 6
  fmt.Println(a)

  // Foor-loop
  for i:=0; i<len(a); i++ {
    a[i]=i+1
  }
  fmt.Println(a)

  // Cast of datatypes
  var a1 float64
  for i:=0; i<len(a); i++ {
    a1 += float64(a[i])
  }  
  fmt.Println(a1/float64(len(a)))

 // For-loop with rage
 for i, val := range a {
   // Format text output
   s := fmt.Sprintf("index: %d, value: %d", i, val)
   fmt.Println(s)
 }

 // For-loop with rage and _
 for _, val := range strArray {
   fmt.Println(val)
 } 

}

func paralel() {
  // Create channel
  ch := make(chan int)
  
  // Start routines 
  go goRoutine0(ch)
  go goRoutine1(ch)
    
  for i:=0; i<11; i++ {
    ch <- i
  }
  fmt.Println("Paralel finished") 
 
}

func goRoutine0(c chan int){
  for{
    i := <- c
    res := i
    for j:=0; j<i; j++{
      res = res*i
    }
    fmt.Printf("go0: %d^2 %d \n", i, res)
  }
}

func goRoutine1(c chan int){
  for{
    i := <- c
    fmt.Printf("go1: %d \n", i)
  }
}
