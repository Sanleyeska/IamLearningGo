package main
import "fmt"
func main(){

  var myArray [4] string
  myArray[1] = "Rue"
  myArray[2] = "Lu"
  
  fmt.Printf("Array: %v\n", myArray)
  fmt.Printf("First Name: %v\n", myArray[1])
  fmt.Printf("Last Name: %v\n", myArray[2])
  fmt.Printf("Array type: %T\n", myArray)
  fmt.Printf("Array length: %v\n", len(myArray))

}
