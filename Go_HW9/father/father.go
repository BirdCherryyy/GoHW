package father

import (
	"flag"
	"fmt"
	"net/url"
)

/*
Port: 8080
DbURL: postgres://db-user:db-password@petstore-db:5432/petstore?sslmode=disable
JaegerURL: http://jaeger:16686
sentry_url: http://sentry:9000
kafka_broker: kafka:9092
somEAppID: testid
some_app_key: testkey
*/
type MyConf struct { //Структурки
	Port   *int
	EAppID *string
}
type MyURL struct {
	DbURL     *string
	JaegerURL *string
}

func (conf *MyConf) ReadConf() { //чтение одной структурки
	conf.Port = flag.Int("Port", 5432, "this is Port")
	conf.EAppID = flag.String("EAppID", "nothing", "this is EAppID")
	flag.Parse()
}
func (URL *MyURL) ReadURL() { //чтение другой структурки
	URL.DbURL = flag.String("postgres://db-user:5432", "nothing", "this is database URL")
	URL.JaegerURL = flag.String("http://jaeger:16686", "nothing", "this is jaeger URL")
	flag.Parse()
}

func (conf *MyConf) OutConf() { //вывод одной структурки
	fmt.Println("Port = ", *conf.Port)
	fmt.Println("EAppID = ", *conf.EAppID)
}
func (URL *MyURL) OutConf() { //вывод другой структурки
	fmt.Println("DbURL = ", *URL.DbURL)
	fmt.Println("JaegerURL = ", *URL.JaegerURL)
}

func (URL *MyURL) CheckURL() { //проверка моего неправильного URL
	defer func() {
		err := recover()
		fmt.Println("Oh, noy, this is panic, but we continue:)", err)
	}()

	_, err := url.Parse(*URL.DbURL)
	if err != nil {
		fmt.Println("URL ok")
	} else {
		panic(err)
	}
	/*
		_, err := url.Parse(*URL.JaegerURL)
		if err != nil {
			fmt.Println("URL ok")
		} else {
			fmt.Println("URL err")
		}
	*/
}
