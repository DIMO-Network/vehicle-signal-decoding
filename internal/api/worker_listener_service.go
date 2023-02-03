package api

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/DIMO-Network/shared"
	"github.com/DIMO-Network/shared/db"
	"github.com/DIMO-Network/vehicle-signal-decoding/internal/core/commands"
	"github.com/DIMO-Network/vehicle-signal-decoding/internal/core/services"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
)

const (
	vehicleSignalDecodingEventType = "zone.dimo.canbus.signal.update"
)

type WorkerListenerService struct {
	DBS               func() *db.ReaderWriter
	logger            zerolog.Logger
	userDeviceService services.UserDeviceService
}

type VechicleSignalDecodingData struct {
	Signals map[string]commands.RunTestSignalItemCommandRequest `json:"signals"`
}

func NewWorkerListenerService(dbs func() *db.ReaderWriter, logger zerolog.Logger, userDeviceService services.UserDeviceService) *WorkerListenerService {
	return &WorkerListenerService{DBS: dbs, logger: logger, userDeviceService: userDeviceService}
}

func (i *WorkerListenerService) ProcessWorker(messages <-chan *message.Message) {
	for msg := range messages {
		err := i.processMessage(msg)
		if err != nil {
			i.logger.Err(err).Msg("error processing task status message")
		}
	}
}

func (i *WorkerListenerService) processMessage(msg *message.Message) error {
	// Keep the pipeline moving no matter what.
	defer func() { msg.Ack() }()

	event := new(shared.CloudEvent[VechicleSignalDecodingData])
	if err := json.Unmarshal(msg.Payload, event); err != nil {
		return errors.Wrap(err, "error parsing vehicle signal decoding payload")
	}

	return i.processEvent(event)
}

func (i *WorkerListenerService) processEvent(event *shared.CloudEvent[VechicleSignalDecodingData]) error {
	var (
		ctx = context.Background()
	)

	switch event.Type {
	case vehicleSignalDecodingEventType:
		service := commands.NewRunTestSignalCommandHandler(i.DBS, i.logger, i.userDeviceService)

		command := &commands.RunTestSignalCommandRequest{
			AutoPIUnitID: event.Subject,
			Time:         event.Time,
			Signals:      event.Data.Signals,
		}

		return service.Execute(ctx, command)
	default:
		return fmt.Errorf("unexpected event type %s", event.Type)
	}
}
