package main
import "fmt"

var fib[10]int64

func main(){
	 arr:=fibs()
	 for x:=0;x<=9;x++ {
	 	fmt.Printf("The %dst Fibonacci Number is %d\n",x+1,arr[x])
	 }
}

func fibs() [10]int64 {
	fib[0]=1
	fib[1]=1
	for x:=2;x<=9;x++ {
		fib[x]=fib[x-1]+fib[x-2]
	}
	return fib
}