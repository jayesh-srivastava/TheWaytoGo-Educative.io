package main
import "fmt"

func main() {
	fmt.Println("Enter the string")
	var s string
	fmt.Scanln(&s)
	fmt.Printf("The reversed string is %s", reverse(s))
}

func reverse(s string) string {
	s2 := []byte(s)
	for i,j :=0, len(s2)-1 ; i<j ; i, j = i+1,j-1 {
		s2[i], s2[j] = s2[j], s2[i]
	}
	return string(s2)
}
