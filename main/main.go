package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/Icyvexen/rpgscripts/items"
	"github.com/Icyvexen/rpgscripts/names"
)

const menu string = `
Options:
1. Magic Items
2. Ship Names
3. Exit

Enter number for desired option... `

var read = bufio.NewReader(os.Stdin)

func main() {
	for {
		fmt.Println(menu)
		option := inputNum()
		switch option {
		case 1:
			fmt.Println("How many magic items to generate?")
			num := inputNum()
			for num > 0 {
				gen := items.CreateMagicItem()
				fmt.Println(gen)
				time.Sleep(time.Microsecond) //give the RNG call time to use nanosecond seed
				num--
			}
		case 2:
			generate("ship names", names.NewShipName)
		default:
			fmt.Println("Closing script...")
			os.Exit(0)
		}
	}
}

func input() string {
	in, err := read.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	//Because is input line, take any newline characters and remove them
	return strings.Replace(in, "\r\n", "", -1)
}

func generate(toGen string, fn func(...int64) string) {
	fmt.Println("How many", toGen, "to generate?")
	num := inputNum()
	for num > 0 {
		gen := fn
		ret := gen()
		fmt.Println(ret)
		time.Sleep(time.Microsecond) //give the RNG call time to use nanosecond seed
		num--
	}
}

func inputNum() int {
	var in string
	var numOut int
	for in == "" {
		in = input()
		if in == "" {
			fmt.Println("Empty input. Please enter a number.")
			continue
		}
		numIn, err := strconv.Atoi(in)
		if err != nil {
			fmt.Println("Error!\n", err)
			os.Exit(1)
		}
		if numIn <= 0 {
			fmt.Println("Number must be greater than 0.")
			in = ""
			continue
		}
		numOut = numIn
	}
	return numOut
}
