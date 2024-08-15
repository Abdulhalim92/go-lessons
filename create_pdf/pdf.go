package main

import (
	"github.com/jung-kurt/gofpdf"
	"log"
)

func main() {
	// Создаем новый PDF-документ
	pdf := gofpdf.New("P", "mm", "A4", "")

	// Добавляем страницу
	pdf.AddPage()

	// Устанавливаем шрифт с поддержкой кириллицы
	pdf.AddUTF8Font("DejaVu", "", "./create_pdf/DejaVuSans.ttf")
	pdf.SetFont("DejaVu", "", 12)

	// Заголовок
	pdf.Cell(0, 10, "20 JSON Tasks in Go")
	pdf.Ln(12)

	// Задачи
	tasks := []string{
		"1. Сериализация структуры в JSON\nОписание: Напишите программу, которая сериализует структуру Person в формат JSON и выводит результат на экран.\nСтруктура:\n\ntype Person struct {\n    Name  string `json:\"name\"`\n    Age   int    `json:\"age\"`\n    Email string `json:\"email\"`\n}\n",
		"2. Десериализация JSON в структуру\nОписание: Напишите программу, которая десериализует JSON-строку в структуру Person и выводит результат на экран.\nСтруктура:\n\ntype Person struct {\n    Name  string `json:\"name\"`\n    Age   int    `json:\"age\"`\n    Email string `json:\"email\"`\n}\n",
		"3. Сериализация карты в JSON\nОписание: Напишите программу, которая сериализует карту map[string]int в JSON и выводит результат на экран.\nПример данных:\n\ndata := map[string]int{\n    \"apples\":  5,\n    \"oranges\": 10,\n}\n",
		"4. Десериализация JSON в карту\nОписание: Напишите программу, которая десериализует JSON-строку в карту map[string]int и выводит результат на экран.\nПример данных:\n{\n    \"apples\": 5,\n    \"oranges\": 10\n}\n",
		"5. Пропуск полей при сериализации\nОписание: Напишите программу, которая сериализует структуру, пропуская поля с пустыми значениями.\nСтруктура:\n\ntype Product struct {\n    Name     string  `json:\"name\"`\n    Price    float64 `json:\"price,omitempty\"`\n    InStock  bool    `json:\"in_stock,omitempty\"`\n}\n",
		"6. Работа с вложенными структурами\nОписание: Напишите программу, которая сериализует и десериализует структуру с вложенной структурой.\nСтруктура:\n\ntype Address struct {\n    Street string `json:\"street\"`\n    City   string `json:\"city\"`\n}\n\ntype Person struct {\n    Name    string  `json:\"name\"`\n    Age     int     `json:\"age\"`\n    Email   string  `json:\"email\"`\n    Address Address `json:\"address\"`\n}\n",
		"7. Использование JSON с интерфейсами\nОписание: Напишите программу, которая сериализует и десериализует карту map[string]interface{} и выводит результат на экран.\nПример данных:\n\ndata := map[string]interface{}{\n    \"name\":  \"John\",\n    \"age\":   30,\n    \"email\": \"john@example.com\",\n}\n",
		"8. Чтение JSON из файла\nОписание: Напишите программу, которая читает JSON из файла и десериализует его в структуру.\nФайл: person.json\nСтруктура:\n\ntype Person struct {\n    Name  string `json:\"name\"`\n    Age   int    `json:\"age\"`\n    Email string `json:\"email\"`\n}\n",
		"9. Запись JSON в файл\nОписание: Напишите программу, которая сериализует структуру в JSON и записывает результат в файл.\nФайл: output.json\nСтруктура:\n\ntype Product struct {\n    Name  string  `json:\"name\"`\n    Price float64 `json:\"price\"`\n}\n",
		"10. Пропуск неизвестных полей при десериализации\nОписание: Напишите программу, которая десериализует JSON-строку в структуру, игнорируя неизвестные поля.\nСтруктура:\n\ntype Person struct {\n    Name string `json:\"name\"`\n    Age  int    `json:\"age\"`\n}\n",
		"11. Форматирование JSON-вывода\nОписание: Напишите программу, которая сериализует структуру в JSON с отступами для улучшения читаемости.\nМетод: Используйте json.MarshalIndent.\nСтруктура:\n\ntype Person struct {\n    Name  string `json:\"name\"`\n    Age   int    `json:\"age\"`\n    Email string `json:\"email\"`\n}\n",
		"12. Обновление значения в JSON\nОписание: Напишите программу, которая десериализует JSON в структуру, изменяет одно из значений, а затем сериализует структуру обратно в JSON.\nСтруктура:\n\ntype Person struct {\n    Name  string `json:\"name\"`\n    Age   int    `json:\"age\"`\n    Email string `json:\"email\"`\n}\n",
		"13. Парсинг JSON с разными типами данных\nОписание: Напишите программу, которая десериализует JSON-строку, содержащую данные разных типов (строки, числа, массивы, объекты), в map[string]interface{}.\nПример данных:\n{\n    \"name\": \"John\",\n    \"age\": 30,\n    \"scores\": [100, 98, 95],\n    \"address\": {\"city\": \"New York\", \"street\": \"5th Avenue\"}\n}\n",
		"14. Массив структур в JSON\nОписание: Напишите программу, которая сериализует массив структур Person в JSON и выводит результат на экран.\nСтруктура:\n\ntype Person struct {\n    Name  string `json:\"name\"`\n    Age   int    `json:\"age\"`\n    Email string `json:\"email\"`\n}\n",
		"15. Добавление нового объекта в JSON-массив\nОписание: Напишите программу, которая читает JSON-массив из файла, добавляет новый объект и записывает обновленный массив обратно в файл.\nФайл: people.json\nСтруктура:\n\ntype Person struct {\n    Name  string `json:\"name\"`\n    Age   int    `json:\"age\"`\n    Email string `json:\"email\"`\n}\n",
		"16. Удаление объекта из JSON-массива\nОписание: Напишите программу, которая читает JSON-массив из файла, удаляет объект с определенным значением поля и записывает обновленный массив обратно в файл.\nФайл: people.json\nСтруктура:\n\ntype Person struct {\n    Name  string `json:\"name\"`\n    Age   int    `json:\"age\"`\n    Email string `json:\"email\"`\n}\n",
		"17. Поиск объекта в JSON-массиве\nОписание: Напишите программу, которая читает JSON-массив из файла и находит объект с определенным значением поля.\nФайл: people.json\nСтруктура:\n\ntype Person struct {\n    Name  string `json:\"name\"`\n    Age   int    `json:\"age\"`\n    Email string `json:\"email\"`\n}\n",
		"18. Подсчет объектов в JSON-массиве\nОписание: Напишите программу, которая читает JSON-массив из файла и подсчитывает количество объектов в массиве.\nФайл: people.json\nСтруктура:\n\ntype Person struct {\n    Name  string `json:\"name\"`\n    Age   int    `json:\"age\"`\n    Email string `json:\"email\"`\n}\n",
		"19. Десериализация JSON с произвольными полями\nОписание: Напишите программу, которая десериализует JSON-строку с произвольным набором полей в map[string]interface{}.\nПример данных:\n{\n    \"name\": \"John\",\n    \"age\": 30,\n    \"country\": \"USA\",\n    \"job\": \"Engineer\"\n}\n",
		"20. Сериализация и десериализация времени в JSON\nОписание: Напишите программу, которая сериализует и десериализует структуру с полем времени (time.Time) в JSON.\nСтруктура:\n\ntype Event struct {\n    Name      string    `json:\"name\"`\n    Timestamp time.Time `json:\"timestamp\"`\n}\n",
	}

	// Печать задач в PDF
	for _, task := range tasks {
		pdf.MultiCell(0, 10, task, "", "", false)
		pdf.Ln(4)
	}

	// Сохранение PDF-файла
	err := pdf.OutputFileAndClose("20_JSON_Tasks_in_Go.pdf")
	if err != nil {
		log.Fatalf("Error creating PDF file: %s\n", err)
	}
}
