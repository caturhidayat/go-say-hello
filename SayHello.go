package go_say_hello

import "strconv"

func SayHelloBro(value string, age int) string {
	return "Hello Bro " + value + "You are " + strconv.Itoa(age) + "Year old"
}

func SayGoodBye(name string) string {
	return "Good Bye " + name
}
