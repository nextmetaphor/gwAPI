package controller

import (
	"net/http"
	"io"
	"testing"
)

type TestConnector struct {}

func (tc TestConnector) NewRequest(credentials ConnectionCredentials, method string, urlStr string, body io.Reader) (*http.Request, error) {
	return &http.Request{}, nil
}

func (tc TestConnector) DoHttpRequest(req *http.Request, responseObject interface{}) http.Response {
	return *(&http.Response{})
}

func TestGenericAPICall(*testing.T) {
	creds := &ConnectionCredentials{
		GatewayURL: "gatewayURL",
		AuthToken: "authToken"}

	connector := &TestConnector{}

	genericAPICall(*creds, *connector, http.MethodGet, "", nil, nil)

}