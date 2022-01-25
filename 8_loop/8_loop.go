package main
import "fmt"

func main(){
  /* Loops are simplified in Golang, cuz we've only
  FOR loop.
  */

  /* Forever Loop
  for {

  }
  */

  for x := 1; x <= 9; x++ {
    fmt.Println("Go\n")
  }

  /* FOR-Continued
  The INIT and POST statements are optional.
  */

  sum := 1
  for ; sum < 999;{
    sum += sum
  }
  fmt.Println(sum)

  /* FOR is Go's WHILE */
  firstNumber := 1
  for firstNumber < 999 {
    firstNumber += firstNumber
  }
  fmt.Println(firstNumber)

}
