package main
import "fmt"

func main(){
	fmt.Println("Enter the first number")
	var a int
	fmt.Scanln(&a)
	fmt.Println("Enter the second number")
	var b int
	fmt.Scanln(&b)
	s,p,d := SumProductDiff(a,b)
	fmt.Printf("Sum is %d, Product is %d and Difference is %d",s,p,d)
}

func SumProductDiff(a,b int)(sum int, product int, diff int) {
	sum=a+b
	product=a*b
	diff=a-b
	return sum,product,diff
}
