package main
import "fmt"

func add(a int, b int) int {
	return a + b
}

func getLanguages()(string,string,string){
	return "golang","javaScript","c++"
}

func processIt() func (a int) int{
	return func(a int) int{
		return 4
	}
}

func main() {
	// result := add(3, 5)
	// fmt.Println(result)

	// lang1, lang2, _ := getLanguages()

	// fmt.Println(lang1,lang2)

	// fn := func(a int) int{
	// 	return a
	// }

	fmt.Println(processIt())
}