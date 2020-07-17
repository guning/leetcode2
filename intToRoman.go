package main

import (
	"fmt"
)

func intToRoman(num int) string {

	one := []string{"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"}
	ten := []string{"", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"}
	hundred := []string{"",  "C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"}
	thousand := []string{"", "M", "MM", "MMM"}


	return thousand[num/1000] + hundred[(num%1000)/100] + ten[(num%100)/10] + one[num%10]
}

func main() {
	fmt.Println(intToRoman(1994))
}
