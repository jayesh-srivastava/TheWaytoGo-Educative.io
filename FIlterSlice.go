package main
import "fmt"

func main() {
	s:=[]int{0,1,2,3,4,5,6,7,8,9}
	s=Filter(s, even)
	fmt.Println(s)
}

func Filter(s []int, fn func(int) bool) []int {
	var e []int
	for _, i:=range s {
		if fn(i) {
			e=append(e,i)
		}
	}
	return e
}

func even(n int) bool {
	if n%2 == 0 {
		return true
	}
	return false
}
