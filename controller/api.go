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

func genericAPICall(credentials ConnectionCredentials, connector Connector, method, uri string, body io.Reader, responseDocument interface{}) (http.Response, error) {
	var httpResponse http.Response
	var reqError error

	req, reqError := connector.NewRequest(credentials, method, uri, body)
	if (reqError != nil) {
		log.WithFields(log.Fields{
			"error": reqError}).Debug("Error creating request")
	} else {
		httpResponse = connector.DoHttpRequest(req, responseDocument)
	}

	return httpResponse, reqError
}

func UpdateAPI(credentials ConnectionCredentials, connector Connector, apiID string, apiDefinition *string, responseDocument interface{}) (http.Response, error) {
	apiDefinitionBuffer := bytes.NewBufferString(*apiDefinition)

	return genericAPICall(credentials, connector, http.MethodPut, fmt.Sprintf(API_DEFINITION_UPDATE_URI, apiID), apiDefinitionBuffer, responseDocument)
}

func ReloadGatewayGroup(credentials ConnectionCredentials, connector Connector, responseDocument interface{}) (http.Response, error) {
	return genericAPICall(credentials, connector, http.MethodGet, GATEWAY_RELOAD_GROUP_URI, nil, responseDocument)
}