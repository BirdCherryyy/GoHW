package jconf

import (
	"encoding/json"
	father "family/father"
	"os"
)

func ReadJ(MyConf *father.MyConf, MyURL *father.MyURL) { //считывание сразу 2ух структур
	data, err := os.Open("./JSON/MyConf.json") //открываем и закрываем один файл
	if err != nil {
		panic(err)
	}
	defer data.Close()

	if err = json.NewDecoder(data).Decode(&MyConf); err != nil { //передаем данные
		panic(err)
	}
	data, err = os.Open("./JSON/MyURL.json") //открываем и закрываем другой файл
	if err != nil {
		panic(err)
	}
	defer data.Close()

	if err = json.NewDecoder(data).Decode(&MyURL); err != nil { //передаем данные
		panic(err)
	}
}
