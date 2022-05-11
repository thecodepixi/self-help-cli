package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
)

// TO DO: use Charm to store these in a KV DB
var toDoList = []string{}

func getUserInput() string {
	reader := bufio.NewReader(os.Stdin)

	input, _ := reader.ReadString('\n')
	text := strings.Replace(input, "\n", "", -1)

	if text == "exit" {
		os.Exit(0)
	}

	return text
}

// eventually use this to have some sort of login functionality
func username() {
	fmt.Println("What's your name?...")

	for {
		fmt.Print("-> ")

		userInput := getUserInput()

		if len(userInput) > 0 {
			fmt.Println("It's nice to meet you", userInput, "!")
			break
		} else {
			fmt.Println("you didn't type anything...")
		}
	}
}

func printToDos() {
	if len(toDoList) > 0 {
		fmt.Println("Here is your To Do list:")
		for i , s := range toDoList {
			fmt.Println(i + 1, ".", s)

		}
	} else {
		fmt.Println("Looks like your To Do list is empty, try adding some things to it!")
	}
}

func addToDos() {
	fmt.Println("What do you need to do today?")
	fmt.Println("you can type in a comma separated list, for example: ")
	fmt.Println("''Go Grocery Shopping, Do Dishes, Pet The Cat''")

	for {
		fmt.Print("-> ")

		userInput := getUserInput()

		if len(userInput) > 0 {
			list := strings.Split(userInput, ",")
			for _ , s := range list {
				toDoList = append(toDoList, strings.TrimSpace(s))
			}
			printToDos()
			break
		} else {
			fmt.Println("you didn't type anything...")
		}
	}
}

func deleteToDo(i int64) {
	copy(toDoList[i:], toDoList[i+1:])
	toDoList = toDoList[:len(toDoList)-1]
}

func removeToDos() {
	length := len(toDoList)

	if length > 0 {
		printToDos()
		fmt.Println("Which items would you like to remove?")

		for {
			fmt.Println("type the number associated with the item you'd like to remove from 1 -", length)
			fmt.Print("-> ")

			userInput := getUserInput()
			index, _ := strconv.ParseInt(userInput, 0, 8)

			deletedItem := toDoList[index]

			deleteToDo(index)

			fmt.Println("Deleted ", deletedItem, "from your To Do list...")
			printToDos()
			break
		}
	} else {
		fmt.Println("Your To Do list is currently empty...")
	}
}

// this is messy, maybe clean this up at some point
func promptLoop() {
	fmt.Println("Welcome to a stupid To Do app...")
	fmt.Println("(type 'exit' at any time to quit")
	username()

	for {
		fmt.Println("What would you like to do? (type the number of your selection)")
		fmt.Println("1. View my To Do list")
		fmt.Println("2. Add to my To Do list")
		fmt.Println("3. Remove items from my To Do list")
		fmt.Print("-> ")
		userInput := getUserInput()

		if userInput == "1" {
			printToDos()
		} else if userInput == "2" {
			addToDos()
		} else if userInput == "3" {
			removeToDos()
		}
	}
}

func main() {
	promptLoop()
}