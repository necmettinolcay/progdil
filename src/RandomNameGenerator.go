
  package main

  import "flag"
  import "fmt"
 // import "math/rand"
  import "time"
 
  var elements = make(map[string]string)
  var rnd int = int(time.Now().UnixNano())
  var dict map[string]map[string][]string = map[string]map[string][]string{
		"tr": map[string][]string{
			"name":[]string{
				"Ali",
				"Veli",
				"Can",
				"Yıldırım",
				"İsmet",
			},
			"adjective":[]string{
				"Uzun",
				"Kısa",
				"Sinirli",
				"Komik",
				"Sevecen",
				"Gergin",
			},
		},
	}



  type generator interface {
	  nameGenerator() string
	  adjectiveGenerator() string
  }

  type lang struct {
	  name, adjective []string
  }

  func (s lang) nameGenerator() string {
	  return s.name[rnd%len(dict["tr"]["name"])]
  }

  func (s lang) adjectiveGenerator() string {
	  return s.adjective[rnd%len(dict["tr"]["adjective"])]
  }

  func elementsAdd() {
	  tamlama := g.nameGenerator() +  " " + g.adjectiveGenerator()

  func show (g generator) {
	  fmt.Println(g.nameGenerator())
	  fmt.Println(g.adjectiveGenerator())
  }


  func main() {

	  langPtr := flag.String("lang", "tr", "a language")
	  numPtr := flag.Int("numb", 5, "an total")

	  flag.Parse()

	  fmt.Println("word:", *langPtr)
	  fmt.Println("numb:", *numPtr)

          

	  s := lang{name: dict["tr"]["name"], adjective: dict["tr"]["adjective"]}
          show(s)
  }

