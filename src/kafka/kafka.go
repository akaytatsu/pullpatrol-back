package kafka

import (
	"app/infrastructure/repository"
	kafka_hanlders "app/kafka/hanlders"
	"app/prisma/db"
	usecase_user "app/usecase/user"

	"github.com/segmentio/kafka-go"
)

func StartKafka() {
	repositoryUser := repository.NewRepositoryUser(db.NewClient())
	usecaseUser := usecase_user.NewService(repositoryUser)

	var topicParams []KafkaReadTopicsParams

	topicParams = append(topicParams, KafkaReadTopicsParams{
		Topic: "user",
		Handler: func(m kafka.Message) error {
			return kafka_hanlders.CreateUser(m, usecaseUser)
		},
	})

	startKafkaConnection(topicParams)
	readTopics()
}
