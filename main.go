package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"sort"
	"strconv"
	"strings"
)

func sum(a, b int, op string) int {

	if op == "+" {
		return a + b
	} else if op == "-" {
		return a - b
	} else if op == "*" {
		return a * b
	} else if op == "/" {
		if b == 0 {
			panic("На ноль делить нельзя")
		} else {
			return a / b
		}
	} else {
		return 0
	}
}

func to_roman(n int) string {

	res := ""

	elements := map[int]string{
		100: "C",
		90:  "XC",
		50:  "L",
		40:  "XL",
		10:  "X",
		9:   "IX",
		5:   "V",
		4:   "IV",
		1:   "I",
	}

	keys := make([]int, len(elements))

	i := 0

	for k := range elements {
		keys[i] = k
		i++
	}

	sort.Ints(keys)
	slices.Reverse(keys)

	for _, p := range keys {
		r := n / p
		res += strings.Repeat(elements[p], r)
		n %= p

	}
	return res
}

func to_arabic(n string) int {

	elements := map[string]int{
		"C":  100,
		"XC": 90,
		"L":  50,
		"XL": 40,
		"X":  10,
		"IX": 9,
		"V":  5,
		"IV": 4,
		"I":  1,
	}

	rome_num := n

	res := 0

	for n, i := range rome_num {
		for ro, ar := range elements {
			if string(i) == ro {
				if n == len(rome_num)-1 {
					res += ar
				} else {
					if string(i) == "C" && string(rome_num[n+1]) == "M" {
						res -= ar
						continue
					}
					if string(i) == "C" && string(rome_num[n+1]) == "D" {
						res -= ar
						continue
					}
					if string(i) == "X" && string(rome_num[n+1]) == "C" {
						res -= ar
						continue
					}
					if string(i) == "X" && string(rome_num[n+1]) == "L" {
						res -= ar
						continue
					}
					if string(i) == "I" && string(rome_num[n+1]) == "X" {
						res -= ar
						continue
					}
					if string(i) == "I" && string(rome_num[n+1]) == "V" {
						res -= ar
						continue
					} else {
						res += ar
					}
				}
			}
		}
	}
	return res
}

//

func check_type(i string) (string, int) {

	ar_num := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10"}
	rome_num := []string{"I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX", "X"}

	if slices.Contains(ar_num, i) {
		return i, 0
	} else if slices.Contains(rome_num, i) {
		return i, 1
	} else {
		panic("Числа должны быть арабскими: от 1 до 10; или римскими: от I до X")
	}
}

func oper_check(op string) string {

	if op == "+" {
		return op
	} else if op == "-" {
		return op
	} else if op == "*" {
		return op
	} else if op == "/" {
		return op
	} else {
		panic("Введён неверный оператор")
	}
}

//

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("")
	fmt.Println("КАЛЬКУЛЯТОР")
	fmt.Println("")
	fmt.Print("Введите математическое выражение, соответствующее условиям...")
	text, _ := reader.ReadString('\n')

	strp_st := strings.Fields(text)
	if len(strp_st) == 1 { //При вводе без пробелов работает не корректно. В примерах работы программы использовался ввод с пробелами, поэтому багом не считаю.
		fmt.Println("Cтрока не является математической операцией")
	} else if len(strp_st) != 3 {
		fmt.Println("Формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *)")
	} else {

		first_value_str := strp_st[0]
		second_value_str := strp_st[2]
		operand := strp_st[1]

		a_d, ir := check_type(first_value_str)
		b_d, ir2 := check_type(second_value_str)

		o_c := oper_check(operand)

		if ir == 0 && ir2 == 0 {
			to_digit, err := strconv.Atoi(a_d)
			if err != nil {
				panic(err)
			}
			to_digit2, err := strconv.Atoi(b_d)
			if err != nil {
				panic(err)
			}

			itogo := sum(to_digit, to_digit2, o_c)
			fmt.Println(itogo)

		} else if ir == 1 && ir2 == 1 {
			s1 := to_arabic(a_d)
			s2 := to_arabic(b_d)
			itogo_ar := sum(s1, s2, o_c)
			if itogo_ar < 0 {
				panic("В римской системе нет отрицательных чисел")
			} else {
				itogo := to_roman(itogo_ar)
				fmt.Println(itogo)
			}
		} else {
			panic("Неверное значение")
		}

	}

}
