package main
import "fmt"

func main(){
	fmt.Println("Enter the number")
	var a int
	fmt.Scanln(&a)
	fmt.Println(fibarray(a))
}
func fibarray(term int) []int {
	slice := make([]int, term)
	slice[0]=1
	slice[1]=1
	for x:=2;x<term;x++ {
		slice[x]=slice[x-1]+slice[x-2]
	}
	return slice
}
