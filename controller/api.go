package controller

import (
	"bytes"
	"net/http"
	"fmt"
	"github.com/nextmetaphor/gwAPI/connection"
	"github.com/TykTechnologies/tykcommon"
)

func SelectAPIs(credentials connection.ConnectionCredentials, connector connection.Connector) (*[]struct { tykcommon.APIDefinition }, http.Response, error) {
	responseDocument := new([]struct {tykcommon.APIDefinition})
	httpResponse, httpRequestError := connection.GenericRESTCall(credentials, connector, http.MethodGet, API_DEFINITION_SELECT_URI, nil, responseDocument)

	return responseDocument, httpResponse, httpRequestError
}

func ReadAPI(credentials connection.ConnectionCredentials, connector connection.Connector, apiID string) (*tykcommon.APIDefinition, http.Response, error) {
	responseDocument := new(tykcommon.APIDefinition)
	httpResponse, httpRequestError := connection.GenericRESTCall(credentials, connector, http.MethodGet, fmt.Sprintf(API_DEFINITION_READ_URI, apiID), nil, responseDocument)

	return responseDocument, httpResponse, httpRequestError
}

func UpdateAPI(credentials connection.ConnectionCredentials, connector connection.Connector, apiID string, apiDefinition *string) (*tykcommon.APIDefinition, http.Response, error) {
	responseDocument := new(tykcommon.APIDefinition)
	apiDefinitionBuffer := bytes.NewBufferString(*apiDefinition)

	httpResponse, httpRequestError := connection.GenericRESTCall(credentials, connector, http.MethodPut, fmt.Sprintf(API_DEFINITION_UPDATE_URI, apiID), apiDefinitionBuffer, responseDocument)

	return responseDocument, httpResponse, httpRequestError
}