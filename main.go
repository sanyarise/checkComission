package main

// Импортируем библиотеки вывода и конвертации
import (
	"fmt"
	"strconv"
)

// Функция проверки комиссии
func check_commission(number string) uint64 {
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
	// Преобразуем значение в нужный тип данных
	result := uint64(tempValue)
	// Возвращаем корректное значение
	return result
}

// Ввиду ограничений формы входящее значение от пользователя всегда в string
// Пеменная input будет эмулировать входящее значение
func main() {
	// Проверяем маленькое значение
	input := "0.01"
	fmt.Println("Первое значение:", input)
	fmt.Println("Преобразованное значение:", check_commission(input))

	// Проверяем большое значение
	input = "99.99"
	fmt.Println("\nВторое значение:", input)
	fmt.Println("Преобразованное значение:", check_commission(input))

	// Проверяем если ввели не число
	input = "неЧисло"
	fmt.Println("\nСтроковое значение:", input)
	fmt.Println("Преобразованное значение:", check_commission(input))
}
