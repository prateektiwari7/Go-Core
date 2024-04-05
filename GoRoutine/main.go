package main

import (
	"fmt"
	"sync"
	"time"
)

var wg = sync.WaitGroup{}
var m = sync.Mutex{}

var dbData = []string{"id1", "id2", "id3", "id4", "id5", "id6", "id7", "id8"}

var result = []string{}

func main() {
	t0 := time.Now()
	for i := 0; i < len(dbData); i++ {
		wg.Add(1)
		go DbCall(i)
	}
	wg.Wait()
	fmt.Println("Total time take is", time.Since(t0))
	fmt.Println("Result is ", result)

}

func DbCall(i int) {
	var delay float32 = 2000
	time.Sleep(time.Duration(delay) * time.Millisecond)
	fmt.Println("The result from database", dbData[i])
	m.Lock()
	result = append(result, dbData[i])
	m.Unlock()
	wg.Done()
}
