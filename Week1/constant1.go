package main
import "fmt"

const Pi = 3.14

const(
	_ = iota + 2
	Monday
	Tuesday
	Wednesday
	Thurday
	Friday
	Saturday
	Sunday
)

func main() {
	fmt.Println("Happy", Pi, "Day")
	fmt.Println(Monday,Tuesday,Wednesday,Thurday,Friday,Saturday,Sunday)
}
