
  package main
  import "flag"
  import "fmt"
  import "math/rand"
  import "time"
  import "strings"


  var langPtr *string = flag.String("lang", "tr", "a language")
  var numbPtr *int = flag.Int("numb", 5, "an total")
  var path string = "../ProjectNames/" + *langPtr + "/" 
  var dict = map[string]map[string][]string {
		"en": {
			"name": {
				"Lion",
				"Leopard",
				"Bird",
				//"Bear",
				//"Monkey",
			},
			"adjective": {
				"Mountain",
				"Grazzy",
				"Crazy",
				//"Nervius",
				//"Funny",
				//"Helpness",
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
	  elements := []string {}
	  dirNames := readDirNames(path)
	  for *numbPtr > 0 {
		  projectName := strings.Join([]string{g.adjectiveGenerator(),g.nameGenerator()}, " " )
		  existPname := controlName(elements, projectName)
		  existDir := controlName(dirNames, projectName)
		  if  existPname == false && existDir == false {
			  elements = append(elements, projectName)
			  *numbPtr --
		  }
	  }
	  show(elements)
	  createDir(elements)
  }

  func show(list []string) {

	  for _, pName := range elements {
		  fmt.Println(pName)
	  }

  }
  func readDir(dirname string) ([]string, error) {
   		f, err := os.Open(dirname)
   		if err != nil {
   			return nil, err
   		}
   		names, err := f.Readdirnames(-1)
   		f.Close()
   		if err != nil {
   			return nil, err
   		}
   		sort.Strings(names)
   		return names, nil	
}
  func createDir(list []string, pName string) {
	os.Mkdir(path + pName, 0777)

}
 func controlName (list []string, pName string) bool {
          
          sort.Strings(list)
          i := sort.SearchStrings(list, pName)
          if  i < len(list) && list[i] == pName {
                  return true
          } else {
                  return false
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
          

  }

