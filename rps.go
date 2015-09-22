package main

import "fmt"
import "strings"
import "math"
import "math/rand"
import "time"


var validinput = map[string]float64{
  "r":  0,
  "p":  1,
  "s":  2,
  "rock": 0,
  "paper": 1,
  "scissors": 2,
}

// not sure if really need reverse map, anyway to do this with original?
var reverseinput = map[float64]string{
  0 : "Rock",
  1 : "Paper",
  2 : "Scissors",
}

var input string
var player float64
var exists bool
var result string

func getinput() string {
  // get input and use strings package to convert to lowercase
  fmt.Print("Pick [r]ock, [p]aper, or [s]issors:   ")
  var input string
  fmt.Scanln(&input)
  input = strings.ToLower(input)
  return input
}


func checkwinner(player, computer float64) string {
  compare := math.Mod((computer - player) ,3)
  if compare < 0 {
      // fixing for the way remainders work with neg numbers in golang
      compare = compare + 3
  }

  if compare == 1 {
    result = "You Lost :("
    return result
  }
  if compare == 2 {
    result = "You Win!!  :)"
    return result
  } else {
    result = "You Tied"
    return result
  }

}

func checkinput() string{
  // start a loop until we get valid input
  for {
    input := getinput()
    // Does the user input match the map validinput?
    player, exists =validinput[input]
      if exists == true{
        fmt.Println("Player Chooses", reverseinput[player])
        return input
      }  else {
        fmt.Println("I didn't understand your choice, please retry")
      }
  }

}

func random(min, max int) int {
    rand.Seed(time.Now().Unix())
    return rand.Intn(max - min) + min
}

func main() {

  // Get Valid Player Input
  input = checkinput()

  // Randomly Assign Computer Choice
  computer := float64(random(0,3))
  cpuinput, _ := reverseinput[computer]
  fmt.Println("Computer Chooses" , cpuinput)


  // Check to see who won
  result = checkwinner(player, computer)
  fmt.Println("Result:", result)


}