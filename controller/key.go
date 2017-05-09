package controller

import (
	"fmt"
	"github.com/nextmetaphor/gwAPI/connection"
	"github.com/nextmetaphor/gwAPI/definition"
	"net/http"
)

func SelectKeys(credentials connection.ConnectionCredentials, connector connection.Connector, apiID string) (*definition.APIAllKeys, http.Response, error) {
	responseDocument := new(definition.APIAllKeys)
	httpResponse, httpRequestError := connection.GenericRESTCall(credentials, connector, http.MethodGet, fmt.Sprintf(keySelectURI, apiID), nil, responseDocument)

	return responseDocument, httpResponse, httpRequestError
}
