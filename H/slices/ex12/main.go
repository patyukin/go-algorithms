package main

import (
	"fmt"
)

func main() {
	planets := []string{
		"Меркурий", "Венера", "Земля", "Марс",
		"Юпитер", "Сатурн", "Уран", "Нептун",
	}

	worlds := planets[0:4:5]
	worlds = append(worlds, "Марс2")
	worlds = append(worlds, "Марс3")

	fmt.Println("planets", planets)
	fmt.Println("worlds", worlds)
}
