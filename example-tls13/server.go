package main

import (
	"crypto/tls"
	"flag"
	"io"
	"log"
	"net"
)

func main() {
	port := flag.String("port", "8360", "listening port")
	certFile := flag.String("cert", "testdata/example-cert.pem", "certificate PEM file")
	keyFile := flag.String("key", "testdata/example-key.pem", "key PEM file")
	flag.Parse()
	cert, err := tls.LoadX509KeyPair(*certFile, *keyFile)
	if err != nil {
		log.Fatal(err)
	}
	config := &tls.Config{Certificates: []tls.Certificate{cert}}
	log.Printf("listening on port %s\n", *port)
	l, err := tls.Listen("tcp", ":"+*port, config)
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("accepted connection from %s\n", conn.RemoteAddr())
		go func(c net.Conn) {
			io.Copy(c, c)
			c.Close()
			log.Printf("closing connection from %s\n", conn.RemoteAddr())
		}(conn)
	}

}
