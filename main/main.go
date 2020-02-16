package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/Icyvexen/rpgscripts/rpgscripts"
)

const menu string = `
Options:
1. Magic Items
2. Exit 

Enter number for desired option... `

var read = bufio.NewReader(os.Stdin)

func main() {
	for {
		fmt.Println(menu)
		option := inputNum()
		switch option {
		case 1:
			fmt.Println("How many magic items to generate?")
			gen := inputNum()
			for gen > 0 {
				mi := rpgscripts.CreateMagicItem()
				fmt.Println(mi)
				gen--
			}
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

func inputNum() int {
	in := ""
	for in == "" {
		in = input()
	}
	num, err := strconv.Atoi(in)
	if err != nil {
		fmt.Println("Error!\n", err)
		os.Exit(1)
	}
	return num
}
