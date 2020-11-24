package main
import "fmt"
// func average(scores [5]float64) float64 {
// 		total := 0.0
//   for _, score := range scores {
//     total += score
// 	}
// 	return total / float64(len(scores))
// }

// func main() {
// 	scores := [5]float64{1,2,3,4,5}
//   fmt.Println(average(scores))
// }

// Part 2
// var initialPets map[string]string = map[string]string {
// 	"frido" : "dog",
// 	"nuffi" : "cat",
// 	"kubus" : "horse",
// }

// func petExists(petName string) bool{
// 	_, ok := initialPets[petName] 
// 	return ok
// }

// func main() {
// 	petName := "frido"
// 	fmt.Println(petExists(petName))
// }

// Part 3
// var groceries = []string{"apple", "lemon", "eggs"}
// func totalGroceries(additionalGroceries ...string) []string{
// 	foods := groceries
// 	for _, g := range additionalGroceries {
// 		foods = append(groceries, g)
// 	}
// 	return foods
// }
// func main() {
// 	fmt.Println(totalGroceries("milk"))
// }