package actioninfo

import (
	"log"
)

type DataParser interface {
	Parse(datastring string) error
	ActionInfo() (string, error)
}

func Info(dataset []string, dp DataParser) {
	for _, data := range dataset {
		if err := dp.Parse(data); err != nil {
			log.Println("Ошибка парсинга:", err)
			continue
		}
		info, err := dp.ActionInfo()
		if err != nil {
			log.Println("Ошибка получения информации:", err)
			continue
		}
		log.Println(info)
	}
}
