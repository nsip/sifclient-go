package sifclient

import (
    "github.com/go-resty/resty"
    // "fmt"
    "crypto/x509"
    "crypto/tls"
    "io/ioutil"
    "log"
)

func (s *SIFClientData) GetClient() *resty.Client {
    c := resty.New()
    // Config requjired
    c.SetDebug(s.Debug)

    rootCAs, _ := x509.SystemCertPool()
    if rootCAs == nil {
        rootCAs = x509.NewCertPool()
    }

    // Read in the cert file (XXX quick hack on my Mac)
    //  - must be a better way
    localCertFile:= "/usr/local/etc/openssl/cert.pem"
    certs, err := ioutil.ReadFile(localCertFile)
    if err != nil {
        log.Println("Failed to append %q to RootCAs: %v", localCertFile, err)
    }

    // Append our cert to the system pool
    if ok := rootCAs.AppendCertsFromPEM(certs); !ok {
        log.Println("No certs appended, using system certs only")
    }

    // Trust the augmented cert pool in our client

    c.SetTLSClientConfig(&tls.Config{RootCAs: rootCAs})

    // XXX note - skipping certificates is working
    c.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})

    return c
}

func (s *SIFClientData) GetRequest() *resty.Request {
    r := s.GetClient().R()
    if s.Debug {
    	r.EnableTrace()
    }
    r.SetHeader("Content-Type", "application/xml; charset=UTF-8")
    r.SetHeader("Timestamp", s.Timestamp)
    r.SetHeader("Authorization", s.GetAuthorization())
    r.SetHeader("applicationKey", s.ApplicationKey)
    r.SetHeader("authenticateduser", s.UserToken)
    r.SetHeader("requestType", "IMMEDIATE")
	r.SetHeader("navigationpage", "1")
	r.SetHeader("navigationpagesize", "100")
    // XXX debugging
    // fmt.Println("Headers", r.Header)
    return r
}
