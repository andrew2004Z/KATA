package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Decode(roman string) int {
	var sum int
	var Roman = map[byte] int {'I':1, 'V':5, 'X':10, 'L':50, 'C':100, 'D':500, 'M':1000,}
	for k, v := range roman {
	  if k < len(roman) - 1 && Roman[byte(roman[k + 1])] > Roman[byte(roman[k])] {
		sum -= Roman[byte(v)]
	  } else {
		sum += Roman[byte(v)]
	  }
	}
	return sum
  }

  func arabicToRoman(num int) string {
    roman := ""
    for num > 0 {
        switch {
        case num >= 1000:
            roman += "M"
            num -= 1000
        case num >= 900:
            roman += "CM"
            num -= 900
        case num >= 500:
            roman += "D"
            num -= 500
        case num >= 400:
            roman += "CD"
            num -= 400
        case num >= 100:
            roman += "C"
            num -= 100
        case num >= 90:
            roman += "XC"
            num -= 90
        case num >= 50:
            roman += "L"
            num -= 50
        case num >= 40:
            roman += "XL"
            num -= 40
        case num >= 10:
            roman += "X"
            num -= 10
        case num >= 9:
            roman += "IX"
            num -= 9
        case num >= 5:
            roman += "V"
            num -= 5
        case num >= 4:
            roman += "IV"
            num -= 4
        default:
            roman += "I"
            num -= 1
        }
    }
    return roman
}



func StrtoInt(str string) int {
	b, _ := strconv.Atoi(str)
	return b
}

func isNumber(str string) bool {
	 _, err := strconv.ParseFloat(str, 64) 
	return err == nil 
}
func main() {
	var example string
	//var result string
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	example = scanner.Text()
	choper := strings.Fields(example)
	if len(choper) > 3 {
		fmt.Println("Не правильный формат")
		return
	} else if len(choper) < 3{
		fmt.Println("Не математическая операция")
		return
	} 
	res1 := Decode(choper[0])
	res2 := Decode(choper[2])
	if res1 != 0 && res2 != 0 && res1 < res2{
		fmt.Println("Результат не может быть отрицательным")
		return
	} else if res1 == 0 && res2 != 0 || res2 == 0 && res1 != 0  {                                           
        fmt.Println("Ошибка используються разные системы исчесления")
		return
    }
	if isNumber(choper[0]) && isNumber(choper[2]) && StrtoInt(choper[0]) >= 1 && StrtoInt(choper[0]) <= 10 && StrtoInt(choper[2]) >= 1 && StrtoInt(choper[2]) <= 10 {
		if choper[1] == "+" {
			res := StrtoInt(choper[0]) + StrtoInt(choper[2])
			fmt.Println(res)
			return
		} else if choper[1] == "-"{
			res := StrtoInt(choper[0]) - StrtoInt(choper[2])
			fmt.Println(res)
			return
		}else if choper[1] == "/"{
			res := StrtoInt(choper[0]) / StrtoInt(choper[2])
			fmt.Println(res)
			return
		}else if choper[1] == "*"{
			res := StrtoInt(choper[0]) * StrtoInt(choper[2])
			fmt.Println(res)
			return
		}
	} else if isNumber(choper[0]) == false && isNumber(choper[2]) == false && Decode(choper[0]) >= 1 && Decode(choper[0]) <= 10 && Decode(choper[2]) >= 1 && Decode(choper[2]) <= 10 {
		res1 := Decode(choper[0])
		res2 := Decode(choper[2])
		if choper[1] == "+" {
			res := res1 + res2
			fmt.Println(arabicToRoman(res))
			return
		} else if choper[1] == "-"{
			res := res1 - res2
			if res > 0 {
				fmt.Println(arabicToRoman(res))
			}
			return
		}else if choper[1] == "/"{
			res := res1 / res2
			fmt.Println(arabicToRoman(res))
			return
		}else if choper[1] == "*"{
			res := res1 * res2
			fmt.Println(arabicToRoman(res))
			return
		}
	}
	if (StrtoInt(choper[0]) > 10 || StrtoInt(choper[0]) < 1 || StrtoInt(choper[2]) > 10 || StrtoInt(choper[2]) < 1) || (res1 != 0 && (res1 > 10 || res1 < 1)) || (res2 != 0 && (res2 > 10 || res2 < 1)) {
		fmt.Println("Введенные числа не подходят условию")
		return
	}
}
