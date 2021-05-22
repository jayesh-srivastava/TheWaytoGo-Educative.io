package main
import "fmt"

func main(){
	fmt.Println("Enter the month number")
	var n int
	fmt.Scanln(&n)
	fmt.Println(Season(n))
}
func Season(number int) (season string) {
	switch number {
	case 1:
		season="January and it's Winter"
	case 2:
		season="February and it's Winter"
	case 3:
		season="March and it's Spring"
	case 4:
		season="April and it's Spring"
	case 5:
		season="May and it's Spring"
	case 6:
		season="June and it's Summer"
	case 7:
		season="July and it's Summer"
	case 8:
		season="August and it's Summer"
	case 9:
		season="September and it's Autumn"
	case 10:
		season="October and it's Autumn"
	case 11:
		season="November and it's Autumn"
	case 12:
		season="December and it's Winter"
	default:
		season="Invalid Input"
	}
	return season
}
