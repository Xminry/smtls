package main

import (
	"flag"
	"github.com/tjfoc/gmsm/gmtls"
	"io"
	"log"
	"net"
)

func main() {
	port := flag.String("port", "8360", "listening port")
	certFile := flag.String("cert", "testdata/gm-example-cert.pem", "certificate PEM file")
	keyFile := flag.String("key", "testdata/gm-example-key.pem", "key PEM file")
	flag.Parse()
	cert, err := gmtls.LoadX509KeyPair(*certFile, *keyFile)
	if err != nil {
		log.Fatal(err)
	}

	if err != nil {
		log.Fatal(err)
	}
	config := &gmtls.Config{Certificates: []gmtls.Certificate{cert},
		CipherSuites: []uint16{gmtls.GMTLS_ECDHE_SM2_WITH_SM4_SM3},
		MinVersion:   gmtls.VersionTLS12}
	log.Printf("listening on port %s\n", *port)
	l, err := gmtls.Listen("tcp", ":"+*port, config)
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
