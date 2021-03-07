package connect

import (
	"testing"
)

func TestConnect(t *testing.T) {
	// Not sure how useful this test really is.
	host := "github.com"
	certs, err := Connect(host, host, 443)
	if err != nil {
		t.Errorf("Error connecting:\n %s", err)
	}
	if len(certs) == 0 {
		t.Errorf("No certificates received:\n %s", err)
	}
}
