package api

import (
	"testing"

	mock_commands "github.com/DIMO-Network/vehicle-signal-decoding/internal/core/commands/mocks"
	"github.com/DIMO-Network/vehicle-signal-decoding/internal/infrastructure/dbtest"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestWorkerListenerService_processMessage(t *testing.T) {
	// setup mock etc
	ctrl := gomock.NewController(t)
	signalHandler := mock_commands.NewMockRunTestSignalCommandHandler(ctrl)

	worker := NewWorkerListenerService(*dbtest.Logger(), signalHandler)
	signalHandler.EXPECT().Execute(gomock.Any(), gomock.Any())
	// deal with incorrectly formatted message
	msg := &message.Message{
		UUID:     uuid.New().String(),
		Metadata: nil,
		Payload:  []byte(payload),
	}
	err := worker.processMessage(msg)

	assert.NoError(t, err)
}

const payload = `{
    "data": {
        "signals": [
            {
                "canbus_vin_toyota580v1": {
                    "_stamp ": "2023-03-16T13:12:17.815155 ",
                    "value ": 0.0
                },
                "canbus_vin_toyota580v2": {
                    "_stamp ": "2023-03-16T13:12:17.826589 ",
                    "value ": 160.0
                },
                "canbus_vin_toyota580v4": {
                    "_stamp ": "2023-03-16T13:12:17.825384 ",
                    "value ": 2.0
                }
            }
        ]
    },
    "id": "2N66pVsRwgj0jzoKBxCuLJ7JpBe",
    "source": "autopi/status/transform",
    "specversion": "1.0 ",
    "subject": "4918de88-d252-dfe7-e028-fbd9e3cdaabf",
    "time": "2023-03-16T13:12:47.591Z",
    "type": "zone.dimo.canbus.signal.update"
}`
