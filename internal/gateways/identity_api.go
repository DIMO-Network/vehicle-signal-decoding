package gateways

// this was all copied from edge-network, could be candidate for shared lib or SDK exposed by identity-api
import (
	"encoding/json"
	"io"
	"time"

	"github.com/DIMO-Network/vehicle-signal-decoding/internal/config"

	"github.com/DIMO-Network/shared/pkg/http"
	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
)

var ErrNotFound = errors.New("not found")
var ErrBadRequest = errors.New("bad request")

//go:generate mockgen -source identity_api.go -destination mocks/identity_api_mock.go
type IdentityAPI interface {
	GetVehicleByDeviceAddr(ethAddress common.Address) (*VehicleInfo, error)
	GetDefinitionByID(definitionID string) (*DeviceDefinition, error)
}

type identityAPIService struct {
	httpClient     http.ClientWrapper
	logger         zerolog.Logger
	identityAPIURL string
}

// NewIdentityAPIService creates a new instance of IdentityAPI, initializing it with the provided logger, settings, and HTTP client.
// httpClient is used for testing really
func NewIdentityAPIService(logger *zerolog.Logger, settings *config.Settings, httpClient http.ClientWrapper) IdentityAPI {
	if httpClient == nil {
		h := map[string]string{}
		h["Content-Type"] = "application/json"
		httpClient, _ = http.NewClientWrapper("", "", 10*time.Second, h, false) // ok to ignore err since only used for tor check
	}

	return &identityAPIService{
		httpClient:     httpClient,
		logger:         *logger,
		identityAPIURL: settings.IdentityAPIURL,
	}
}

func (i *identityAPIService) GetDefinitionByID(definitionID string) (*DeviceDefinition, error) {
	query := `{
  deviceDefinition(by: {id: "` + definitionID + `"}) {
    	model,
    	year
		manufacturer {
    	  id
    	  tokenId
    	  name
    	  tableId
		}
    	attributes {
      	name
    	  value
    	}
  	  }
	}`
	var wrapper struct {
		Data struct {
			DeviceDefinition DeviceDefinition `json:"deviceDefinition"`
		} `json:"data"`
	}
	err := i.fetchWithQuery(query, &wrapper)
	if err != nil {
		return nil, err
	}
	if wrapper.Data.DeviceDefinition.Model == "" {
		return nil, errors.Wrapf(ErrNotFound, "identity-api did not find definition for id: %s", definitionID)
	}
	return &wrapper.Data.DeviceDefinition, nil
}

func (i *identityAPIService) GetVehicleByDeviceAddr(deviceEthAddr common.Address) (*VehicleInfo, error) {
	// GraphQL query
	graphqlQuery := `{
        aftermarketDevice(by: {address: "` + deviceEthAddr.Hex() + `"}) {
			vehicle {
			  tokenId,
			  definition {
				make
				model
				year
			  }
			}
  		}
	}`

	return i.fetchVehicleWithQuery(graphqlQuery)
}

func (i *identityAPIService) fetchWithQuery(query string, result interface{}) error {
	// GraphQL request
	requestPayload := GraphQLRequest{Query: query}
	payloadBytes, err := json.Marshal(requestPayload)
	if err != nil {
		return err
	}

	// POST request
	res, err := i.httpClient.ExecuteRequest(i.identityAPIURL, "POST", payloadBytes)
	if err != nil {
		i.logger.Err(err).Str("func", "fetchWithQuery").Msgf("request payload: %s", string(payloadBytes))
		if _, ok := err.(http.ResponseError); !ok {
			return errors.Wrapf(err, "error calling identity api to get definition from url %s", i.identityAPIURL)
		}
	}
	defer res.Body.Close() // nolint

	if res.StatusCode == 404 {
		return ErrNotFound
	}
	if res.StatusCode == 400 {
		return ErrBadRequest
	}

	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		return errors.Wrapf(err, "error reading response body from url %s", i.identityAPIURL)
	}

	if err := json.Unmarshal(bodyBytes, result); err != nil {
		return err
	}

	return nil
}

func (i *identityAPIService) fetchVehicleWithQuery(query string) (*VehicleInfo, error) {
	var vehicleResponse struct {
		Data struct {
			AftermarketDevice struct {
				Vehicle VehicleInfo `json:"vehicle"`
			} `json:"aftermarketDevice"`
		} `json:"data"`
	}

	err := i.fetchWithQuery(query, &vehicleResponse)
	if err != nil {
		return nil, err
	}
	if vehicleResponse.Data.AftermarketDevice.Vehicle.TokenID == 0 {
		return nil, ErrNotFound
	}

	return &vehicleResponse.Data.AftermarketDevice.Vehicle, nil
}

// potential issue having these not in a separate models package for testing

type VehicleInfo struct {
	TokenID    uint64            `json:"tokenId"`
	Definition VehicleDefinition `json:"definition"`
}

type VehicleDefinition struct {
	Make  string `json:"make"`
	Model string `json:"model"`
	Year  int    `json:"year"`
}

type GraphQLRequest struct {
	Query string `json:"query"`
}

type DeviceDefinition struct {
	Model        string `json:"model"`
	Year         int    `json:"year"`
	Manufacturer struct {
		ID      string `json:"id"`
		TokenID int    `json:"tokenId"`
		Name    string `json:"name"`
		TableID int    `json:"tableId"`
	} `json:"manufacturer"`
	Attributes []struct {
		Name  string `json:"name"`
		Value string `json:"value"`
	} `json:"attributes"`
}
