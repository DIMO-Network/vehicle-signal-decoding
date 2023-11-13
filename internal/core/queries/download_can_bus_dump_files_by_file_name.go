package queries

import (
	"context"
	"fmt"
	"github.com/DIMO-Network/vehicle-signal-decoding/internal/config"
	"github.com/DIMO-Network/vehicle-signal-decoding/internal/infrastructure/exceptions"
	p_grpc "github.com/DIMO-Network/vehicle-signal-decoding/pkg/grpc"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
	"io"
)

type DownloadCanBusDumpFileByFileNameQueryHandler struct {
	logger   *zerolog.Logger
	s3Client *s3.Client
	settings *config.Settings
}

func NewDownloadCanBusDumpFileByFileNameQueryHandler(logger *zerolog.Logger, s3Client *s3.Client, settings *config.Settings) DownloadCanBusDumpFileByFileNameQueryHandler {
	return DownloadCanBusDumpFileByFileNameQueryHandler{
		logger:   logger,
		s3Client: s3Client,
		settings: settings,
	}
}

type DownloadCanBusDumpFileByFileNameQueryRequest struct {
	FileName string
}

func (h DownloadCanBusDumpFileByFileNameQueryHandler) Handle(ctx context.Context, query *DownloadCanBusDumpFileByFileNameQueryRequest) (*p_grpc.DownloadCanBusDumpFileContentResponse, error) {
	response, err := h.s3Client.GetObject(ctx, &s3.GetObjectInput{
		Bucket: aws.String(h.settings.AWSCandumpsBucketName),
		Key:    aws.String(query.FileName),
	})
	if err != nil {
		var nsk types.NoSuchKey
		if errors.As(err, &nsk) {
			return nil, &exceptions.NotFoundError{
				Err: fmt.Errorf("the document does not exist"),
			}
		}
		return nil, err
	}
	defer response.Body.Close()

	bs, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return &p_grpc.DownloadCanBusDumpFileContentResponse{Content: bs}, nil
}
