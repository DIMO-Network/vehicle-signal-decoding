package api

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/DIMO-Network/shared"
	"github.com/DIMO-Network/vehicle-signal-decoding/internal/core/commands"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
)

const (
	vehicleSignalDecodingEventType = "zone.dimo.canbus.signal.update"
)

type WorkerListenerService struct {
	logger  zerolog.Logger
	handler commands.RunTestSignalCommandHandler
}

type VechicleSignalDecodingData struct {
	Signals []map[string]commands.RunTestSignalItemCommandRequest `json:"signals"`
}

func NewWorkerListenerService(logger zerolog.Logger, handler commands.RunTestSignalCommandHandler) *WorkerListenerService {
	return &WorkerListenerService{
		logger:  logger,
		handler: handler,
	}
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
		i.logger.Warn().Str("payload", string(msg.Payload)).Msg("failed to unmarshall processMessage payload")
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
		command := &commands.RunTestSignalCommandRequest{
			AutoPIUnitID: event.Subject,
			Time:         event.Time,
			Signals:      event.Data.Signals[0], // due to way ingest is processing the message we get an array of 1 with a map of the signals
		}

		return i.handler.Execute(ctx, command)
	default:
		return fmt.Errorf("unexpected event type %s", event.Type)
	}
}
