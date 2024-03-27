package kafka

import (
	"context"

	"github.com/Shopify/sarama"
	"github.com/ThreeDotsLabs/watermill"
	wmkafka "github.com/ThreeDotsLabs/watermill-kafka/v2/pkg/kafka"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/rs/zerolog"
)

type Consumer struct {
	subscriber *wmkafka.Subscriber
	topic      string
	logger     *zerolog.Logger
}

// NewConsumer sets up watermill subscriber and returns our consumer
func NewConsumer(config *Config, logger *zerolog.Logger) (*Consumer, error) {
	saramaSubscriberConfig := wmkafka.DefaultSaramaSubscriberConfig()
	// equivalent of auto.offset.reset: earliest
	saramaSubscriberConfig.Consumer.Offsets.Initial = sarama.OffsetOldest
	// note that autocommit is enabled
	saramaSubscriberConfig.Version = sarama.V2_6_0_0

	subscriber, err := wmkafka.NewSubscriber(
		wmkafka.SubscriberConfig{
			Brokers:               config.BrokerAddresses,
			Unmarshaler:           wmkafka.DefaultMarshaler{},
			OverwriteSaramaConfig: saramaSubscriberConfig,
			ConsumerGroup:         config.GroupID,
		},
		watermill.NewStdLogger(false, false), // if we don't like this logger we'll need an adapter for zerologger
	)
	if err != nil {
		return nil, err
	}
	return &Consumer{
		subscriber: subscriber,
		topic:      config.Topic,
		logger:     logger,
	}, nil
}

// Start reads messages from subscriber and processes them with passed in function.
//
//	eg: for msg := range messages {
//			fmt.Printf("received message: %s, payload: %s", msg.UUID, string(msg.Payload))
//			msg.Ack() }
func (c *Consumer) Start(ctx context.Context, process func(messages <-chan *message.Message)) {
	messages, err := c.subscriber.Subscribe(ctx, c.topic)
	if err != nil {
		c.logger.Fatal().Err(err).Msgf("could not subscribe to topic: %s", c.topic)
	}
	go process(messages)
}
