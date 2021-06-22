package main
import "fmt"

func main() {
	var arr [15]int
	for x:=0;x<=14;x++ {
		arr[x]=x
	}
	fmt.Println(arr)
}
