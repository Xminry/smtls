# smtls-gm

Shangmi(SM) TLS1.3

[![Go Reference](https://pkg.go.dev/badge/github.com/quic-go/smtls-go1-19.svg)](https://pkg.go.dev/github.com/quic-go/smtls-go1-19)
[![.github/workflows/go-test.yml](https://github.com/quic-go/smtls-go1-19/actions/workflows/go-test.yml/badge.svg)](https://github.com/quic-go/smtls-go1-19/actions/workflows/go-test.yml)

This repository contains a modified version of the standard library's TLS implementation, modified for the QUIC protocol. It is used by [quic-go](https://github.com/lucas-clemente/quic-go).

based on smtls, support sm2,sm3,sm4.
satisfy RFC8998
基于QTLS开发的国密TLS库
支持RFC8998标准
支持TLS1.3

gmtls1.3

a example for tls1.3 is in example-tls13
example-tls13为一个非国密的普通tls1.3示例

a example for gmtls1.3 without CA is in example-gmtls13-noCA
example-gmtls13-noCA为一个不包含CA验证的国密tls1.3示例


a example for gmtls1.3 with CA is in example-gmtls13-CA
example-gmtls13-CA为一个包含CA验证的国密tls1.3示例

