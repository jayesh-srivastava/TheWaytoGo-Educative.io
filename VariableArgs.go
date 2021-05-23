package main
import "fmt"

func main(){
	fmt.Printf("The sum is %d\n", sumsInt(1,2,3,4,5))
	fmt.Printf("The sum is %d\n", sumsInt(1,4,6,8,4,-6,3,-8))
}

func sumsInt(n ...int)(sum int){
	for _, i := range n {
		sum +=i
	}
	return sum
}
