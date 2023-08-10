package env

import (
	log "github.com/sirupsen/logrus"

	"os"

	env "github.com/joho/godotenv"
)

var (
	MessageBrokerHost     string
	MessageBrokerPort     string
	MessageBrokerUser     string
	MessageBrokerPassword string
	MessageBrokerQueue    string
)

func init() {
	var envPath string
	if envPath = os.Getenv("ENVIRONMENT_FILE"); envPath == "" {
		envPath = ".env"
	}
	env.Load(envPath)
	log.Infof("env file '%s' loaded\n", envPath)

	MessageBrokerHost = os.Getenv("MESSAGE_BROKER_HOST")
	log.Debugln("MessageBrokerHost: ", MessageBrokerHost)

	MessageBrokerPort = os.Getenv("MESSAGE_BROKER_PORT")
	log.Debugln("MessageBrokerPort: ", MessageBrokerPort)

	MessageBrokerUser = os.Getenv("MESSAGE_BROKER_USER")
	log.Debugln("MessageBrokerUser: ", MessageBrokerUser)

	MessageBrokerPassword = os.Getenv("MESSAGE_BROKER_PASSWORD")
	log.Debugln("MessageBrokerPassword: ", MessageBrokerPassword)

	MessageBrokerQueue = os.Getenv("MESSAGE_BROKER_QUEUE")
	log.Debugln("MessageBrokerQueue: ", MessageBrokerQueue)
}
