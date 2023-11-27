package commands

import (
	"context"
	"fmt"
	"github.com/DIMO-Network/vehicle-signal-decoding/internal/core/common"
	p_grpc "github.com/DIMO-Network/vehicle-signal-decoding/pkg/grpc"
	"github.com/segmentio/ksuid"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/DIMO-Network/vehicle-signal-decoding/internal/infrastructure/exceptions"
	"github.com/pkg/errors"

	"github.com/DIMO-Network/shared/db"
	"github.com/DIMO-Network/vehicle-signal-decoding/internal/infrastructure/db/models"
)

type CreateJobByEtherumAddressCommandHandler struct {
	DBS func() *db.ReaderWriter
}

func NewCreateJobByEtherumAddressCommandHandler(dbs func() *db.ReaderWriter) CreateJobByEtherumAddressCommandHandler {
	return CreateJobByEtherumAddressCommandHandler{DBS: dbs}
}

type CreateJobCommandRequest struct {
	Command        string
	EtherumAddress string
}

func (h CreateJobByEtherumAddressCommandHandler) Execute(ctx context.Context, command *CreateJobCommandRequest) (*p_grpc.GetJobByEtherumAddressItemResponse, error) {

	ethAddrBytes, err := common.ResolveEtherumAddressFromString(command.EtherumAddress)
	if err != nil {
		return nil, &exceptions.ValidationError{
			Err: fmt.Errorf("invalid ethereum address: %w", err),
		}
	}

	job := &models.Job{}
	job.ID = ksuid.New().String()
	job.Command = command.Command
	job.DeviceEthereumAddress = ethAddrBytes

	err = job.Insert(ctx, h.DBS().Writer, boil.Infer())
	if err != nil {
		return nil, &exceptions.InternalError{Err: errors.Wrapf(err, "error inserting job: %s", command.Command)}
	}

	jobItem, err := models.Jobs(models.JobWhere.DeviceEthereumAddress.EQ(ethAddrBytes)).One(ctx, h.DBS().Reader)
	if err != nil {
		return nil, &exceptions.InternalError{
			Err: fmt.Errorf("failed to get job"),
		}
	}

	result := &p_grpc.GetJobByEtherumAddressItemResponse{
		Id:            jobItem.ID,
		Command:       jobItem.Command,
		Status:        jobItem.Status,
		CreatedAt:     timestamppb.New(jobItem.CreatedAt),
		LastExecution: timestamppb.New(jobItem.LastExecution),
	}

	return result, nil
}
