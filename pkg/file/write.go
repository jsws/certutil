package file

import (
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io"
	"os"
)

// WritePEM writes a slice of certificates to an io.writer in the PEM format.
func WritePEM(certs []*x509.Certificate, output io.Writer) error {

	for _, cert := range certs {
		block := &pem.Block{
			Type:  "CERTIFICATE",
			Bytes: []byte(cert.Raw),
		}

		if err := pem.Encode(output, block); err != nil {
			return err
		}
	}
	return nil
}

// GetWriter return an io.Writer for use by WritePEM.
func GetWriter(outputFilePath string) (io.Writer, error) {
	if outputFilePath != "" {
		file, err := os.Create(outputFilePath)
		if err != nil {
			return nil, err
		}

		fmt.Printf("Saving to %s\n", outputFilePath)
		return file, nil
	}

	// File path not set, output to stdout.
	return os.Stdout, nil
}
