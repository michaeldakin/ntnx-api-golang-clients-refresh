//Api classes for lcm's golang SDK
package api

import (
    "github.com/nutanix/ntnx-api-golang-clients/lcm-go-client/v4/client"
	import3 "github.com/nutanix/ntnx-api-golang-clients/lcm-go-client/v4/models/lcm/v4/common"
	import2 "github.com/nutanix/ntnx-api-golang-clients/lcm-go-client/v4/models/lcm/v4/operations"
	"encoding/json"
	"net/http"
    "net/url"
)

type PrecheckApi struct {
  ApiClient *client.ApiClient
}

func NewPrecheckApi(apiClient *client.ApiClient) *PrecheckApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &PrecheckApi{
		ApiClient: apiClient,
	}
	return a
}


/**
    Perform a LCM precheck operation
    Perform LCM precheck

    parameters:-
    -> body (lcm.v4.common.PrecheckSpec) (optional)
    -> args (map[string]interface{}) (optional) : Additional Arguments

    returns: (*lcm.v4.operations.PrecheckApiResponse, error)
*/
func (api *PrecheckApi) Precheck(body *import3.PrecheckSpec, args ...map[string]interface{}) (*import2.PrecheckApiResponse, error) {
    argMap := make(map[string]interface{})
	if len(args) > 0 {
        argMap = args[0]
    }

    uri := "/api/lcm/v4.0.a1/operations/$actions/performPrecheck"


	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := url.Values{}

	// to determine the Content-Type header 
    contentTypes := []string{"application/json"}

    // to determine the Accept header 
	accepts := []string{"application/json"} 

    // Header Params
    if ifMatch, ifMatchOk := argMap["If-Match"].(string); ifMatchOk {
        headerParams["If-Match"] = ifMatch
    }
    if ifNoneMatch, ifNoneMatchOk := argMap["If-None-Match"].(string); ifNoneMatchOk {
        headerParams["If-None-Match"] = ifNoneMatch
    }
    authNames := []string{"basicAuthScheme"}

    responseBody, err := api.ApiClient.CallApi(&uri, http.MethodPost, body, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
    if nil != err || nil == responseBody{
    	return nil, err
	}
    unmarshalledResp := new(import2.PrecheckApiResponse)
    json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

