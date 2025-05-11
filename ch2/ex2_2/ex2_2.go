/*
Program to convert given value from on unit to another.
Usage:
	go run ex2_2.go -convT convertion_type arg_1 arg_2 ... arg_n
Convertion types are:
	-CToF: celsisus to fahrenheit;
	-FToC: fahrenheit to celsius;
	-CToK: celsius to kelvin;
	-KToC: kelvin to celsius;
	-FToK: fahrenheit to kelvin;
	-KToF: kelvin to fahrenheit;
	-MToFt: meter to feet;
	-FtToM: feet to meter;
	-KgToLb: kilogram to poud;
	-LbToKg: pound to kilogram.
*/

package main

import (
	"flag"
	"fmt"
	"github.com/RickSanch3z/ch2/ex2_2/conv"
	"bufio"
	"os"
	"strings"
)

func main() {
	var convType = flag.String("convT", "", "Convertion types:\n" +
							"\tCToF: celsisus to fahrenheit;\n" +
							"\tFToC: fahrenheit to celsius;\n" +
							"\tCToK: celsius to kelvin;\n" +
							"\tKToC: kelvin to celsius;\n" +
							"\tFToK: fahrenheit to kelvin;\n" +
							"\tKToF: kelvin to fahrenheit;\n" +
							"\tMToFt: meter to feet;\n" +
							"\tFtToM: feet to meter;\n" +
							"\tKgToLb: kilogram to poud;\n" +
							"\tLbToKg: pound to kilogram.\n",
						)
	flag.Parse()

	nFlag := flag.NFlag()
	nArg := flag.NArg()

	if nFlag == 1 {
		err := conv.ValidConv(*convType)
		if err != nil {
			flag.PrintDefaults()
		} else {
			if (nArg > 0) {
				for _, arg := range flag.Args() {
					retVal, err := conv.ConvertUnit(*convType, arg)
					if err != nil {
						fmt.Printf("%s\n", err)
					} else {
						fmt.Printf("%s\n", retVal)
					}
				}
			} else {
				fmt.Println("To exit input mode write \"EOF\" and press \"Enter\".")
				input := bufio.NewScanner(os.Stdin)
				for input.Scan() {
					line_in := strings.TrimSpace(input.Text())
					if line_in == "EOF" {
						break
					} else {
						retVal, err := conv.ConvertUnit(*convType, line_in)
						if err != nil {
							fmt.Printf("%s\n", err)
						} else {
							fmt.Printf("%s\n", retVal)
						}
					}
				}
			}
		}
	}
}
