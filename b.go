package main

import (
  "fmt"
  "time"
)

func main(){
  st := time.Now()
  n := 0
  for i := 0; i < 100000000; i++{
   n += 1+1-1*1/1
  }
  t := time.Now()
  fmt.Println("b:",t.Sub(st),n)
}
