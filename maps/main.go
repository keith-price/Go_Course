package main

import "fmt"

type Colors map[string]string

func main() {
	colors := Colors{}

	colors["red"] = "#fgfgfg"

	colors.addColorToColors("blue", "#5656dty")
	colors.addColorToColors("Limauge", "#68uhh")

	colors.printColorsMap()
}

func (c Colors) addColorToColors(colorName string, colorHex string) {
	c[colorName] = colorHex
}

func (c Colors) printColorsMap() {
	for key, value := range c {
		fmt.Printf("The hex for %v is %v \n", key, value)
	}
}
