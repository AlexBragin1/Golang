package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var romanNumbers = map[byte]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
}

// операция сложения
func add(op1, op2 int) int {

	return op1 + op2
}

// операция вычетания
func sub(op1, op2 int, operandsIsRoman bool) int {

	resultSub := op1 - op2
	if !operandsIsRoman && resultSub < 0 {
		return -101
	}
	return resultSub
}

// операция деления
func div(op1, op2 int, operandsIsRoman bool) int {
	resultDiv := op1 / op2
	if !operandsIsRoman && resultDiv < 1 {
		return -102
	}
	return resultDiv
}

// операция умножения
func multy(op1, op2 int) int {
	return op1 * op2
}

// функция проверяет является чило римсим
func isRoman(number string) bool {
	flag := true
	for i := range number {
		if romanNumbers[number[i]] == 0 {
			flag = false
			break
		}

	}
	return flag
}

// фунция переводит из целого числа в строку
func intToRoman(number int) string {
	value := []int{100, 90, 50, 40, 10, 9, 5, 4, 1}
	simbol := []string{"C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

	strRoman := ""

	for i := 0; i < len(value); i++ {
		for number >= value[i] {
			strRoman += simbol[i]
			number -= value[i]
		}
	}
	return strRoman
}

// функция перевода из римского числа в целое
func romanToInt(strRoman string) int {
	var romanStrToInt int
	for i := 0; i < len(strRoman); i++ {
		if i+1 < len(strRoman) && romanNumbers[strRoman[i]] < romanNumbers[strRoman[i+1]] {
			romanStrToInt -= romanNumbers[strRoman[i]]
		} else {
			romanStrToInt += romanNumbers[strRoman[i]]
		}

	}
	return romanStrToInt
}

// перевод строки в целое число
func strToInt(digits string) (int, error) {
	num, err := strconv.Atoi(digits)
	if err != nil {
		return -103, err
	}
	return num, err
}

// проверка является число
func isArabic(numbers string) bool {
	_, err := strToInt(numbers)
	if err != nil {
		return false
	}
	return true
}

// функция проверяет ограничений  на операнды
func limitsNumber(num int) bool {
	if num < 1 || num > 10 {
		return false
	}
	return true
}

// функция ищет находиткакую математическую операцию надо совершить  над операндами и подсчитывает количество операций
func operationWithOperands(exp string) (int, byte) {
	countOperation := 0
	iOperation := -1
	var opertationWithOp byte
	for i := 0; i < len(exp); i++ {
		if exp[i] == '-' || exp[i] == '+' || exp[i] == '*' || exp[i] == '/' {
			iOperation = i
			opertationWithOp = exp[i]
			countOperation++
		}
	}
	if countOperation > 1 {
		iOperation = -2
	}
	return iOperation, opertationWithOp
}

func main() {
	var result int
	numberIsArabic := false

	expressionToTrim, err := bufio.NewReader(os.Stdin).ReadString('\n')

	if err != nil {
		fmt.Println(err)
	}
	// удаляем переход на другую строку
	expressionTrim := strings.Trim(expressionToTrim, "\n")
	//убираем все пробелы в нашей строке
	expression := strings.Replace(expressionTrim, " ", "", -1)

	indexOperation, operation := operationWithOperands(expression)

	if indexOperation != -1 && indexOperation != -2 {
		//из выражения берем операнды и  переводим вверхний регистр
		operandOneStr := strings.ToUpper(expression[:indexOperation])
		operandTwoStr := strings.ToUpper(expression[indexOperation+1:])

		if isArabic(operandOneStr) == isArabic(operandTwoStr) {
			numberIsArabic = isArabic(operandOneStr)
			operandOneInt := 0
			operandTwoInt := 0
			if numberIsArabic {
				operandOneInt, _ = strToInt(operandOneStr)
				operandTwoInt, _ = strToInt(operandTwoStr)
			} else if !numberIsArabic && isRoman(operandOneStr) {
				operandOneInt = romanToInt(operandOneStr)
				operandTwoInt = romanToInt(operandTwoStr)

			} else {
				fmt.Println("Вывод ошибки, так как используются одновременно разные системы счисления.")
			}
			if limitsNumber(operandOneInt) && limitsNumber(operandTwoInt) {
				switch operation {
				case '+':
					result = add(operandOneInt, operandTwoInt)
				case '-':
					result = sub(operandOneInt, operandTwoInt, numberIsArabic)
				case '/':
					result = div(operandOneInt, operandTwoInt, numberIsArabic)
				case '*':
					result = multy(operandOneInt, operandTwoInt)
				}

				if !numberIsArabic && result == -101 {
					fmt.Println("Вывод ошибки, так как в римской системе нет отрицательных чисел")
				} else if !numberIsArabic && result == -102 {
					fmt.Println(0)
					fmt.Println("Исключение, так как результат от деления меньше 1 для римской системе")
				} else if numberIsArabic {
					fmt.Println(result)
				} else {
					fmt.Println(intToRoman(result))
				}

			} else {
				fmt.Println("Вывод ошибки, Числа должны быть в пределах от 1 до 10")
			}

		} else {
			fmt.Println("Вывод ошибки, так как используются одновременно разные системы счисления.")
		}
	}
	if indexOperation == -1 {
		fmt.Println("Вывод ошибки, так как строка не является математической операцией.")
	}
	if indexOperation == -2 {
		fmt.Println("Вывод ошибки, так как формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *).")
	}

}
