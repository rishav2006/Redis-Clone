package controllers

import (
	"bufio"
	"fmt"
	"strconv"
	"time"

	// "io"
	"os"
	"strings"

	// "github.com/gin-gonic/gin"
	"github.com/rishav2006/redis-clone/internals/store"
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

func TakeInput() string {
	fmt.Println("Enter the command")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	if input == "" {
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

func isExists(str string) (bool, string) { // Checks if the key exists in the hashmap
	store.DB.Mu.RLock()
	var val string
	expiry, exists := store.DB.Expiration[str]
	if exists {
		if time.Now().After(expiry) {
			delete(store.DB.Data, str)
			delete(store.DB.Expiration, str)
			val = ""
		} else {
			val = store.DB.Data[str]
		}
	} else {
		val = store.DB.Data[str]
	}
	store.DB.Mu.RUnlock()
	if val == "" {
		return false, ""
	} else {
		return true, val
	}
}

func EXISTS(str string) string { // Calls isExists and prints value accordingly
	check, _ := isExists(str)
	if check == true {
		return "YES"
	} else {
		return "NO"
	}
}

func TTLchecker(str []string) (string, int) {
	s := str[0]
	num, err := strconv.Atoi(str[1])
	if err != nil {
		fmt.Println("Error occured while TTL Checking: ", err)
	}
	return s, num
}

func GetArranger(str string) string { // GET
	check, val := isExists(str)
	if check == true {
		return val
	} else {
		return "Invalid Operation - No such value exists"
	}
}

func TimeDeterminer(num int) time.Time {
	newTime := time.Now().Add(time.Duration(num) * time.Second)
	return newTime
}

func SetExArranger(rem []string) string {
	store.DB.Mu.Lock()
	firstWord := rem[0] // extract the first word...ex - name, city
	words := rem[1:]    // put the rest words all along...ex - Rishi 60
	str, num := TTLchecker(words)
	store.DB.Data[firstWord] = str
	store.DB.Expiration[firstWord] = TimeDeterminer(num) // what to write here ?
	store.DB.Mu.Unlock()
	resultStr := fmt.Sprintf("Okay. Data will expire after %d secs", num)
	return resultStr
}

func SetArranger(rem []string) string {
	store.DB.Mu.Lock()
	store.DB.Data[rem[0]] = rem[1]
	store.DB.Mu.Unlock()
	return "Okay"
}

func Checker(firstWord string, rem []string) string {
	if firstWord == "SET" {
		if len(rem) != 2 {
			return "Check again"
		}
		return SetArranger(rem)
	} else if firstWord == "GET" {
		if len(rem) != 1 {
			return "Check again"
		}
		return GetArranger(rem[0])
	} else if firstWord == "EXISTS" {
		if len(rem) != 1 {
			return "Check again"
		}
		return EXISTS(rem[0])
	} else if firstWord == "SETEX" {
		if len(rem) != 3 {
			return "Check again"
		}
		return SetExArranger(rem)
	}
	return "Invalid Input: Please check your command"
}

func Organizer(input string) string {
	// var input string = TakeInput()
	firstWord, rem := StringParser(input)
	return Checker(firstWord, rem)
}
