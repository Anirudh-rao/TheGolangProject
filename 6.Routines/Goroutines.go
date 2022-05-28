package main
//Goroutine are LightWeight Threads
import (
	"fmt"
	"time"
)

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}
func main() {
	f("direct")
	//calling Goroutine thread
	go f("goroutine")
	//Calling Goroutine anonymously
	go func(msg string) {
		fmt.Println(msg)
	}("going")
	time.Sleep(time.Second)
	fmt.Println("done")

}
