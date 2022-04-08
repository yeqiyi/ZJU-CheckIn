package client

import (
	"net/http"
	"net/http/cookiejar"

	"golang.org/x/net/publicsuffix"
)


func NewClient() *http.Client{
	cjar,_:=cookiejar.New(&cookiejar.Options{PublicSuffixList:publicsuffix.List})
	client:=&http.Client{Jar: cjar}
	return client
}