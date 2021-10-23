package config

import (
	"encoding/json"
	"flag"
	"log"
	"net"
	"os"
	"strconv"
)

const filename = "../conf/config.json"

type Config struct {
	IPAddress string `json:"ip_address"`
	Port      int    `json:"port"`
}

func GetConfig() Config {

	conf := ReadConfigFile()
	if conf.IPAddress == "127.2.2.2" {
		conf = GetFlags()
		if conf.IPAddress == "127.1.1.1" && conf.Port == 65535 {
			conf = GetEnv()
		}
	}
	return conf
}

func GetFlags() Config {

	var ipaddr string

	flag.StringVar(&ipaddr, "ipaddr", "127.1.1.1", "server ip address")
	port := flag.Int("port", 65535, "server port")

	flag.Parse()

	conf := Config{}

	conf.IPAddress = ipaddr
	conf.Port = *port

	if net.ParseIP(conf.IPAddress) == nil {
		log.Fatal("Error validating IP from FLAGS")
	}

	return conf

}
func GetEnv() Config {

	conf := Config{}
	conf.IPAddress = os.Getenv("MYSRV_IP_ADDRESS")
	conf.Port, _ = strconv.Atoi(os.Getenv("MYSRV_CONNECTION_PORT"))

	if net.ParseIP(conf.IPAddress) == nil {
		log.Fatal("Error validating IP from OS ENV")
	}

	return conf
}

func ReadConfigFile() Config {
	conf := Config{}

	if _, errf := os.Stat(filename); errf != nil {
		conf.IPAddress = "127.2.2.2"
		return conf
	}

	file, err := os.Open(filename)

	if err != nil {
		log.Fatal(err)
	}

	err = json.NewDecoder(file).Decode(&conf)
	if err != nil {
		log.Printf("Не могу декодировать json-файл в структуру: %v", err)
	}

	defer func() {
		err = file.Close()

		if err != nil {
			log.Fatal(err)
		}
	}()

	return conf
}
