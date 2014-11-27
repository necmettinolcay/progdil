
  package main
  import "os"
  import "fmt"
  import s "strings"
  import "regexp"

  type MyString struct {
	  arg string
  }

<<<<<<< HEAD
 func (a *rect) ReplaceUnderscoreWithSpace() {
	fmt.Println(a.arg)
	//r, _ := regexp.Compile("[a-z]+_+[a-z]")
	
	r, _ := regexp.Compile("[a-zA-Z0-9]+_+[a-zA-Z0-9]")
	fmt.Println(r.ReplaceAllString(a.arg, s.Replace(r.FindString(a.arg), "_", " " , -1)))

 }
 
=======
  func (m *MyString) ReplaceUnderscoreWithSpace() {
>>>>>>> afe90a0d67de10a2ba5e6f7605c83f8144045562

	fmt.Println(m.arg)
        r, _ := regexp.Compile("[a-zA-Z0-9]+_+\\w*[^_]")
	fmt.Println(r.ReplaceAllString(m.arg, s.Replace(r.FindString(m.arg), "_", " " , -1)))

 }

  func main() {

	arg:= MyString{arg: os.Args[1]}
	arg.ReplaceUnderscoreWithSpace()

 }



