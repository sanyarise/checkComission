package main

// Импортируем библиотеки вывода и конвертации
import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"unicode"
)

// Функция проверки комиссии
func check_commission(number string) uint64 {
	// Удаляем пробелы из строки
	numberWithoutSpaces := cleanSpaces(number)

	// Удаляем непечатаемые символы из строки
	number = cleanInvisibleSymbols(numberWithoutSpaces)

	// Проверяем, что строка не пустая
	if number == "" {
		// Если пустая, скажем об этом
		fmt.Println("Введено пустое значение")
		// И выйдем из функции
		return 0
	}

	// Преобразуем входящую строку в слайс рун
	runesNumber := []rune(number)

	// Итерируемся по слайсу рун
	for index, val := range runesNumber {
		// Если руна равна запятой (вместо точки часто используется запятая)
		if val == ',' {
			// Меняем ее значение в слайсе на руну точка
			runesNumber[index] = '.'
			// И меняем само значение на точку для дальнейшей проверки
			val = '.'
		}
		// Если руна равна точке
		if val == '.' {
			// И длина слайса после точки больше двух (то есть больше двух знаков после запятой)
			if len(runesNumber[index+1:]) > 2 {
				// Скажем об этом
				fmt.Println("Больше двух знаков после запятой")
				// И выйдем из функции
				return 0
			}
		}
	}
	// Преобразуем слайс рун обратно в строку
	number = string(runesNumber)
	// Конвертируем присланное значение из строки в число с запятой во временную переменную
	tempValue, err := strconv.ParseFloat(number, 64)
	// Если конвертация не получилась
	if err != nil {
		// Скажем об этом
		fmt.Println("Введено не число:", number)
		// И выйдем из функции
		return 0
	}
	// Если конвертация удалась, проверим полученное значение -
	// если оно меньше нуля или больше 99.99
	if tempValue < 0 || tempValue > 99.99 {
		// Скажем об этом
		fmt.Println("Введенное число меньше 0 или больше 99.99: ", tempValue)
		// И выйдем из функции
		return 0
	}
	// Умножим значение на 100, тем самым сдвинем запятую
	tempValue = tempValue * 100

	// Округляем полученное значение, т.к. некоторые числа
	// при преобразовании дают неверный результат, например 0.29
	// (0.29 * 100 = 28.999999999999996, что при приведении к uint64 дает 28, а не 29)
	tempValue = math.Round(tempValue)
	fmt.Printf("value after math round: %f\n", tempValue)
	// Преобразуем значение в нужный тип данных
	result := uint64(tempValue)
	// Возвращаем корректное значение
	fmt.Println("result is: \n", result)
	return result
}

// Удаляем все пробелы из строки
func cleanSpaces(s string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsSpace(r) {
			return -1
		}
		return r
	}, s)
}

// Очищаем строку от непечатаемых символов
func cleanInvisibleSymbols(s string) string {
	s = strings.TrimFunc(s, func(r rune) bool {
		return !unicode.IsGraphic(r)
	})
	return s
}

// Ввиду ограничений формы входящее значение от пользователя всегда в string
