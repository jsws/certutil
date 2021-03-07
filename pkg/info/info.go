package info

import (
	"crypto/x509"
	"fmt"
	"time"

	"github.com/jsws/certutil/pkg/aia"
)

// PrintInfo prints information
func PrintInfo(certs []*x509.Certificate, serverName string) {
	fmt.Println()

	for i, cert := range certs {
		fmt.Printf("Certificate %d\n", i)
		printCert(cert)
	}

	// TODO: Do something with the returned chain(s) here.
	_, err := IsValid(certs, serverName)
	if err != nil {
		fmt.Println(err)
		fmt.Println("Certificate not valid ❌")
	} else {
		fmt.Println("Certificate is valid 🔒")
	}
}

// printCert prints information about a single cert.
func printCert(cert *x509.Certificate) {

	fmt.Printf("  Subject: %s\n", cert.Subject.String())
	fmt.Printf("  Issuer:  %s\n", cert.Issuer.String())

	if cert.DNSNames != nil {
		fmt.Printf("  DNS Names: ")
		for n, dnsname := range cert.DNSNames {
			fmt.Printf("%s", dnsname)
			if n != len(cert.DNSNames)-1 {
				fmt.Printf(", ")
			}
		}

		fmt.Println()
	}
	fmt.Println("  Vailidity Period")
	printDates(cert.NotBefore, cert.NotAfter)
	fmt.Println()
}

// printDates is a helper function to print the dates of a certificate nicely.
func printDates(notBefore time.Time, notAfter time.Time) {
	now := time.Now()
	fmt.Printf("      Not Before:  %s", notBefore)
	if now.Before(notBefore) {
		fmt.Printf(" ❌\n")

	} else {
		fmt.Printf(" ✅\n")

	}
	fmt.Printf("      Not After:   %s", notAfter)
	if now.After(notAfter) {
		fmt.Printf(" ❌\n")

	} else {
		fmt.Printf(" ✅\n")
	}
}

// IsValid checks weather a certificate chain is valid. Returning the chain(s) is valid or error if not.
// Does NOT do: revocation checking or Certificate Transparency checking
func IsValid(certs []*x509.Certificate, serverName string) ([][]*x509.Certificate, error) {
	intermediates := x509.NewCertPool()
	// Add all certs after leaf to intermediates
	for _, cert := range certs[1:] {
		intermediates.AddCert(cert)
	}

	chains, err := certs[0].Verify(x509.VerifyOptions{
		DNSName:       serverName,
		Intermediates: intermediates,
	})
	if err != nil {
		// Type asseration to check error type
		if _, ok := err.(x509.UnknownAuthorityError); ok {
			fmt.Println("Signed by Unknown Authority, performing AIA fetching 🥏 🐕")
			// Perform AIA fetching
			err := aia.Validate(certs)
			if err != nil {
				return nil, err
			}
		} else {
			return nil, err
		}
	}

	// Certificate is valid
	return chains, nil
}
