
  package main
  import "fmt"
  import "time"
  import "math/rand"

  func main() {
	  show()
	  show()

	  
  }

  func show() {
	  
	  rand.Seed(time.Now().UnixNano())
	  fmt.Println(rand.Intn(100))
  }
