package config

import (
	"flag"
	"log"
	"net"
	"os"
	"strconv"
)

type Config struct {
	IPAddress string
	Port      int
}

func GetConfig() Config {

	conf := GetFlags()
	if conf.IPAddress == "127.1.1.1" && conf.Port == 65535 {
		conf = GetEnv()
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
		log.Fatal("Error validating IP")
	}

	return conf

}
func GetEnv() Config {

	conf := Config{}
	conf.IPAddress = os.Getenv("MYSRV_IP_ADDRESS")
	conf.Port, _ = strconv.Atoi(os.Getenv("MYSRV_CONNECTION_PORT"))

	if net.ParseIP(conf.IPAddress) == nil {
		log.Fatal("Error validating IP")
	}

	return conf
}
