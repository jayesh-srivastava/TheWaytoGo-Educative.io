package main
import "fmt"

func main(){
	sl:=[]int{2,3,6,8,4,9,5,1}
	fmt.Println("Original Slice")
	fmt.Println(sl)
	fmt.Println("Sorted Slice")
	fmt.Println(bubbleSort(sl))

}

func bubbleSort(sl []int) []int{
	for i:=0;i< len(sl)-1;i++ {
		for j:=0;j< len(sl)-i-1;j++ {
			if sl[j]>sl[j+1] {
				sl[j], sl[j+1]=sl[j+1], sl[j]
			}
		}
	}
	return sl
}
