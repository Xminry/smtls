package main

import (
	"crypto/tls"
	"flag"
	"fmt"
	"github.com/quic-go/qtls-go1-19"
	"io"
	"log"
	"net"
)

func main() {
	port := flag.String("port", "8360", "listening port")
	certFile := flag.String("cert", "testdata/gm-example-cert.pem", "certificate PEM file")
	keyFile := flag.String("key", "testdata/gm-example-key.pem", "key PEM file")
	flag.Parse()
	cert9 := *certFile
	key9 := *keyFile
	fmt.Println(cert9, key9)
	cert, err := qtls.LoadX509KeyPair(*certFile, *keyFile)
	if err != nil {
		log.Fatal(err)
	}

	config := &qtls.Config{
		Certificates:     []tls.Certificate{cert},
		CipherSuites:     []uint16{qtls.TLS_SM4_GCM_SM3},
		MinVersion:       qtls.VersionTLS13,
		CurvePreferences: []qtls.CurveID{qtls.CurveSM2},
	}
	log.Printf("listening on port %s\n", *port)
	l, err := qtls.Listen("tcp", ":"+*port, config, nil)
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
			_, err := io.Copy(c, c)
			if err != nil {
				log.Println(err)
			}
			err = c.Close()
			if err != nil {
				log.Println(err)
			}
			log.Printf("closing connection from %s\n", conn.RemoteAddr())
		}(conn)
	}
}
