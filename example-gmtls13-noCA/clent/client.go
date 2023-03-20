package main

import (
	"crypto/tls"
	"github.com/emmansun/gmsm/smx509"

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
	certPool := smx509.NewCertPool()
	if ok := certPool.AppendCertsFromPEM(cert); !ok {
		log.Fatalf("unable to parse cert from %s", *certFile)
	}

	config := &tls.Config{
		InsecureSkipVerify: true,
		CipherSuites:       []uint16{smtls.TLS_SM4_GCM_SM3},
		MinVersion:         smtls.VersionTLS12,
		CurvePreferences:   []smtls.CurveID{smtls.CurveSM2}}
	conn, err := smtls.Dial("tcp", "localhost:"+*port, config, nil)
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
