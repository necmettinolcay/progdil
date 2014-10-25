
package main
import "fmt"
//import "math/rand"
import "time"

var rnd int = int(time.Now().UnixNano())
var x map[string]map[string][]string = map[string]map[string][]string{
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

func main() {
	

	lenghtName := rnd%len(x["tr"]["name"])

	lenghtAdjective := rnd%len(x["tr"]["adjective"])

	fmt.Println(len(x["tr"]["name"]))
	fmt.Println(time.Now().UnixNano()%100)
	if el, ok := x["tr"]; ok {
		fmt.Println(el["name"][lenghtName], el["adjective"][lenghtAdjective])	
	}
}

