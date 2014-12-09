  package main
  import (
           "flag"
	  // "fmt"
	   "math/rand"
	   "time"
	   "strings"
	   "sort"
	   "os"
         )

  var langPtr *string = flag.String("lang", "tr", "a language")
  var numbPtr *int = flag.Int("numb", 5, "an total")

  var dict = map[string]map[string][]string {
		"en": {
			"adjective": {
				"Attractive",
				"Beatiful",
				"Clean",
				"Dark",
				"Fat",
				"Great",
				"High",
				"Little",
				"Misty",
				"Plain",
				"Round",
				"Short",
				"Tall",
				"Ugly",
				"Wide",
			},
			"name": {
				"anvil",
				"balloon",
				"candle",
				"deodorant",
				"eraser",
				"fairy",
				"gel",
				"hockey",
				"ice",
				"javelin",
				"keychain",
				"lace",
				"mallet",
				"nail",
				"picture",
				"quilt",
				"radio",
				"sandpaper",
				"thermometer",
				"vase",
				"wheelbarrow",
				"zipper",
			},
		},
		"tr": {
			"adjective": {
				"Abone",
				"Aceleci",
				"Babacan",
				"Bayraklı",
                                "Canavar",
				"Coşkun",
				"Çemberli",
				"Damgalı",
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
			"name": {
				"anahtarlık",
				"balon",
				"buz",
				"cirit",
				"dantel",
				"deodorant",
				"el arabası",
				"fermuar",
				"hokey",
				"jel",
				"mum",
				"örs",
				"peri",
				"radyo",
				"resim",
				"silgi",
				"termometre",
				"tırnak",
				"tokmak",
				"vazo",
				"yorgan",
				"zımpara",
			},

		},
	}

  type generator interface {
      nameGenerator() string
      adjectiveGenerator() string
  }

  type adjectiveClause struct {
      name, adjective []string
  }

  func (s adjectiveClause) nameGenerator() string {
      return s.name[randomNumber(len(dict[*langPtr]["name"]))]
  }

  func (s adjectiveClause) adjectiveGenerator() string {
      return s.adjective[randomNumber(len(dict[*langPtr]["adjective"]))]
  }

  func elementAdd(g generator) []string {
      elements := []string {}
      path :=  getPath()
      dirNames, _ := readDir(path)

      for *numbPtr > 0 {
          projectName := strings.Join([]string{g.adjectiveGenerator(),g.nameGenerator()}, " " )
	  if  !(isExist(elements, projectName))  && !(isExist(dirNames, projectName))  {
	      elements = append(elements, projectName)
	      *numbPtr --
	  }
      }

      createDir(path, elements)
      return elements
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

  func createDir(path string, list []string) {
      for _, pName := range list {
          os.Mkdir(path + pName, 0777)
      }

  }

  func getPath() string {
      return "../ProjectNames/" + *langPtr + "/"
  }

 func isExist (list []string, pName string) bool {
     sort.Strings(list)
     i := sort.SearchStrings(list, pName)

     return  i < len(list) && list[i] == pName
  }

  func randomNumber(limit int) int {
      rand.Seed(time.Now().UnixNano())
      return rand.Intn(limit)
  }

  func generateName()[]string {
      flag.Parse() //Bayrakların değerlerini öğrenmek için Parse() metodu kullanıldı.
      langOfClause := adjectiveClause{name: dict[*langPtr]["name"], adjective: dict[*langPtr]["adjective"]}
      list := elementAdd(langOfClause)
      return list
  }

