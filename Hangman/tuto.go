package main

import (
	"fmt"
	"time"
)

func tuto() {
	var rep string
	
    
	fmt.Println("\nTo play Hang Fruit, you need a lot of imagination, a dose of logic, and above all a love of fruit =)\n\nFirst of all, let's start with a simple word. Okay ?")
	fmt.Scanln(&rep)
	if rep == "yes"{
		fmt.Println("You will see a number of dashes\n\n_ _ _ _ _\n\nhow many dashes do you see ?")
	}
	fmt.Scanln(&rep)
	if rep == "5"{
		fmt.Println("Fine !")
	}
	time.Sleep(1 * time.Second)
	fmt.Println("We know that the word contains five letters,  Good!")
	time.Sleep(2 * time.Second)
	fmt.Println("As you have tested our tutorial, you will be entitled to a hint. =) ")
	time.Sleep(2 * time.Second)
	fmt.Println("All the hidden words in our game are fruit names. ")
	time.Sleep(2 * time.Second)
	fmt.Println("Now that we know that, we have two clues…\nLet this be a 5-letter word and this a fruit")
	time.Sleep(2 * time.Second)
	fmt.Println("As a five-letter fruit name we have : Peach, Prune , Apple")
	time.Sleep(2 * time.Second)
	fmt.Println("\nLet's try Apple ! Who doesn't love apples ?")
	fmt.Println("\n_ _ _ _ _")
	time.Sleep(1 * time.Second)
	fmt.Println("\n( Insert a )")

	fmt.Scanln(&rep)
	if rep == "a"{
		fmt.Println("\na _ _ _ _")
	}
	time.Sleep(1 * time.Second)
	fmt.Println("\nAh! maybe we're on the right track ")
	time.Sleep(1 * time.Second)
	fmt.Println("\na _ _ _ _")
	fmt.Println("\n( Insert p )")

	fmt.Scanln(&rep)
	if rep == "p"{
		fmt.Println("\na p p _ _")
	}
	fmt.Println("\n( Insert l )")

	fmt.Scanln(&rep)
	if rep == "l"{
		fmt.Println("\na p p l _")
	}

	fmt.Println("\n( Insert e )")

	fmt.Scanln(&rep)
	if rep == "e"{
		fmt.Println("\na p p l e")
	}

	fmt.Println("\nGood Job = )")
	time.Sleep(1 * time.Second)
	fmt.Println("You have now found the right word ")
	time.Sleep(1 * time.Second)
	fmt.Println("You are now ready to play our game. ")
	time.Sleep(1 * time.Second)
	fmt.Println("\nBack to the menu")
	time.Sleep(3 * time.Second)

	fmt.Println("")
	
	menu()
	
}
