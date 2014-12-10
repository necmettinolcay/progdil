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

      if (trimString != "") {
          replaceString := s.Replace(trimString, "_", " ", -1)
          m.arg = s.Replace(m.arg, trimString, replaceString, -1)
      }

      show(m.arg)
  }

  func show (finalString string) {
      fmt.Println(finalString)
  }

  func main() {
      arg:= MyString{arg: os.Args[1]}
      arg.ReplaceUnderscoreWithSpace()
  }

