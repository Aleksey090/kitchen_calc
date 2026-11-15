package main

import "fmt"

func main() {
	// 1. Создайте map[string]float64 с остатками продуктов
	stock := map[string]float64{
		"Salmon":  15.5,
		"Avocado": 4.0,
		"Rice":    25.0,
		"Cheese":  6.5,
	}
	// 2. Заведите переменную для подсчета общего веса (var totalWeight float64)
	var totalWeight float64
	// 3. Пройдитесь по карте циклом for name, weight := range ...
	//    - Проверьте, если weight < 5.0 -> выведите предупреждение
	//    - Прибавьте weight к totalWeight
	for name, weight := range stock {
		if weight < 5.0 {
			fmt.Println("Need to buy:", name)
		}
		totalWeight += weight
	}
	// 4. Выведите итоговый вес всех продуктов
	fmt.Println("Total product weight:", totalWeight)
}
