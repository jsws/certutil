package info

import (
	"testing"

	"github.com/jsws/certutil/pkg/connect"
)

func TestIsValid(t *testing.T) {
	const (
		host = "badssl.com"
		port = 443
	)

	tables := []struct {
		sni   string
		valid bool
	}{
		{"badssl.com", true},
		{"expired.badssl.com", false},
		{"wrong.host.badssl.com", false},
		{"self-signed.badssl.com", false},
		{"untrusted-root.badssl.com", false},
		{"self-signed.badssl.com", false},
		{"no-subject.badssl.com", false},
		{"incomplete-chain.badssl.com", true}, // Valid with AIA fetching.
	}

	for _, table := range tables {
		certs, _ := connect.Connect(host, table.sni, port)
		_, err := IsValid(certs, table.sni)
		if err != nil && table.valid {
			t.Errorf("Validity of %s was incorrect, got: invalid, want: valid.", table.sni)
		} else if err == nil && !table.valid {
			t.Errorf("Validity of %s was incorrect, got: valid, want: invalid.", table.sni)
		}
	}
}
