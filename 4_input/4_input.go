package main
import "fmt"
func main(){
  var userName string

  fmt.Println("Whats ur name? ")
  fmt.Scan(&userName)
  fmt.Printf("\nHello %v!\n", userName)
  
  /* Using a pointer(&) to see the memory address */
  fmt.Println("Memory Address: ", &userName)
}
