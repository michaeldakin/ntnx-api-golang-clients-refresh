//Api classes for lifecycle's golang SDK
package api

import (
	"encoding/json"
	"github.com/michaeldakin/ntnx-api-golang-clients/lifecycle-go-client/v4/client"
	"github.com/michaeldakin/ntnx-api-golang-clients/lifecycle-go-client/v4/models/lifecycle/v4/resources"
	"net/http"
	"net/url"
	"strings"
)

type StatusApi struct {
	ApiClient     *client.ApiClient
	headersToSkip map[string]bool
}

func NewStatusApi(apiClient *client.ApiClient) *StatusApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &StatusApi{
		ApiClient: apiClient,
	}

	headers := []string{"authorization", "cookie", "host", "user-agent"}
	a.headersToSkip = make(map[string]bool)
	for _, header := range headers {
		a.headersToSkip[header] = true
	}

	return a
}

// Get the LCM framework status.
func (api *StatusApi) GetStatus(xClusterId *string, args ...map[string]interface{}) (*resources.GetStatusApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/lifecycle/v4.0.b1/resources/status"

	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := url.Values{}

	// to determine the Content-Type header
	contentTypes := []string{}

	// to determine the Accept header
	accepts := []string{"application/json"}

	if xClusterId != nil {
		headerParams["X-Cluster-Id"] = client.ParameterToString(*xClusterId, "")
	}
	// Headers provided explicitly on operation takes precedence
	for headerKey, value := range argMap {
		// Skip platform generated headers
		if !api.headersToSkip[strings.ToLower(headerKey)] {
			if value != nil {
				if headerValue, headerValueOk := value.(string); headerValueOk {
					headerParams[headerKey] = headerValue
				}
			}
		}
	}

	authNames := []string{"basicAuthScheme"}

	responseBody, err := api.ApiClient.CallApi(&uri, http.MethodGet, nil, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
	if nil != err || nil == responseBody {
		return nil, err
	}

	unmarshalledResp := new(resources.GetStatusApiResponse)
	json.Unmarshal(responseBody.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}
