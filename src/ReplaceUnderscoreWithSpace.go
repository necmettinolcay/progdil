
  package main
  import "os"
  import "fmt"
  import s "strings"
  import "regexp"

  type MyString struct {
	  arg string
  }

  func (m *MyString) ReplaceUnderscoreWithSpace() {
      fmt.Println(m.arg)
      r, _ := regexp.Compile("[a-zA-Z0-9]+_+\\w*[^_]")
      fmt.Println(r.ReplaceAllString(m.arg, s.Replace(r.FindString(m.arg), "_", " " , -1)))
  }

  func main() {
      arg:= MyString{arg: os.Args[1]}
      arg.ReplaceUnderscoreWithSpace()
  }



