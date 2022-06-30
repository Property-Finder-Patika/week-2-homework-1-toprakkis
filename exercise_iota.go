package main

import "fmt"

func main(){
  const (
    C0 = iota
    C1 = iota
    C2 = iota
)
fmt.Println(C0, C1, C2)
  
}
