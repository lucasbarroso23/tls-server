package main

import (
	"crypto/tls"
	"crypto/x509"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	_, err := BuildServerTLSConfig(certPath)
	if err != nil {
		println(err.Error())
	}

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("serverTLS")
	})

	server := http.Server{
		Addr:    ":8000",
		Handler: router,
	}

	fmt.Println(server.ListenAndServeTLS(certPath, keyPath))

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
