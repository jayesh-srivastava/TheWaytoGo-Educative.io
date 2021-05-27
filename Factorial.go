package main
import "fmt"

func main(){
	fmt.Println("Enter the number")
	var n uint64
	fmt.Scanln(&n)
	fmt.Printf("The factorial is %d",Factorial(n))
}

func Factorial(n uint64)(fac uint64) {
	if n<=1 {
		return 1
	}
	fac=n*Factorial(n-1)
	return fac
}
