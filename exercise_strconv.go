// Integer to string with strconv

package main
import(
  "fmt"
  "strconv"
)

func main(){
  a := 32
  b:= strconv.Itoa(a)
  fmt.Printf("%T, %T" ,a,b)
}
