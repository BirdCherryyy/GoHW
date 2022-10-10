package main

import (
	father "family/father"
	jconf "family/j_conf"
)

func main() {
	Conf := new(father.MyConf) //Инициализация структур
	URL := new(father.MyURL)
	Conf.ReadConf()        //считывание структуры conf
	URL.ReadURL()          //считывание структуры URL
	URL.CheckURL()         //проверка URL
	Conf.OutConf()         //Вывод значений структур
	jconf.ReadJ(Conf, URL) // Считывание с JSON
	URL.CheckURL()         //проверка URL
	Conf.OutConf()         //Вывод значений структур
}
