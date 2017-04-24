package controller

import (
	"github.com/nextmetaphor/gwAPI/connection"
	"net/http"
)

func ReloadGatewayGroup(credentials connection.ConnectionCredentials, connector connection.Connector) (interface{}, http.Response, error) {
	var responseDocument interface{}

	httpResponse, httpRequestError := connection.GenericRESTCall(credentials, connector, http.MethodGet, GATEWAY_RELOAD_GROUP_URI, nil, responseDocument)

	return responseDocument, httpResponse, httpRequestError
}