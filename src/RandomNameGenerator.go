
  package main

  import "flag"
  import "fmt"
  import "math/rand"

  type genarator interface {
	  nameGenarator() string
	  adjectiveGenarator() string
  }
  type tr struct {
	  name, adjective []string
  }
  type en struct {
	  name, adjective []string
  }
  
  func (s tr) nameGenarator() string  {
	  return s.name[rand.Intn(2)]   
  }
  
  func (s tr) adjectiveGenarator() string {
	  return s.adjective[rand.Intn(2)]
  }

  func (c en) nameGenarator() string  {
	  return  c.name[0]
  }
  
  func (c en) adjectiveGenarator() string {
	  return  c.adjective[0]
  }

  func show (g genarator) {
	  fmt.Println(g.nameGenarator() )
	  fmt.Println(g.adjectiveGenarator())
	  
  }


  func main() {
	  
	  langPtr := flag.String("lang", "tr", "a language")
	  numPtr := flag.Int("numb", 5, "an total")
	  
	  flag.Parse()

	  fmt.Println("word:", *langPtr)
	  fmt.Println("numb:", *numPtr)

          n := []string{"Ali","Veli","Can","Kalem"}
	  k := []string{"Uzun", "Kisa","Duvar", "Kitap"}

	  s := tr{name: n, adjective: k}
          show(s)

  }

