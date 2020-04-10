package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

var things = make(map[string](chan int))
var forhotel = make(map[string](chan int))
var input = ""

func scanner() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		go client(scanner.Text())
	}
}

var airline = []string{
	"airaisa001",
	"airaisa002",
	"airaisa003",
	"airaisa004",
}
var hotoy = []string{
	"hotel01",
	"hotel02",
	"hotel03",
}

func client(admin string) {
	co := 0
	for true {
		time.Sleep(time.Duration(rand.Intn(5)) * time.Second)

		choosen := airline[rand.Intn(len(airline))]

		value := <-things[choosen]

		choosenH := hotoy[rand.Intn(len(hotoy))]

		value1 := <-forhotel[choosenH]
		co = co + 1
		if value1 > 3 && value <= 3 {
			fmt.Println(admin + "  : " + choosenH + " is full")
			forhotel[choosenH] <- value1
			things[choosen] <- value
			if co == 2 {
				break
			}
			continue
		} else if value1 <= 3 && value > 3 {
			fmt.Println(admin + "  : " + choosen + " is full")
			forhotel[choosenH] <- value1
			things[choosen] <- value
			if co == 2 {
				break
			}
			continue
		} else if value1 > 3 && value > 3 {
			fmt.Println(admin + "  : " + "full")
			things[choosen] <- value
			forhotel[choosenH] <- value1
			if co == 2 {
				break
			}
			continue
		}

		fmt.Println(admin + "  : " + choosen)
		fmt.Println(admin + "  : " + choosenH)

		value = value + 1
		things[choosen] <- value

		value1 = value1 + 1
		forhotel[choosenH] <- value1

		break
	}

}

func main() {
	fmt.Println("run")

	rand.Seed(time.Now().UnixNano())

	for i := 0; i < len(airline); i++ {
		value := airline[i]
		things[value] = make(chan int, 1)
		things[value] <- 0
	}
	for j := 0; j < len(hotoy); j++ {
		namehotel := hotoy[j]
		forhotel[namehotel] = make(chan int, 1)
		forhotel[namehotel] <- 0

	}
	for true {
		scanner()

	}

}
