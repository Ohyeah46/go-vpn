package tlsconfig

import (
	"crypto/tls"
	"crypto/x509"
	"os"
)

func LoadTLSConfig(certFile, keyFile, caFile string, isServer bool) (*tls.Config, error) {
	cert, err := tls.LoadX509KeyPair(certFile, keyFile)
	if err != nil {
		return nil, err
	}

	caCert, err := os.ReadFile(caFile)
	if err != nil {
		return nil, err
	}

	caPool := x509.NewCertPool()
	if !caPool.AppendCertsFromPEM(caCert) {
		return nil, err
	}

	config := &tls.Config{
		Certificates: []tls.Certificate{cert},
		RootCAs:      caPool,      // <- здесь должен быть твой ca.crt, загруженный в caPool
		ServerName:   "localhost", // обязательно, чтобы имя совпадало с CN в сертификате сервера
	}

	if isServer {
		config.ClientCAs = caPool
		config.ClientAuth = tls.RequireAndVerifyClientCert
	} else {
		config.ServerName = "localhost"
		// На время отладки можно раскомментировать:
		// config.InsecureSkipVerify = true
	}

	return config, nil
}
