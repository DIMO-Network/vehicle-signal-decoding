package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/DIMO-Network/shared"
	"github.com/DIMO-Network/vehicle-signal-decoding/internal/config"
	"github.com/Shopify/sarama"
	"github.com/google/uuid"
	"github.com/rs/zerolog"
	"github.com/segmentio/ksuid"
	"github.com/tidwall/sjson"
)

func main() {
	logger := zerolog.New(os.Stdout).With().
		Timestamp().
		Str("app", "devices-api-test-producer").
		Logger()

	settings, err := shared.LoadConfig[config.Settings]("settings.yaml")
	if err != nil {
		logger.Fatal().Err(err).Msg("could not load settings")
	}

	clusterConfig := sarama.NewConfig()
	clusterConfig.Version = sarama.V2_6_0_0
	clusterConfig.Producer.Return.Successes = true
	clusterConfig.Producer.Return.Errors = true

	syncProducer, err := sarama.NewSyncProducer(strings.Split(settings.KafkaBrokers, ","), clusterConfig)
	if err != nil {
		logger.Fatal().Err(err).Msg("could not start test sync producer")
	}

	msg, err := buildTestMessage(uuid.NewString(), []signal{
		{
			SignalName: "canbus_vin_toyota580v1",
			Value:      "5TFSX5EN4LX072756",
		},
	})
	if err != nil {
		logger.Fatal().Err(err).Msg("failed to build test message")
	}

	message := &sarama.ProducerMessage{
		Topic: settings.DBCDecodingTopic,
		Value: sarama.StringEncoder(msg),
		Key:   sarama.StringEncoder(ksuid.New().String()),
	}

	partition, offset, err := syncProducer.SendMessage(message)
	if err != nil {
		logger.Err(err).Msg("could not produce message to topic")
	}
	logger.Info().Msgf("succesfully published message on topic. partition: %d offset: %d", partition, offset)
}

func buildTestMessage(autopiUnitID string, signals []signal) ([]byte, error) {

	json := `{
	  "data": {
		"signals": {
		  "canbus_vin_toyota580v0": {
			"_stamp": "2023-01-30T15:12:17.464970",
			"value": "0"
		  }
		}
	  },
	  "id": "2L3EreOFbikS9QT7c2Iu0NV3t4M",
	  "source": "autopi/status/transform",
	  "specversion": "1.0",
	  "subject": "f20d8f09-9f1b-7fda-fec8-a7b6d3edfb0b",
	  "time": "2023-01-30T15:12:45.911Z",
	  "type": "zone.dimo.canbus.signal.update"
	}`

	set, err := sjson.Set(json, "id", ksuid.New().String())
	if err != nil {
		return nil, err
	}
	set, err = sjson.Set(set, "subject", autopiUnitID)
	if err != nil {
		return nil, err
	}
	// process signals - note it is not an array but a json numeric object
	for _, s := range signals {
		ts := "2023-01-30T15:12:17.464970"
		set, err = sjson.Set(set, fmt.Sprintf("data.signals.:%s._stamp", s.SignalName), ts)
		if err != nil {
			return nil, err
		}
		set, err = sjson.Set(set, fmt.Sprintf("data.signals.:%s.value", s.SignalName), s.Value)
		if err != nil {
			return nil, err
		}
	}

	return []byte(set), nil
}

type signal struct {
	SignalName string
	Value      string
}
