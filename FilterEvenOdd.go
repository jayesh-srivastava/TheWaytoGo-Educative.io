package main
import "fmt"

type flt func(int) bool
func isEven(n int) bool {
	if n%2==0 {
		return true
	}
	return false
}
func filter(s[] int, f flt) (yes, no[] int){
	for _, val := range s {
		if f(val) {
			yes=append(yes, val)
		} else {
			no=append(no, val)
		}
	}
	return
}

func main(){
	slice := [] int {1,2,3,4,5,6,7,8,9,10}
	even, odd := filter(slice, isEven)
	fmt.Printf(" The Even elements are %d\n", even)
	fmt.Printf("The Odd elements are %d", odd)
}
