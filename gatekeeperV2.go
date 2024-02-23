package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now().Format("15:04:05")

	if now > "07:00:00" && now < "12:00:00" {
		fmt.Println("goedemorgen Welkom bij Fonteyn Vakantieparken")
	}
	if now > "12:00:00" && now < "18:00:00" {
		fmt.Println("goedemiddag Welkom bij Fonteyn Vakantieparken")
	}
	if now > "18:00:00" && now < "23:00:00" {
		fmt.Println("goedenavond Welkom bij Fonteyn Vakantieparken")
	}
	if now > "23:00:00" && now < "7:00:00" {
		fmt.Println("Sorry, de parkeerplaats is 's nachts gesloten")
	}
}
