package test

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	pb "github.com/DIMO-Network/users-api/pkg/grpc"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
)

func BuildRequest(method, url, body string) *http.Request {
	req, _ := http.NewRequest(
		method,
		url,
		strings.NewReader(body),
	)
	req.Header.Set("Content-Type", "application/json")

	return req
}

func Logger() *zerolog.Logger {
	l := zerolog.New(os.Stdout).With().
		Timestamp().
		Str("app", "devices-api").
		Logger()
	return &l
}

// AuthInjectorTestHandler injects fake jwt with sub
func AuthInjectorTestHandler(userID string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"sub": userID,
			"nbf": time.Now().Unix(),
		})

		c.Locals("user", token)
		return c.Next()
	}
}

// SetupAppFiber sets up app fiber with defaults for testing, like our production error handler.
func SetupAppFiber(logger zerolog.Logger) *fiber.App {
	app := fiber.New(fiber.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return ErrorHandler(c, err, &logger, false)
		},
	})
	return app
}

// ErrorHandler custom handler to log recovered errors using our logger and return json instead of string
func ErrorHandler(c *fiber.Ctx, err error, logger *zerolog.Logger, isProduction bool) error {
	logger = getLogger(c, logger)

	code := fiber.StatusInternalServerError // Default 500 statuscode
	var e *fiber.Error
	isFiberErr := errors.As(err, &e)
	if isFiberErr {
		// Override status code if fiber.Error type
		code = e.Code
	}
	c.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)
	codeStr := strconv.Itoa(code)

	logger.Err(err).Str("httpStatusCode", codeStr).
		Str("httpMethod", c.Method()).
		Str("httpPath", c.Path()).
		Msg("caught an error from http request")
	// return an opaque error if we're in a higher level environment and we haven't specified an fiber type err.
	if !isFiberErr && isProduction {
		err = fiber.NewError(fiber.StatusInternalServerError, "Internal error")
	}

	return c.Status(code).JSON(fiber.Map{
		"code":    code,
		"message": err.Error(),
	})
}

func getLogger(c *fiber.Ctx, d *zerolog.Logger) *zerolog.Logger {
	m := c.Locals("logger")
	if m == nil {
		return d
	}

	l, ok := m.(*zerolog.Logger)
	if !ok {
		return d
	}

	return l
}

//
//type UsersClient struct {
//	Store map[string]*pb.User
//}
//
//func (c *UsersClient) GetUser(_ context.Context, in *pb.GetUserRequest, _ ...grpc.CallOption) (*pb.User, error) {
//	u, ok := c.Store[in.Id]
//	if !ok {
//		return nil, status.Error(codes.NotFound, "No user with that id found.")
//	}
//	return u, nil
//}

type UsersClient struct {
	Store map[string]*pb.User
}

func (c *UsersClient) GetUserByEthAddr(_ context.Context, _ *pb.GetUserByEthRequest, _ ...grpc.CallOption) (*pb.User, error) {
	return nil, status.Error(codes.NotFound, "Not implemented.")
}

func (c *UsersClient) GetUser(_ context.Context, in *pb.GetUserRequest, _ ...grpc.CallOption) (*pb.User, error) {
	u, ok := c.Store[in.Id]
	if !ok {
		return nil, status.Error(codes.NotFound, "No user with that id found.")
	}
	return u, nil
}
