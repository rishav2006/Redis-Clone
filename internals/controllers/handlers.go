package controllers

import (
	"bufio"
	"fmt"
	// "io"
	"os"
	"strings"

	// "github.com/gin-gonic/gin"
	"github.com/rishav2006/redis-clone/internals/models"
)

// func TakeInput1(c *gin.Context) string { // Responsible for taking the input
// 	fmt.Println("Enter the command")
// 	input := string(jsonData)
// 	// reader := bufio.NewReader(os.Stdin)
// 	if input == "" {
// 		fmt.Println("Error : Please provide some input")
// 		return ""
// 	}
// 	return input
// }

func TakeInput() string{
	fmt.Println("Enter the command")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	if input == ""{
		fmt.Println("Error : Please provide some input")
		return ""
	}
	return input
}

func StringParser(str string) (string, []string) { // Responsible for Parsing the string
	words := strings.Fields(str)
	firstWord := words[0]
	remainingWords := words[1:]
	return firstWord, remainingWords
}

func isExists(str string) (bool, string) {	// Checks if the key exists in the hashmap
	val := models.SETHashCommand[str]
	if val == "" {
		return false, ""
	} else {
		return true, val
	}
}

func EXISTS(str string) {	// Calls isExists and prints value accordingly
	check, _ := isExists(str)
	if check == true {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}

func GetArranger(str string) {	// GET 
	check, val := isExists(str)
	if check == true {
		fmt.Println(val)
	} else {
		fmt.Println("Invalid Operation - No such value exists")
	}
}

func SetArranger(firstWord string, rem []string) {
	models.SETHashCommand[rem[0]] = rem[1]
	fmt.Println("OK")
}

func Checker(firstWord string, rem []string) {
	if firstWord == "SET" {
		if len(rem) != 2 {
			fmt.Println("Check again")
			return
		}
		SetArranger(firstWord, rem)
	} else if firstWord == "GET" {
		if len(rem) != 1 {
			fmt.Println("Check again")
			return
		}
		GetArranger(rem[0])
	} else if firstWord == "EXISTS" {
		if len(rem) != 1 {
			fmt.Println("Check again")
			return
		}
		EXISTS(rem[0])
	}
}

func Organizer() {
	var input string = TakeInput()
	firstWord, rem := StringParser(input)
	Checker(firstWord, rem)
}
