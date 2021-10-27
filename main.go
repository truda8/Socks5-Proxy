package main

import (
	"flag"

	"github.com/armon/go-socks5"
)

type Config struct {
	host string
	port string
	user string
	pass string
}

func main() {
	config := getArgs()

	// Username and password auth config
	cred := socks5.StaticCredentials{
		config.user: config.pass,
	}
	auth := socks5.UserPassAuthenticator{Credentials: cred}

	// Create a SOCKS5 server
	conf := &socks5.Config{AuthMethods: []socks5.Authenticator{auth}}
	server, err := socks5.New(conf)
	if err != nil {
		panic(err)
	}

	// Create SOCKS5 proxy
	if err := server.ListenAndServe("tcp", config.host+":"+config.port); err != nil {
		panic(err)
	}
}

func getArgs() Config {
	var config Config
	flag.StringVar(&config.host, "host", "127.0.0.1", "Listen host.")
	flag.StringVar(&config.port, "port", "1080", "Listen port.")
	flag.StringVar(&config.user, "user", "admin", "Socks5 username.")
	flag.StringVar(&config.pass, "pass", "admin", "Socks5 password.")
	flag.Parse()
	return config
}
