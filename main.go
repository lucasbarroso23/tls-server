package main

import (
	"crypto/tls"
	"crypto/x509"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	goTLS "github.com/lucasbarroso23/poc-https/goTls"
)

func main() {
	addr := flag.String("addr", ":8000", "HTTPS network address")
	certFile := flag.String("certfile", certPath, "certificate PEM file")
	keyFile := flag.String("keyfile", keyPath, "key PEM file")
	flag.Parse()

	goTLS.GenerateCA()

	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
		}
		fmt.Fprintf(w, "Proudly served with GO and HTTPS!")
	})

	tlsConf, err := BuildServerTLSConfig(certPath)
	if err != nil {
		println(err.Error())
	}

	srv := http.Server{
		Addr:      *addr,
		Handler:   mux,
		TLSConfig: tlsConf,
	}

	log.Printf("Starting server on %s", *addr)
	err = srv.ListenAndServeTLS(*certFile, *keyFile)
	log.Fatal(err)

}

// RegisterTLSConfig registers a custom tls.Config
func BuildServerTLSConfig(certFilePath string) (*tls.Config, error) {
	pool := x509.NewCertPool()
	pem, err := ioutil.ReadFile(certFilePath)
	if err != nil {
		return nil, errors.New("Error executing ioutil.ReadFile in RegisterTLSConfig function")
	}

	if ok := pool.AppendCertsFromPEM(pem); !ok {
		return nil, errors.New("Failed to append PEM on RegisterTLSConfig")
	}

	return &tls.Config{
		RootCAs:            pool,
		MinVersion:         tls.VersionTLS12,
		InsecureSkipVerify: false,
	}, nil
}
