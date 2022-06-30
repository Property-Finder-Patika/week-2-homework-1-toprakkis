// We can show the version of our Golang runtime with the help of Version function from runtime package

package main

 import (
   
   "fmt"
   "runtime"
 )

func main(){
  fmt.Println(runtime.Version()) 
}
