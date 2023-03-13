package main

import (
	"github.com/tjfoc/gmsm/gmtls"
	"github.com/tjfoc/gmsm/x509"

	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	port := flag.String("port", "8360", "port to connect")
	certFile := flag.String("certfile", "testdata/gm-example-cert.pem", "trusted CA certificate")
	flag.Parse()
	cert, err := os.ReadFile(*certFile)
	if err != nil {
		log.Fatal(err)
	}
	certPool := x509.NewCertPool()
	if ok := certPool.AppendCertsFromPEM(cert); !ok {
		log.Fatalf("unable to parse cert from %s", *certFile)
	}

	config := &gmtls.Config{InsecureSkipVerify: true, CipherSuites: []uint16{gmtls.GMTLS_ECDHE_SM2_WITH_SM4_SM3},
		MinVersion: gmtls.VersionTLS12}
	conn, err := gmtls.Dial("tcp", "localhost:"+*port, config)
	if err != nil {
		log.Fatal(err)
	}
	_, err = io.WriteString(conn, "Hello simple secure Server\n")
	if err != nil {
		log.Fatal("client write error:", err)
	}
	if err = conn.CloseWrite(); err != nil {
		log.Fatal(err)
	}
	buf := make([]byte, 256)
	n, err := conn.Read(buf)
	if err != nil && err != io.EOF {
		log.Fatal(err)
	}
	fmt.Println("client read:", string(buf[:n]))
	conn.Close()
}
