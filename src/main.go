package main

import (
	"bufio"
	"eletronic-distribution/src/core"
	"fmt"
	"os"
	"strconv"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Eletronic Distribution by Zackwn")
	fmt.Println("--------------------------------")

	for {
		fmt.Print("e: ")
		text, _ := reader.ReadString('\n')
		text = text[:len(text)-2] // remove "\n"

		textInt, err := strconv.Atoi(text)
		if err != nil {
			fmt.Println("Electons should be a number.")
			continue
		}
		electrons, valid := core.ValidateElectrons(textInt)
		if valid == false {
			fmt.Println("Invalid electrons.")
			continue
		}

		result := core.Run(electrons)

		fmt.Printf("\n%v\n", result)
	}
}
