package father

import (
	"flag"
	"fmt"
	"net/url"
)

/*
port: 8080
db_url: postgres://db-user:db-password@petstore-db:5432/petstore?sslmode=disable
jaeger_url: http://jaeger:16686
sentry_url: http://sentry:9000
kafka_broker: kafka:9092
some_app_id: testid
some_app_key: testkey
*/
type MyConf struct {
	port     *int
	e_app_id *string
}
type MyURL struct {
	db_url     *string
	jaeger_url *string
}

func (conf *MyConf) ReadConf() {
	conf.port = flag.Int("port", 5432, "this is port")
	conf.e_app_id = flag.String("e_app_id", "nothing", "this is e_app_id")
	flag.Parse()
}
func (URL *MyURL) ReadURL() {
	URL.db_url = flag.String("URL1", "nothing", "this is database URL")
	URL.jaeger_url = flag.String("URL2", "nothing", "this is jaeger URL")
	flag.Parse()
}

func (conf *MyConf) OutConf() {
	fmt.Println("port = ", *conf.port)
	fmt.Println("e_app_id = ", *conf.e_app_id)
}
func (URL *MyURL) OutConf() {
	fmt.Println("db_url = ", *URL.db_url)
	fmt.Println("jaeger_url = ", *URL.jaeger_url)
}

func (URL *MyURL) CheckURL() {
	_, err := url.Parse(*URL.db_url)
	if err != nil {
		fmt.Println("URL ok")
	} else {
		fmt.Println("URL err")
	}
	/*
		_, err := url.Parse(*URL.jaeger_url)
		if err != nil {
			fmt.Println("URL ok")
		} else {
			fmt.Println("URL err")
		}
	*/
}
