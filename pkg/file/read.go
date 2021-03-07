package file

import (
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"io/ioutil"
)

// ReadPEM read the file at path and returns the certificate chain inside.
func ReadPEM(path string) ([]*x509.Certificate, error) {

	fileBytes, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	certChain, err := ParsePEM(fileBytes)
	if err != nil {
		return nil, err
	}

	if certChain == nil {
		return nil, errors.New("Error parsing certificate")
	}
	fmt.Printf("Read %d certificate(s) from file. \n", len(certChain))
	return certChain, nil
}

// ParsePEM steps through a byte array and returns a slice of certificates.
func ParsePEM(pemBytes []byte) ([]*x509.Certificate, error) {
	var certchain []*x509.Certificate

	// Step through pemBytes decoding PEM certificate until there's nothing left.
	for {
		block, rest := pem.Decode(pemBytes)
		if block == nil {
			// When no PEM data is found block is nil, all PEM data has been read.
			break
		}

		cert, err := x509.ParseCertificate(block.Bytes)
		if err != nil {
			return nil, err
		}
		certchain = append(certchain, cert)
		pemBytes = rest
	}

	return certchain, nil
}
