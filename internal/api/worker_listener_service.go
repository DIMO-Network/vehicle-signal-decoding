package api

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/DIMO-Network/shared"
	"github.com/DIMO-Network/shared/db"
	"github.com/DIMO-Network/vehicle-signal-decoding/internal/core/commands"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
)

const (
	vehicleSignalDecodingEventType = "zone.dimo.task.tesla.poll.status.update"
)

type WorkerListenerService struct {
	DBS    func() *db.ReaderWriter
	logger *zerolog.Logger
}

type VechicleSignalDecodingData struct {
}

func NewWorkerListenerService(dbs func() *db.ReaderWriter, logger *zerolog.Logger) *WorkerListenerService {
	return &WorkerListenerService{DBS: dbs, logger: logger}
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
		service := commands.NewRunTestSignalCommandHandler(i.DBS)
		command := &commands.RunTestSignalCommandRequest{}

		return service.Execute(ctx, command)
	default:
		return fmt.Errorf("unexpected event type %s", event.Type)
	}
}
