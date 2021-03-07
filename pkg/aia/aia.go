package aia

import (
	"crypto/x509"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

// Validate performs  Authority Information Access (AIA) fetching for cert validation.
func Validate(certs []*x509.Certificate) error {
	//https://github.com/golang/go/issues/31773
	intermediates := x509.NewCertPool()

	firstCert := certs[0]
	// TODO: Check all IssuingCertificateURL not just the first.
	if len(firstCert.IssuingCertificateURL) >= 1 && firstCert.IssuingCertificateURL[0] != "" {
		fmt.Printf("Getting %s\n", firstCert.IssuingCertificateURL[0])
		intermediate, err := getCertificate(firstCert.IssuingCertificateURL[0])
		if err != nil {
			return err
		}
		intermediates.AddCert(intermediate)

		_, err = certs[0].Verify(x509.VerifyOptions{
			Intermediates: intermediates,
		})
		if err != nil {
			fmt.Println("Still not valid")
			certs = append(certs, intermediate)
			return Validate(certs)
		}
		fmt.Println(intermediate.Subject)
		fmt.Println(intermediate.Issuer)
		fmt.Println("Valid with AIA fetching, full certificate chain not given ⛓️")

		return nil

	}
	return errors.New("No issuing URL provided in certificate, can't perform AIA fetching")
}

// getCertificate return a certificate for a given URL or an error if the URL scheme is not
// implemented.
func getCertificate(URL string) (*x509.Certificate, error) {
	u, err := url.Parse(URL)
	if err != nil {
		return nil, err
	}
	if u.Scheme == "http" {
		return getCertificateHTTP(URL)
	}
	return nil, errors.New("Can't get certificate with " + u.Scheme + " URL")
}

// getCertificateHTTP return an x509 certifate from a given HTTP URL.
func getCertificateHTTP(URL string) (*x509.Certificate, error) {
	resp, err := http.Get(URL)
	if resp != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		fmt.Println("Error fetching certificate")
		return nil, err
	}
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error getting AIA certificate from CA")
		return nil, err
	}
	intermediate, err := x509.ParseCertificate(data)
	if err != nil {
		fmt.Println("Error parsing certificate")
		return nil, err
	}

	return intermediate, nil
}
