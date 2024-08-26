package gateways

// this was all copied from edge-network, could be candidate for shared lib or SDK exposed by identity-api
import (
	"encoding/json"
	"io"
	"time"

	"github.com/DIMO-Network/vehicle-signal-decoding/internal/config"

	"github.com/DIMO-Network/shared"
	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
)

var ErrNotFound = errors.New("not found")
var ErrBadRequest = errors.New("bad request")

//go:generate mockgen -source identity_api.go -destination mocks/identity_api_mock.go
type IdentityAPI interface {
	GetVehicleByDeviceAddr(ethAddress common.Address) (*VehicleInfo, error)
}

type identityAPIService struct {
	httpClient     shared.HTTPClientWrapper
	logger         zerolog.Logger
	identityAPIURL string
}

func NewIdentityAPIService(logger *zerolog.Logger, settings *config.Settings) IdentityAPI {
	h := map[string]string{}
	h["Content-Type"] = "application/json"
	hcw, _ := shared.NewHTTPClientWrapper("", "", 10*time.Second, h, false) // ok to ignore err since only used for tor check

	return &identityAPIService{
		httpClient:     hcw,
		logger:         *logger,
		identityAPIURL: settings.IdentityAPIURL,
	}
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

func (i *identityAPIService) fetchVehicleWithQuery(query string) (*VehicleInfo, error) {
	// GraphQL request
	requestPayload := GraphQLRequest{Query: query}
	payloadBytes, err := json.Marshal(requestPayload)
	if err != nil {
		return nil, err
	}

	// POST request
	res, err := i.httpClient.ExecuteRequest(i.identityAPIURL, "POST", payloadBytes)
	if err != nil {
		i.logger.Err(err).Send()
		if _, ok := err.(shared.HTTPResponseError); !ok {
			return nil, errors.Wrapf(err, "error calling identity api to get vehicles definition from url %s", i.identityAPIURL)
		}
	}
	defer res.Body.Close() // nolint
	if res.StatusCode == 404 {
		return nil, ErrNotFound
	}

	if res.StatusCode == 400 {
		return nil, ErrBadRequest
	}

	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, errors.Wrapf(err, "error get vehicles definition from url %s", i.identityAPIURL)
	}

	var vehicleResponse struct {
		Data struct {
			AftermarketDevice struct {
				Vehicle VehicleInfo `json:"vehicle"`
			} `json:"aftermarketDevice"`
		} `json:"data"`
	}

	if err := json.Unmarshal(bodyBytes, &vehicleResponse); err != nil {
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
