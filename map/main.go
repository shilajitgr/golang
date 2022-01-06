package main

import "fmt"

func main() {

	// following are types of declaration of map objects :
	// var colors map[string]string
	// colors := make(map[string]string)

	colors := map[interface{}]interface{}{
		"red":   "#ff0000",
		"blue":  "#00ff00",
		"green": "#0000ff",
	}

	colors["black"] = "#ffffff"

	// to delete a map element

	delete(colors, "black")

	// fmt.Println(colors)
	// printcolors(colors)
	new_colors := colors
	new_colors["black"] = "new0000"
	colors["zman"] = "eneha"
	colors["anma"] = "23dwd"
	colors[23] = 34567
	printcolors(new_colors)
	printcolors(colors)
}

func printcolors(c map[interface{}]interface{}) {
	for color, hex := range c {
		fmt.Printf("%v : %v\n", color, hex)
	}
}
