package main

import (
	father "family/father"
)

func main() {
	Conf := new(father.MyConf)
	URL := new(father.MyURL)
	Conf.ReadConf()
	URL.ReadURL()
	URL.CheckURL()
	Conf.OutConf()
}
