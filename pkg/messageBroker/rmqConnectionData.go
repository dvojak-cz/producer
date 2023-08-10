package messagebroker

type ConnectionDataRMQ struct {
	Host     string
	Port     string
	User     string
	Password string
}

func NewConnectionDataRMQ(host, port, user, password string) *ConnectionDataRMQ {
	return &ConnectionDataRMQ{host, port, user, password}
}

func (c *ConnectionDataRMQ) ConnectionString() string {
	return "amqp://" + c.User + ":" + c.Password + "@" + c.Host + ":" + c.Port
}