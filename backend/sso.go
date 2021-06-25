package main

import (
	"context"
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"log"
	"net/http"
	"net/url"

	"github.com/crewjam/saml/samlsp"
)

func loadSSO() (*samlsp.Middleware, error) {
	keyPair, err := tls.LoadX509KeyPair("dh-application.nerdwiese.de.crt", "dh-application.nerdwiese.de.key")
	if err != nil {
		return nil, err
	}
	keyPair.Leaf, err = x509.ParseCertificate(keyPair.Certificate[0])
	if err != nil {
		return nil, err
	}

	// alternative test idp: https://samltest.id/saml/idp
	idpMetadataURL, err := url.Parse("https://idp.dhbw-loerrach.de/idp/shibboleth")
	if err != nil {
		return nil, err
	}
	idpMetadata, err := samlsp.FetchMetadata(context.Background(), http.DefaultClient,
		*idpMetadataURL)
	if err != nil {
		return nil, err
	}
	log.Println(idpMetadata.EntityID)

	rootURL, err := url.Parse("https://dh-application.nerdwiese.de")
	if err != nil {
		return nil, err
	}

	samlSP, _ := samlsp.New(samlsp.Options{
		URL:         *rootURL,
		Key:         keyPair.PrivateKey.(*rsa.PrivateKey),
		Certificate: keyPair.Leaf,
		IDPMetadata: idpMetadata,
		SignRequest: true,
	})

	return samlSP, nil
}
