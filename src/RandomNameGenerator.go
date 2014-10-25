
  package main

  import "flag"
  import "fmt"
 // import "math/rand"
  import "time"

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
	  return s.name[int(time.Now().UnixNano())%len(dict["tr"]["name"])]
  }

  func (s lang) adjectiveGenerator() string {
	  return s.adjective[int(time.Now().UnixNano())%len(dict["tr"]["adjective"])]
  }

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

