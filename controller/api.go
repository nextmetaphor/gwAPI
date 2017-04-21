package controller

import (
	"bytes"
	"net/http"
	"fmt"
	"github.com/nextmetaphor/gwAPI/connection"
)

const (
	API_DEFINITION_READ_URI = "/tyk/apis/%s"
	API_DEFINITION_UPDATE_URI = "/tyk/apis/%s"
)

func ReadAPI(credentials connection.ConnectionCredentials, connector connection.Connector, apiID string, responseDocument interface{}) (http.Response, error) {
	return connection.GenericRESTCall(credentials, connector, http.MethodGet, fmt.Sprintf(API_DEFINITION_READ_URI, apiID), nil, responseDocument)
}

func UpdateAPI(credentials connection.ConnectionCredentials, connector connection.Connector, apiID string, apiDefinition *string, responseDocument interface{}) (http.Response, error) {
	apiDefinitionBuffer := bytes.NewBufferString(*apiDefinition)

	return connection.GenericRESTCall(credentials, connector, http.MethodPut, fmt.Sprintf(API_DEFINITION_UPDATE_URI, apiID), apiDefinitionBuffer, responseDocument)
}

func ReloadGatewayGroup(credentials connection.ConnectionCredentials, connector connection.Connector, responseDocument interface{}) (http.Response, error) {
	return connection.GenericRESTCall(credentials, connector, http.MethodGet, GATEWAY_RELOAD_GROUP_URI, nil, responseDocument)
}