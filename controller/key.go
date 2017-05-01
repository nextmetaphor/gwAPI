package controller

import (
	"fmt"
	"github.com/TykTechnologies/tykcommon"
	"github.com/nextmetaphor/gwAPI/connection"
	"net/http"
)

func SelectKeys(credentials connection.ConnectionCredentials, connector connection.Connector, apiID string) (*[]struct{ tykcommon.APIDefinition }, http.Response, error) {
	responseDocument := new([]struct{ tykcommon.APIDefinition })
	httpResponse, httpRequestError := connection.GenericRESTCall(credentials, connector, http.MethodGet, fmt.Sprintf(keySelectURI, apiID), nil, responseDocument)

	return responseDocument, httpResponse, httpRequestError
}
