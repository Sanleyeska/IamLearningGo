package main
import "fmt"
func main(){

  /*  **Dynamic Size** */

  mySlice := [] string{}
  var fisrtName = "Rue"
  var lastName = "Lu"

  mySlice = append(mySlice, fisrtName + " " + lastName)

  fmt.Printf("\nSlice: %v\n", mySlice)
  fmt.Printf("Name: %v\n", mySlice[0])
  fmt.Printf("Slice type: %T\n", mySlice)
  fmt.Printf("Slice length: %v\n", len(mySlice))

}
