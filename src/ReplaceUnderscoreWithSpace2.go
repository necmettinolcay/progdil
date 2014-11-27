
  package main
  import "os"
  import "fmt"
  import s "strings"

  type MyString struct {
	  arg string
  }

  func (m *MyString) ReplaceUnderscoreWithSpace() {
	fmt.Println(m.arg)
	trimString := s.Trim(m.arg, "_")
	replaceString := s.Replace(trimString, "_", " ", -1)
	finalString := s.Replace(m.arg, trimString, replaceString, -1)
        Show(finalString)
  }
  func Show (finalString string) {
	   fmt.Println(finalString)
  }

  func main() {
	arg:= MyString{arg: os.Args[1]}
	arg.ReplaceUnderscoreWithSpace()
  }



