package main

import (
	"fmt"
	"time"
)

var (
	consume_a_drink       chan int
	empty_bottles_on_hand chan int
	empty_caps_on_hand    chan int
)

func exchanges_with_empty_bottle() {
	for {
		// Give away 2 empty_bottles_on_hand exchange for 1 new drink !
		<-empty_bottles_on_hand
		<-empty_bottles_on_hand

		consume_a_drink <- 1
	}
}

func exchanges_with_empty_cap() {
	for {
		// Give away 4 empty_caps_on_hand exchange for 1 new drink !
		<-empty_caps_on_hand
		<-empty_caps_on_hand
		<-empty_caps_on_hand
		<-empty_caps_on_hand

		consume_a_drink <- 1
	}
}

func consume_drinks() {
	count := 0
	for {
		<-consume_a_drink
		// consume a drink, added one empty bottle and one empty cap on hand
		empty_bottles_on_hand <- 1
		empty_caps_on_hand <- 1
		count = count + 1
		fmt.Print(count, " ")
	}
}

func main() {
	say_the_rules()
	fmt.Print("Be patient, now I am counting the beers I consumed: ")

	max := 10 // just a magic number as a maximum for our processes communications
	consume_a_drink = make(chan int, max)
	empty_bottles_on_hand = make(chan int, max)
	empty_caps_on_hand = make(chan int, max)

	// 10 dollar, means we start with 5 drinks
	consume_a_drink <- 1
	consume_a_drink <- 1
	consume_a_drink <- 1
	consume_a_drink <- 1
	consume_a_drink <- 1

	go consume_drinks()
	go exchanges_with_empty_bottle()
	go exchanges_with_empty_cap()

	time.Sleep(time.Second * 2) // Simply end by wait for a maximum of 10 seconds for the calculation
	say_the_end()
}

func say_the_rules() {
	fmt.Println("IQ Test Rules:")
	fmt.Println("     Beer sold at $2 each.")
	fmt.Println("     4 empty bottle caps can be exchanged for a new Beer")
	fmt.Println("     2 empty bottles can be exchanged for a new Beer")
	fmt.Println("The Question is:")
	fmt.Println("     How many beers can I drink with $10 ?")
	fmt.Println("")
	fmt.Println("")
	time.Sleep(time.Second * 1)
}
func say_the_end() {
	fmt.Println("")
	fmt.Println("That's All.")
	fmt.Println("So now, I am drunk, you can try to see how many bears can I drink with $20.")
}
