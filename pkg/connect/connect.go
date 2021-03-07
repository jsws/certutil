package connect

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"strconv"
)

// Connect connects to a host and port and returns the the list of certificates provided.
func Connect(host string, sni string, port uint16) ([]*x509.Certificate, error) {
	socket := host + ":" + strconv.Itoa(int(port))

	fmt.Printf("Connecting to %s\n", socket)
	conn, err := tls.Dial("tcp", socket, &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         sni,
	})
	if err != nil {
		return nil, err
	}
	conn.Close()
	fmt.Printf("Connected to %s\n", conn.RemoteAddr())
	fmt.Printf("Recieved %d certificate(s).\n", len(conn.ConnectionState().PeerCertificates))

	return conn.ConnectionState().PeerCertificates, nil
}
