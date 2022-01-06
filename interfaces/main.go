package main

import "fmt"

type bot interface {
	getGreeting() string
}

type englishBot struct{}
type spanishBot struct{}

func main() {

	english_b := englishBot{}
	printGreeting(english_b)
	spanish_b := spanishBot{}
	printGreeting(spanish_b)

}

func (eb englishBot) getGreeting() string {
	// since we are not making use of the reciever object, so we can remove the object from func declaration, like so
	// func (englishBot) getGreeting() string {

	return "Hi there!"

}

func (sb spanishBot) getGreeting() string {

	return "Hola!"

}

func printGreeting(b bot) {

	fmt.Println(b.getGreeting())

}

// same func name with different signatures will not be allowed
// func printGreeting(eb englishBot) {

// 	fmt.Println(eb.getGreeting())

// }

// func printGreeting(sb spanishBot) {
// 	fmt.Println(sb.getGreeting())
// }
