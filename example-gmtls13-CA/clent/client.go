package main

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/pem"
	"flag"
	"fmt"
	"github.com/Xminry/smtls"
	"github.com/emmansun/gmsm/smx509"
	"io"
	"log"
	"os"
)

func main() {
	port := flag.String("port", "8360", "port to connect")
	certFile := flag.String("certfile", "testdata/gmca.cert.pem", "trusted CA certificate")
	flag.Parse()
	certPem, err := os.ReadFile(*certFile)
	if err != nil {
		log.Fatal(err)
	}
	certDERBlock, _ := pem.Decode(certPem)

	certl, err := smx509.ParseCertificates(certDERBlock.Bytes)
	if err != nil {
		log.Fatal(err)
	}
	cert := certl[0]
	certPool := x509.NewCertPool()
	var exkusg []x509.ExtKeyUsage
	for ii := 0; ii < len(cert.ExtKeyUsage); ii++ {
		exkusg = append(exkusg, x509.ExtKeyUsage(cert.ExtKeyUsage[ii]))
	}
	var cert1 x509.Certificate = x509.Certificate{
		Raw:                         cert.Raw,
		RawTBSCertificate:           cert.RawTBSCertificate,
		RawSubjectPublicKeyInfo:     cert.RawSubjectPublicKeyInfo,
		RawSubject:                  cert.RawSubject,
		RawIssuer:                   cert.RawIssuer,
		Signature:                   cert.Signature,
		SignatureAlgorithm:          x509.SignatureAlgorithm(cert.SignatureAlgorithm),
		PublicKeyAlgorithm:          x509.PublicKeyAlgorithm(cert.PublicKeyAlgorithm),
		PublicKey:                   cert.PublicKey,
		Version:                     cert.Version,
		SerialNumber:                cert.SerialNumber,
		Issuer:                      cert.Issuer,
		Subject:                     cert.Subject,
		NotBefore:                   cert.NotBefore,
		NotAfter:                    cert.NotAfter,
		KeyUsage:                    x509.KeyUsage(cert.KeyUsage),
		Extensions:                  cert.Extensions,
		ExtraExtensions:             cert.ExtraExtensions,
		UnhandledCriticalExtensions: cert.UnhandledCriticalExtensions,
		ExtKeyUsage:                 exkusg,
		UnknownExtKeyUsage:          cert.UnknownExtKeyUsage,
		BasicConstraintsValid:       cert.BasicConstraintsValid,
		IsCA:                        cert.IsCA,
		MaxPathLen:                  cert.MaxPathLen,
		MaxPathLenZero:              cert.MaxPathLenZero,
		SubjectKeyId:                cert.SubjectKeyId,
		AuthorityKeyId:              cert.AuthorityKeyId,
		OCSPServer:                  cert.OCSPServer,
		IssuingCertificateURL:       cert.IssuingCertificateURL,
		DNSNames:                    cert.DNSNames,
		EmailAddresses:              cert.EmailAddresses,
		IPAddresses:                 cert.IPAddresses,
		URIs:                        cert.URIs,
		PermittedDNSDomainsCritical: cert.PermittedDNSDomainsCritical,
		PermittedDNSDomains:         cert.PermittedDNSDomains,
		ExcludedDNSDomains:          cert.ExcludedDNSDomains,
		PermittedIPRanges:           cert.PermittedIPRanges,
		ExcludedIPRanges:            cert.ExcludedIPRanges,
		PermittedEmailAddresses:     cert.PermittedEmailAddresses,
		ExcludedEmailAddresses:      cert.ExcludedEmailAddresses,
		PermittedURIDomains:         cert.PermittedURIDomains,
		ExcludedURIDomains:          cert.ExcludedURIDomains,
		CRLDistributionPoints:       cert.CRLDistributionPoints,
		PolicyIdentifiers:           cert.PolicyIdentifiers,
	}
	certPool.AddCert(&cert1)

	config := &tls.Config{
		//InsecureSkipVerify: true,
		RootCAs:          certPool,
		CipherSuites:     []uint16{smtls.TLS_SM4_GCM_SM3},
		MinVersion:       smtls.VersionTLS13,
		CurvePreferences: []smtls.CurveID{smtls.CurveSM2}}
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
