package main
import "fmt"

func main(){
	fmt.Println
}
func enlarge(s []int, factor int) []int {
	newslice:=make([]int, len(s)*factor)
	copy(newslice,s)
	return newslice
}
