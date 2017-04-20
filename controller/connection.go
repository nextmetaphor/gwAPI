package controller

import (
	"crypto/tls"
	log "github.com/Sirupsen/logrus"
	"github.com/square/go-jose/json"
	"io"
	"io/ioutil"
	"net/http"
)

type Connector interface {
	NewRequest(credentials ConnectionCredentials, method string, urlStr string, body io.Reader) (*http.Request, error)
	DoHttpRequest(req *http.Request, responseObject interface{}) http.Response
}

type ConnectionCredentials struct {
	GatewayURL string
	AuthToken  string
}

type Connection struct {
}

func (con Connection) NewRequest(credentials ConnectionCredentials, method string, urlStr string, body io.Reader) (*http.Request, error) {
	req, reqErr := http.NewRequest(method, credentials.GatewayURL + urlStr, body)
	req.Header.Add("x-tyk-authorization", credentials.AuthToken)
	return req, reqErr
}

func (con Connection) DoHttpRequest(req *http.Request, responseObject interface{}) http.Response {
	log.WithFields(log.Fields{
		"URL":    req.URL,
		"Method": req.Method}).Debug("Calling DoHttpRequest")

	tlsConfig := &tls.Config{
		Renegotiation:      tls.RenegotiateFreelyAsClient,
		InsecureSkipVerify: true}
	transport := &http.Transport{TLSClientConfig: tlsConfig}
	client := &http.Client{Transport: transport}

	// Make the HTTP request
	resp, respErr := client.Do(req)

	log.WithFields(log.Fields{
		"StatusCode": resp.StatusCode}).Debug("Got initial response")

	if (respErr == nil) && (resp.StatusCode >= http.StatusOK) && (resp.StatusCode <= http.StatusIMUsed) {

		// Got a response - try and read it all
		if importRes, importResErr := ioutil.ReadAll(resp.Body); importResErr != nil {
			log.WithFields(log.Fields{"importResErr": importResErr}).
				Debug("Could not read all of the message")
		} else {
			log.WithFields(log.Fields{
				"Body": string(importRes)}).Debug("Got complete response")

			// Marshall the response into the responseObject parameter
			if unmarshallErr := json.Unmarshal(importRes, responseObject); unmarshallErr != nil {
				log.WithFields(log.Fields{"unmarshallErr": unmarshallErr}).
					Debug("Cannot unmarshall object")
			}
		}

	} else {
		if respErr != nil {
			log.WithFields(log.Fields{"resErr": respErr}).
				Debug("HTTP response error")
		}
	}

	return *resp
}
