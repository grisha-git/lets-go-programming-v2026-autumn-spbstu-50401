package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	configPath := flag.String("config", "config.yaml", "Path to config file")
	flag.Parse()

	cfg, err := config.Load(*configPath)
	if err != nil {
		// По условию задания при ошибке — аварийное завершение (panic)
		panic(fmt.Sprintf("Ошибка загрузки конфигурации: %v", err))
	}

	// 3. Чтение входного XML файла
	xmlData, err := os.ReadFile(cfg.InputFile)
	if err != nil {
		panic(fmt.Sprintf("Ошибка чтения входного файла '%s': %v", cfg.InputFile, err))
	}

	// 4. Вызов парсера
	// Передаем сырые байты, получаем готовый отсортированный слайс
	currencies, err := parser.ParseAndSortXML(xmlData)
	if err != nil {
		panic(fmt.Sprintf("Ошибка парсинга XML: %v", err))
	}

	// 5. Подготовка выходного файла (Шаг 1, Примечание 2)
	// Получаем путь к папке (например, "result/")
	outputDir := filepath.Dir(cfg.OutputFile)

	// Создаем папку, если её не существует (0755 - стандартные права доступа)
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		panic(fmt.Sprintf("Ошибка создания директории '%s': %v", outputDir, err))
	}

	// 6. Кодирование данных в JSON (Шаг 4)
	// MarshalIndent создает красивый JSON с отступами в 4 пробела
	jsonData, err := json.MarshalIndent(currencies, "", "    ")
	if err != nil {
		panic(fmt.Sprintf("Ошибка кодирования в JSON: %v", err))
	}

	// 7. Запись результата в файл
	// os.WriteFile создаст файл, если его нет, или перезапишет, если он есть
	err = os.WriteFile(cfg.OutputFile, jsonData, 0644)
	if err != nil {
		panic(fmt.Sprintf("Ошибка записи в файл '%s': %v", cfg.OutputFile, err))
	}

	// 8. Успешное завершение
	fmt.Printf("Готово! Данные успешно сохранены в %s\n", cfg.OutputFile)
}
