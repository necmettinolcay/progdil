
  package main

  import "flag"
  import "fmt"
 // import "math/rand"
  import "time"
  import "strings"

 
  var langPtr *string = flag.String("lang", "tr", "a language")
  var numbPtr *int = flag.Int("numb", 1, "an total")
  var elements = make(map[string]string)

  var rnd int = int(time.Now().UnixNano())
  var dict map[string]map[string][]string = map[string]map[string][]string{
		"en": map[string][]string{
			"name":[]string{
				"Lion",
				"Leopard",
				"Bird",
				"Bear",
				"Monkey",
			},
			"adjective":[]string{
				"Mountain",
				"Grazzy",
				"Crazy",
				"Nervius",
				"Funny",
				"Helpness",
			},
		},
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
	  return s.name[rnd%len(dict[*langPtr]["name"])]
  }

  func (s lang) adjectiveGenerator() string {
	  return s.adjective[rnd%len(dict[*langPtr]["adjective"])]
  }

  func elementsAdd(g generator) {
	  
	  for *numbPtr > 0 {
		  projectName := strings.Join([]string{g.adjectiveGenerator(),g.nameGenerator()}, " " )
		  if  _, ok := elements[projectName]; !ok {
			  elements[projectName] = "exist"
			  *numbPtr --
		  }
	  }
  }
  
  func show() {

	  for pName, _ := range elements {
		  fmt.Println(pName)
	  }
  }


  func main() {


	  flag.Parse()

	  fmt.Println("word:", *langPtr)
	  fmt.Println("numb:", *numbPtr)

	  s := lang{name: dict[*langPtr]["name"], adjective: dict[*langPtr]["adjective"]}
	  elementsAdd(s)
          show()
  }

