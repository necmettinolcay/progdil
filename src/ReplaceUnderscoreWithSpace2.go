  package main
  import "os"
  import "fmt"
  import s "strings"

  type MyString struct {
	arg string
  }

  func (m *MyString) ReplaceUnderscoreWithSpace() string {
      trimString := s.Trim(m.arg, "_")

      if (trimString != "") {
          replaceString := s.Replace(trimString, "_", " ", -1)
          m.arg = s.Replace(m.arg, trimString, replaceString, -1)
      }

      return m.arg
  }

  func show (finalString string) {
      fmt.Println(finalString)
  }

  func main() {
      arg:= MyString{arg: os.Args[1]}
      fmt.Println(arg)
      fmt.Println(arg.ReplaceUnderscoreWithSpace())
  }

