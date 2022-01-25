package main
import (
  "fmt"
  "strings"
)
func main(){

  var firstName string
  var lastName string
  var userEmail string
  var nextLang string

  fmt.Println("Enter your first name : ")
  fmt.Scan(&firstName)
  fmt.Println("Enter your last name : ")
  fmt.Scan(&lastName)
  fmt.Println("Enter your email : ")
  fmt.Scan(&userEmail)
  fmt.Println("Next lang to learn :\nRust or Javascript\n-->")
  fmt.Scan(&nextLang)
  
  validEmail := strings.Contains(userEmail, "@")
  if validEmail == true {
    fmt.Printf("Is a valid email!\n")
  } else {
    fmt.Printf("Isn't a valid email, try again!\n")
  }

  validName := len(firstName) >= 2 && len(lastName) >= 2
  if validName == true {
    fmt.Printf("Is a valid name!\n")
  } else {
    fmt.Printf("Isn't a valid name!\n")
  }

  validLang := nextLang == "Rust" || nextLang == "Javascript"
  if !validLang {
    fmt.Printf("Invalid Lang!\n") 
  } else {
    fmt.Printf("Next language to learn: %v\n", nextLang)
  }
}
