package main

import (
	"fmt"
	"time"
)

func kenteken() {
	now := time.Now().Format("15:04:05")
	licensePlateslist := []string{"ABC123", "DEF456", "GHI789"}

	licensePlate := "ABC123"

	found := false
	for _, plate := range licensePlateslist {
		if plate == licensePlate {
			found = true
			break
		}
	}

	if found {
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
	} else {
		fmt.Println("U heeft helaas geen toegang tot het parkeerterrein")
	}
}
