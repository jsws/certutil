package file

import (
	"bytes"
	"strings"
	"testing"
)

func TestWritePEM(t *testing.T) {
	// badssl.com certificate chaub

	certs, _ := ParsePEM([]byte(badsslcomChainPEM))

	var b bytes.Buffer
	err := WritePEM(certs, &b)
	if err != nil {
		t.Errorf("Error writing PEM file %s", err)
	}

	if strings.TrimSpace(b.String()) != strings.TrimSpace(badsslcomChainPEM) {
		t.Errorf("PEM not written correctly")
	}

}
