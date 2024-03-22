package main

import (
	"fmt"
)

func main() {
	planets := []string{"Меркурий", "Венера", "Земля", "Марс", "Юпитер", "Сатурн", "Уран", "Нептун"}

	terrestrial := planets[0:4:4] // Длина 4, вместимость 4
	worlds := append(terrestrial, "Церера")
	terrestrial[0] = "Меркурий2"

	fmt.Println(planets, terrestrial, worlds)
}
