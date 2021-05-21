package main
import "fmt"

type(
	Celsius float64
	Fahrenheit float64
)

func toFahrenheit(celsius Celsius) (fahrenheit Fahrenheit) {
	fahrenheit=Fahrenheit(celsius*9/5)+32
	return fahrenheit
}
func main() {
	fmt.Println("Enter the temp in Celsius")
	var temp Celsius
	fmt.Scanln(&temp)
	fmt.Printf("The temp in Fahrenheit is %v\n",toFahrenheit(temp))
}
