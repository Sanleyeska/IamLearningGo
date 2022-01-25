package main
import "fmt"
func main(){

  var userNickname = "Rue"
  var userNumberPlayer = 99
  var userNumberFloat = 9.9
  fmt.Printf("User nickname is %T, user number player is %T\n", userNickname, userNumberPlayer)
  fmt.Printf("User number %T\n", userNumberFloat)

  /* GoLang can detecte the type o variable, bu if you can set the type manually.*/
  
  var userName string
  var userNumber int
  userName = "Ru"
  userNumber = 99
  fmt.Printf("User %v, number %v", userName, userNumber)

}
