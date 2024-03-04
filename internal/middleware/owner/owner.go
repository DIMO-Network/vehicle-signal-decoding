package owner

import (
	"context"
	"fmt"

	"github.com/DIMO-Network/devices-api/pkg/grpc"

	"github.com/DIMO-Network/vehicle-signal-decoding/internal/core/services"
	"github.com/golang-jwt/jwt/v5"

	pb "github.com/DIMO-Network/users-api/pkg/grpc"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// todo can we move this to shared, then share with valuations api and device-data api too

var errNotFound = fiber.NewError(fiber.StatusNotFound, "Device not found.")

// New creates a new middleware handler that checks whether a user is authorized to access
// a user device. For the middleware to allow the request to proceed:
//
//   - The request must have a valid JWT, identifying a user.
//   - There must be a userDeviceID or ethAddr path parameter, and that device must exist.
//   - Either the user owns the device, or the user's account has an Ethereum address that
//     owns the corresponding NFT.
func New(usersClient pb.UserServiceClient, devicesClient services.UserDeviceService, logger *zerolog.Logger) fiber.Handler {
	return func(c *fiber.Ctx) error {
		userID := GetUserID(c)
		udi := c.Params("userDeviceID")
		ethAddr := c.Params("ethAddr")
		logger := logger.With().Str("userId", userID).Str("userDeviceId", udi).Str("device_address", ethAddr).Logger()

		// i don't get what these are used for?
		c.Locals("userID", userID)
		c.Locals("userDeviceID", udi)
		c.Locals("logger", &logger)
		var device *grpc.UserDevice
		var err error

		if udi != "" {
			device, err = devicesClient.GetUserDevice(context.Background(), udi)
		} else if ethAddr != "" {
			address := common.HexToAddress(ethAddr)
			device, err = devicesClient.GetUserDeviceByEthAddr(context.Background(), address) // if identity api had a way to filterBy address, could use it instead
		} else {
			return fmt.Errorf("no userDeviceID or ethAddr params found for owner validation")
		}
		if err != nil {
			if st, ok := status.FromError(err); ok && st.Code() == codes.NotFound {
				return errNotFound
			}
			return err
		}

		if device.UserId == userID {
			return c.Next()
		}

		user, err := usersClient.GetUser(c.Context(), &pb.GetUserRequest{Id: userID})
		if err != nil {
			if st, ok := status.FromError(err); ok && st.Code() == codes.NotFound {
				return errNotFound
			}
			return err
		}

		if user.EthereumAddress == nil {
			return errNotFound
		}

		if common.HexToAddress(*user.EthereumAddress) == common.BytesToAddress(device.OwnerAddress) {
			return c.Next()
		}

		return errNotFound
	}
}

func GetUserID(c *fiber.Ctx) string {
	token := c.Locals("user").(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)
	userID := claims["sub"].(string)
	return userID
}
