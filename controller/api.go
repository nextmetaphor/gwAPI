package controller

import (
	log "github.com/Sirupsen/logrus"
	"bytes"
	"net/http"
	"io"
	"fmt"
)

const (
	API_DEFINITION_UPDATE_URI = "/tyk/apis/%s"
	GATEWAY_RELOAD_GROUP_URI = "/tyk/reload/group"
)

func (connection Connection) genericAPICall(method, uri string, body io.Reader, responseDocument interface{}) (http.Response, error) {
	var httpResponse http.Response
	var reqError error

	req, reqError := connection.NewRequest(method, uri, body)
	if (reqError != nil) {
		log.WithFields(log.Fields{
			"error": reqError}).Debug("Error creating request")
	} else {
		httpResponse = connection.DoHttpRequest(req, responseDocument)
	}

	return httpResponse, reqError
}

func (connection Connection) UpdateAPI(apiID string, apiDefinition *string, responseDocument interface{}) (http.Response, error) {
	apiDefinitionBuffer := bytes.NewBufferString(*apiDefinition)

	return connection.genericAPICall(http.MethodPut, fmt.Sprintf(API_DEFINITION_UPDATE_URI, apiID), apiDefinitionBuffer, responseDocument)
}

func (connection Connection) ReloadGatewayGroup(responseDocument interface{}) (http.Response, error) {
	return connection.genericAPICall(http.MethodGet, GATEWAY_RELOAD_GROUP_URI, nil, responseDocument)
}