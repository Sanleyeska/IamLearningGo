package main
import "fmt"
func main(){

  /* Example IF */ 
  var1 := 1
  if var1 <= 6 {
    fmt.Printf("%v is less than 6", var1)
  }

  /* IF with a short statement */
  if v := 5; v <= 5 {
    fmt.Printf("\n%v is less or equal to 5", v)
  }

  /* IF and ELSE */
  if x := 5; x == 4 {
    fmt.Printf("\n5 is equal to 4!", x)
  } else {
    fmt.Printf("\n%v Isn't equal to 4!", x)
  }
}
