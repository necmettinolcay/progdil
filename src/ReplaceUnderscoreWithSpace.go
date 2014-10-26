 
  package main
  import "os"
  import "fmt"
  import s "strings"
  import "regexp"

 

  type rect struct {
	  arg string
  }

 func (a *rect) ReplaceUnderscoreWithSpace() {
	fmt.Println(a.arg)
	//r, _ := regexp.Compile("[a-z]+_+[a-z]")
	
	r, _ := regexp.Compile("[a-zA-Z0-9]+_+[a-zA-Z0-9]")
	fmt.Println(r.ReplaceAllString(a.arg, s.Replace(r.FindString(a.arg), "_", " " , -1)))

 }
 



  func main() {

	  arg := rect{arg: os.Args[1]}
	  arg.ReplaceUnderscoreWithSpace()

  }


