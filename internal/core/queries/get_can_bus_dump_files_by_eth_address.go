package queries

import (
	"context"
	"fmt"
	"path/filepath"
	"regexp"

	"github.com/DIMO-Network/vehicle-signal-decoding/internal/config"
	"github.com/DIMO-Network/vehicle-signal-decoding/internal/infrastructure/exceptions"
	pgrpc "github.com/DIMO-Network/vehicle-signal-decoding/pkg/grpc"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/rs/zerolog"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type GetCanBusDumpFileByEthAddressQueryHandler struct {
	logger   *zerolog.Logger
	s3Client *s3.Client
	settings *config.Settings
}

func NewGetCanBusDumpFileByEthAddressQueryHandler(logger *zerolog.Logger, s3Client *s3.Client, settings *config.Settings) GetCanBusDumpFileByEthAddressQueryHandler {
	return GetCanBusDumpFileByEthAddressQueryHandler{
		logger:   logger,
		s3Client: s3Client,
		settings: settings,
	}
}

type GetCanBusDumpFileByEthAddressQueryRequest struct {
	EthAddress string
}

func (h GetCanBusDumpFileByEthAddressQueryHandler) Handle(ctx context.Context, query *GetCanBusDumpFileByEthAddressQueryRequest) (*pgrpc.GetCanBusDumpFileResponse, error) {

	response, err := h.s3Client.ListObjectsV2(ctx, &s3.ListObjectsV2Input{
		Bucket: aws.String(h.settings.AWSCandumpsBucketName),
		Prefix: aws.String(query.EthAddress + "/"),
	})

	if err != nil {
		return nil, &exceptions.NotFoundError{
			Err: fmt.Errorf("the bucket does not exist"),
		}
	}

	files := []*pgrpc.GetCanBusDumpFileItemResponse{}

	for _, item := range response.Contents {
		if *item.Size > 0 {
			files = append(files, &pgrpc.GetCanBusDumpFileItemResponse{
				Id:        RemoveSpecialCharacter(*item.ETag),
				Name:      filepath.Base(*item.Key),
				FullName:  *item.Key,
				Type:      filepath.Ext(*item.Key),
				CreatedAt: timestamppb.New(*item.LastModified),
			})
		}
	}

	return &pgrpc.GetCanBusDumpFileResponse{Items: files}, nil
}

func RemoveSpecialCharacter(input string) string {
	expresionRegular := regexp.MustCompile(`[^\w]`)

	cadenaSinEspeciales := expresionRegular.ReplaceAllString(input, "")

	return cadenaSinEspeciales
}
