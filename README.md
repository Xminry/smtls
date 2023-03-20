# qtls-gm

[![Go Reference](https://pkg.go.dev/badge/github.com/quic-go/qtls-go1-19.svg)](https://pkg.go.dev/github.com/quic-go/qtls-go1-19)
[![.github/workflows/go-test.yml](https://github.com/quic-go/qtls-go1-19/actions/workflows/go-test.yml/badge.svg)](https://github.com/quic-go/qtls-go1-19/actions/workflows/go-test.yml)

This repository contains a modified version of the standard library's TLS implementation, modified for the QUIC protocol. It is used by [quic-go](https://github.com/lucas-clemente/quic-go).

based on qtls, support sm2,sm3,sm4.
satisfy RFC8998

gmtls1.3

a example for tls1.3 is in example-tls13
a example for gmtls1.3 without CA is in example-gmtls13-noCA
a example for gmtls1.3 with CA is in example-gmtls13-CA

