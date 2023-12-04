package GetStatus

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func getClient(caCertPath string, certPath string) *http.Client {

	// Load the certificate from disk.
	cacert, err := ioutil.ReadFile(certPath)
	if err != nil {
		log.Fatal(err)
	}

	cert, err := tls.LoadX509KeyPair(certPath, certPath)
	if err != nil {
		log.Fatal(err)
	}

	// Create a certificate pool and add the certificate to it.
	certPool := x509.NewCertPool()
	certPool.AppendCertsFromPEM(cacert)

	// Create a TLS configuration that uses the certificate pool.
	// Because the Nifi server uses a self-signed cert we skip the
	// client side verification of the server
	tlsConfig := &tls.Config{
		RootCAs:      certPool,
		Certificates: []tls.Certificate{cert},
	}

	// Create an HTTP transport that uses the TLS configuration.
	transport := &http.Transport{
		TLSClientConfig: tlsConfig,
	}

	// Create an HTTP client that uses the transport.
	return &http.Client{
		Transport: transport,
	}

}

func GetStatus(cacert string, cert string) string {

	//url := "https://apexetl-nifi-worker1.amd.com:8443/nifi-api/access/token"
	url := "https://apexetl-nifi-worker1.amd.com:8443/nifi-api/system-diagnostics"

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		fmt.Println("Error creating request: ", err)
		return "Error creating request: " + err.Error()
	}

	client := getClient(cacert, cert)

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request: ", err)
		return "Error sending request: " + err.Error()
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response: ", err)
		return "Error reading response: " + err.Error()
	}

	return string(body)

}
