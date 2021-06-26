package main
import "fmt"

func main(){
	s:=[]int{1,2,3}
	fmt.Println("Enter a factor")
	var factor int
	fmt.Scanln(&factor)
	s=enlarge(s,factor)
	fmt.Printf("The new slice is %d\n",s)
}
func enlarge(s []int, factor int) []int {
	newslice:=make([]int, len(s)*factor)
	copy(newslice,s)
	return newslice
}
