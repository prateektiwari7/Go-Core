package main

import (
	"fmt"
	"math/rand"
	"time"
)

var max_val float32 = 5

func main() {
	var channel = make(chan string)
	var tofurChannel = make(chan string)
	var website = []string{"BSC", "CoinPayment", "CoinGeko"}
	for i := range website {
		go checkPrice(website[i], channel)
		go checkTourPrice(website[i], tofurChannel)
	}
	sendmessage(channel, tofurChannel)
}

func sendmessage(channel chan string, channelTorf chan string) {
	select {
	case website := <-channel:
		fmt.Println("Text Send, Found deal", website)
	case website := <-channelTorf:
		fmt.Println("Email Send Found deal", website)
	}

}

func checkTourPrice(website string, priceChannel chan string) {
	for {
		time.Sleep(time.Second * 1)
		var price = rand.Float32() * 20
		if price <= max_val {
			priceChannel <- website
			break
		}
	}
}

func checkPrice(website string, priceChannel chan string) {
	for {
		time.Sleep(time.Second * 1)
		var price = rand.Float32() * 20
		if price <= max_val {
			priceChannel <- website
			break
		}
	}
}

func process(c chan int) {
	for i := 0; i < 5; i++ {
		c <- i
	}
	close(c)

}

func simple() {
	var c = make(chan int)
	go process(c)
	for i := range c {
		fmt.Println(i)
	}
	fmt.Println(<-c)
}
