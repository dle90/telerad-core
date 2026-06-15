package thread

// import (
// 	"log"
// 	"time"
// )

// VD1: Multi thread
// func main() {
// 	pool := NewPool(3) // creates pool with limited concurrency of 3
// 	for i := 0; i < 10; i++ {
// 		pool.Add(1) // This will wait until a slot is available
// 		go work(i, pool)
// 	}

// 	pool.Wait()
// 	log.Println("All Done !")
// }

// func work(i int, pool *GoPool) {
// 	defer pool.Done() // just like with sync.WaitGroup
// 	log.Printf("working hard on %v", i)
// 	time.Sleep(time.Second)
// 	log.Printf("%v is done", i)
// }
