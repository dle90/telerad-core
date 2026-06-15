package thread

// import (
// 	"fmt"
// 	"log"
// )

// // VD2: Multi thread https://200lab.io/blog/golang-channel-la-gi/
// func fibonacci(c, quit chan int) {
// 	log.Printf("fibonacci start")
// 	x, y := 0, 1
// 	for {
// 		select {
// 		case c <- x:
// 			log.Printf("fibonacci x:%v y:%v", x, y)
// 			x, y = y, x+y
// 		case <-quit:
// 			fmt.Println("fibonacci quit")
// 			return
// 		}
// 	}
// }

// func main() {
// 	log.Printf("MAIN start")
// 	c := make(chan int)
// 	quit := make(chan int)
// 	go func() {
// 		for i := 0; i < 10; i++ {
// 			log.Printf("MAIN - for with index:%v", i)
// 			fmt.Println(<-c)
// 		}
// 		log.Printf("MAIN - for END_______-")
// 		quit <- 0
// 	}()
// 	fibonacci(c, quit)
// }
