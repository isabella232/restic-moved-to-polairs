package backend

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/pem"
	"io/ioutil"
	"net"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/restic/restic/internal/debug"
	"github.com/restic/restic/internal/errors"
)

// TransportOptions collects various options which can be set for an HTTP based
// transport.
type TransportOptions struct {
	// contains filenames of PEM encoded root certificates to trust
	RootCertFilenames []string

	// contains the name of a file containing the TLS client certificate and private key in PEM format
	TLSClientCertKeyFilename string
}

// readPEMCertKey reads a file and returns the PEM encoded certificate and key
// blocks.
func readPEMCertKey(filename string) (certs []byte, key []byte, err error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, nil, errors.Wrap(err, "ReadFile")
	}

	var block *pem.Block
	for {
		if len(data) == 0 {
			break
		}
		block, data = pem.Decode(data)
		if block == nil {
			break
		}

		switch {
		case strings.HasSuffix(block.Type, "CERTIFICATE"):
			certs = append(certs, pem.EncodeToMemory(block)...)
		case strings.HasSuffix(block.Type, "PRIVATE KEY"):
			if key != nil {
				return nil, nil, errors.Errorf("error loading TLS cert and key from %v: more than one private key found", filename)
			}
			key = pem.EncodeToMemory(block)
		default:
			return nil, nil, errors.Errorf("error loading TLS cert and key from %v: unknown block type %v found", filename, block.Type)
		}
	}

	return certs, key, nil
}

// Transport returns a new http.RoundTripper with default settings applied. If
// a custom rootCertFilename is non-empty, it must point to a valid PEM file,
// otherwise the function will return an error.
func Transport(opts TransportOptions) (http.RoundTripper, error) {
	var tlsConfig tls.Config
	envVar := os.Getenv("FORCE_CERT_VALIDATION")
	envVar = strings.ToLower(strings.TrimSpace(envVar))
	if "true" == envVar {
		tlsConfig.InsecureSkipVerify = false
	} else {
		tlsConfig.InsecureSkipVerify = true
	}

	// copied from net/http
	tr := &http.Transport{
		Proxy: http.ProxyFromEnvironment,
		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
			DualStack: true,
		}).DialContext,
		ForceAttemptHTTP2:     true,
		MaxIdleConns:          100,
		MaxIdleConnsPerHost:   100,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
		TLSClientConfig:       &tlsConfig,
	}

	if opts.TLSClientCertKeyFilename != "" {
		certs, key, err := readPEMCertKey(opts.TLSClientCertKeyFilename)
		if err != nil {
			return nil, err
		}

		crt, err := tls.X509KeyPair(certs, key)
		if err != nil {
			return nil, errors.Errorf("parse TLS client cert or key: %v", err)
		}
		tr.TLSClientConfig.Certificates = []tls.Certificate{crt}
	}

	if opts.RootCertFilenames != nil {
		pool := x509.NewCertPool()
		for _, filename := range opts.RootCertFilenames {
			if filename == "" {
				return nil, errors.Errorf("empty filename for root certificate supplied")
			}
			b, err := ioutil.ReadFile(filename)
			if err != nil {
				return nil, errors.Errorf("unable to read root certificate: %v", err)
			}
			if ok := pool.AppendCertsFromPEM(b); !ok {
				return nil, errors.Errorf("cannot parse root certificate from %q", filename)
			}
		}
		tr.TLSClientConfig.RootCAs = pool
	}

	// wrap in the debug round tripper (if active)
	return debug.RoundTripper(tr), nil
}
