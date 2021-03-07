package aia

import (
	"crypto/x509"
	"encoding/pem"
	"log"
	"testing"

	"github.com/jsws/certutil/pkg/file"
)

func TestGetCertificate(t *testing.T) {
	_, err := getCertificate("ldap://myldap")
	if err == nil {
		t.Errorf("Error not produced for LDAP URL %s", err)
	}

}

func TestGetCertificateHTTP(t *testing.T) {
	certificateURL := "http://r3.i.lencr.org/"

	cert, err := getCertificateHTTP(certificateURL)
	if err != nil {
		t.Errorf("Error getting certificate from %s", certificateURL)
	}
	expectedSubject := "CN=R3,O=Let's Encrypt,C=US"
	if cert.Subject.String() != expectedSubject {
		t.Errorf("Certificate fetched doesn't seem to be correct, %s %s", cert.Subject.String(), expectedSubject)
	}
}

func TestValidate(t *testing.T) {
	// incomplete-chain.badssl.com certificate
	const certPEM = `
-----BEGIN CERTIFICATE-----
MIIGqDCCBZCgAwIBAgIQCvBs2jemC2QTQvCh6x1Z/TANBgkqhkiG9w0BAQsFADBN
MQswCQYDVQQGEwJVUzEVMBMGA1UEChMMRGlnaUNlcnQgSW5jMScwJQYDVQQDEx5E
aWdpQ2VydCBTSEEyIFNlY3VyZSBTZXJ2ZXIgQ0EwHhcNMjAwMzIzMDAwMDAwWhcN
MjIwNTE3MTIwMDAwWjBuMQswCQYDVQQGEwJVUzETMBEGA1UECBMKQ2FsaWZvcm5p
YTEVMBMGA1UEBxMMV2FsbnV0IENyZWVrMRwwGgYDVQQKExNMdWNhcyBHYXJyb24g
VG9ycmVzMRUwEwYDVQQDDAwqLmJhZHNzbC5jb20wggEiMA0GCSqGSIb3DQEBAQUA
A4IBDwAwggEKAoIBAQDCBOz4jO4EwrPYUNVwWMyTGOtcqGhJsCK1+ZWesSssdj5s
wEtgTEzqsrTAD4C2sPlyyYYC+VxBXRMrf3HES7zplC5QN6ZnHGGM9kFCxUbTFocn
n3TrCp0RUiYhc2yETHlV5NFr6AY9SBVSrbMo26r/bv9glUp3aznxJNExtt1NwMT8
U7ltQq21fP6u9RXSM0jnInHHwhR6bCjqN0rf6my1crR+WqIW3GmxV0TbChKr3sMP
R3RcQSLhmvkbk+atIgYpLrG6SRwMJ56j+4v3QHIArJII2YxXhFOBBcvm/mtUmEAn
hccQu3Nw72kYQQdFVXz5ZD89LMOpfOuTGkyG0cqFAgMBAAGjggNhMIIDXTAfBgNV
HSMEGDAWgBQPgGEcgjFh1S8o541GOLQs4cbZ4jAdBgNVHQ4EFgQUne7Be4ELOkdp
cRh9ETeTvKUbP/swIwYDVR0RBBwwGoIMKi5iYWRzc2wuY29tggpiYWRzc2wuY29t
MA4GA1UdDwEB/wQEAwIFoDAdBgNVHSUEFjAUBggrBgEFBQcDAQYIKwYBBQUHAwIw
awYDVR0fBGQwYjAvoC2gK4YpaHR0cDovL2NybDMuZGlnaWNlcnQuY29tL3NzY2Et
c2hhMi1nNi5jcmwwL6AtoCuGKWh0dHA6Ly9jcmw0LmRpZ2ljZXJ0LmNvbS9zc2Nh
LXNoYTItZzYuY3JsMEwGA1UdIARFMEMwNwYJYIZIAYb9bAEBMCowKAYIKwYBBQUH
AgEWHGh0dHBzOi8vd3d3LmRpZ2ljZXJ0LmNvbS9DUFMwCAYGZ4EMAQIDMHwGCCsG
AQUFBwEBBHAwbjAkBggrBgEFBQcwAYYYaHR0cDovL29jc3AuZGlnaWNlcnQuY29t
MEYGCCsGAQUFBzAChjpodHRwOi8vY2FjZXJ0cy5kaWdpY2VydC5jb20vRGlnaUNl
cnRTSEEyU2VjdXJlU2VydmVyQ0EuY3J0MAwGA1UdEwEB/wQCMAAwggF+BgorBgEE
AdZ5AgQCBIIBbgSCAWoBaAB2ALvZ37wfinG1k5Qjl6qSe0c4V5UKq1LoGpCWZDaO
HtGFAAABcQhGXioAAAQDAEcwRQIgDfWVBXEuUZC2YP4Si3AQDidHC4U9e5XTGyG7
SFNDlRkCIQCzikrA1nf7boAdhvaGu2Vkct3VaI+0y8p3gmonU5d9DwB2ACJFRQdZ
VSRWlj+hL/H3bYbgIyZjrcBLf13Gg1xu4g8CAAABcQhGXlsAAAQDAEcwRQIhAMWi
Vsi2vYdxRCRsu/DMmCyhY0iJPKHE2c6ejPycIbgqAiAs3kSSS0NiUFiHBw7QaQ/s
GO+/lNYvjExlzVUWJbgNLwB2AFGjsPX9AXmcVm24N3iPDKR6zBsny/eeiEKaDf7U
iwXlAAABcQhGXnoAAAQDAEcwRQIgKsntiBqt8Au8DAABFkxISELhP3U/wb5lb76p
vfenWL0CIQDr2kLhCWP/QUNxXqGmvr1GaG9EuokTOLEnGPhGv1cMkDANBgkqhkiG
9w0BAQsFAAOCAQEA0RGxlwy3Tl0lhrUAn2mIi8LcZ9nBUyfAcCXCtYyCdEbjIP64
xgX6pzTt0WJoxzlT+MiK6fc0hECZXqpkTNVTARYtGkJoljlTK2vAdHZ0SOpm9OT4
RLfjGnImY0hiFbZ/LtsvS2Zg7cVJecqnrZe/za/nbDdljnnrll7C8O5naQuKr4te
uice3e8a4TtviFwS/wdDnJ3RrE83b1IljILbU5SV0X1NajyYkUWS7AnOmrFUUByz
MwdGrM6kt0lfJy/gvGVsgIKZocHdedPeECqAtq7FAJYanOsjNN9RbBOGhbwq0/FP
CC01zojqS10nGowxzOiqyB4m6wytmzf0QwjpMw==
-----END CERTIFICATE-----`
	certificate, _ := file.ParsePEM([]byte(certPEM))

	err := Validate([]*x509.Certificate{certificate[0]})
	if err != nil {
		t.Errorf("Erroring validating certificate with AIA %s", err)
	}
}

// PEMtoCertificate converts a PEM string to a certificate.
func PEMtoCertificate(certPEM string) *x509.Certificate {
	block, _ := pem.Decode([]byte(certPEM))
	if block == nil {
		panic("failed to parse certificate PEM")
	}
	cert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		log.Fatal(err)
	}
	return cert
}
