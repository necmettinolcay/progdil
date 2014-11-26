
  package main
  import "flag"
  import "fmt"
  import "math/rand"
  import "time"
  import "strings"


  var langPtr *string = flag.String("lang", "tr", "a language")
  var numbPtr *int = flag.Int("numb", 5, "an total")
  var elements = make(map[string]string)
  var dict = map[string]map[string][]string {
		"en": {
			"name": {
				"Lion",
				"Leopard",
				"Bird",
				"Bear",
				"Monkey",
			},
			"adjective": {
				"Mountain",
				"Grazzy",
				"Crazy",
				"Nervius",
				"Funny",
				"Helpness",
			},
		},
		"tr": {
			"name": {
				"Ali",
				"Atay",
				"Veli",
				"Can",
				"Yıldırım",
				"İsmet",
			},
			"adjective": {
				"Abone",
				"Aceleci",
				"Babacan",
				"Bayraklı",
                                "Canavar",
				"Coşkun",
				"Çemberli",
				"Damgalı",
				"Demokrat",
				"Eğitimli",
				"Gergin",
				"Ilımlı",
				"Kaba",
				"Komik",
				"Nahif",
				"Paçalı",
				"Rafadan",
				"Romantik",
				"Sevecen",
				"Sinirli",
				"Şakacı",
				"Uzun",

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

	  return s.name[randomNumber(len(dict[*langPtr]["name"]))]

  }

  func (s lang) adjectiveGenerator() string {

	  return s.adjective[randomNumber(len(dict[*langPtr]["adjective"]))]

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

  func randomNumber(limit int) int {

	  rand.Seed(time.Now().UnixNano())
	  return rand.Intn(limit)

  }

  func main() {

	  flag.Parse() //Bayrakların değerlerini öğrenmek için Parse() metodu kullanıldı . ör: --lang = en
          selectedLang := lang{name: dict[*langPtr]["name"], adjective: dict[*langPtr]["adjective"]}
	  elementsAdd(selectedLang)
          show()

  }

