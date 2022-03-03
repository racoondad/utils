package utils

import (
	"bytes"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
)

func GenHttpsReqWithDataWitoutTSL(urlReq, method string, v interface{}, headers map[string][]string) (*http.Response, error) {
	data, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}
	url, err := url.Parse(urlReq)
	if err != nil {
		return nil, err
	}
	req := &http.Request{
		Method: method,
		URL:    url,
		Header: headers,
		Body:   ioutil.NopCloser(bytes.NewBuffer(data)),
	}
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}
	return client.Do(req)
}

func GenHttpsReqWinoutTSL(urlReq, method string, headers map[string][]string) (*http.Response, error) {
	url, err := url.Parse(urlReq)
	if err != nil {
		return nil, err
	}
	req := &http.Request{
		Method: method,
		URL:    url,
		Header: headers,
	}
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}
	return client.Do(req)
}

func GenHttpsReqWithData(urlReq, method string, v interface{}, headers map[string][]string, caCrt []byte) (*http.Response, error) {
	pool := x509.NewCertPool()
	pool.AppendCertsFromPEM(caCrt)
	data, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}
	url, err := url.Parse(urlReq)
	if err != nil {
		return nil, err
	}
	req := &http.Request{
		Method: method,
		URL:    url,
		Header: headers,
		Body:   ioutil.NopCloser(bytes.NewBuffer(data)),
	}
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				RootCAs: pool,
			},
		},
	}
	return client.Do(req)
}

func GenHttpsReq(urlReq, method string, headers map[string][]string, caCrt []byte) (*http.Response, error) {
	pool := x509.NewCertPool()
	pool.AppendCertsFromPEM(caCrt)
	url, err := url.Parse(urlReq)
	if err != nil {
		return nil, err
	}
	req := &http.Request{
		Method: method,
		URL:    url,
		Header: headers,
	}
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				RootCAs: pool,
			},
		},
	}
	return client.Do(req)
}
