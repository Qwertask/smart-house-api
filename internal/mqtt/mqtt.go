package mqtt

import (
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"log"
)

type MQTTService struct {
	client   mqtt.Client
	handlers map[string]func(topic string, payload []byte)
}

func (s *MQTTService) Publish(topic string, qos byte, retained bool, payload interface{}) error {
	token := s.client.Publish(topic, qos, retained, payload)
	token.Wait()
	return token.Error()
}
func (s *MQTTService) Client() mqtt.Client {
	return s.client
}

func New(brokerURL string, pass string, login string) *MQTTService {
	service := &MQTTService{
		handlers: make(map[string]func(string, []byte)),
	}

	opts := mqtt.NewClientOptions().
		AddBroker(brokerURL).
		SetClientID("iot_go_server").
		SetCleanSession(true).
		SetUsername(login).
		SetPassword(pass).
		SetOnConnectHandler(func(c mqtt.Client) {
			log.Println("✅ MQTT подключено")

			for topic := range service.handlers {
				if token := c.Subscribe(topic, 1, service.handle); token.Wait() && token.Error() != nil {
					log.Printf("❌ Ошибка подписки на %s: %v\n", topic, token.Error())
				} else {
					log.Printf("📩 Подписка на %s\n", topic)
				}
			}
		})

	service.client = mqtt.NewClient(opts)
	return service
}

func (s *MQTTService) Connect() error {
	if token := s.client.Connect(); token.Wait() && token.Error() != nil {
		return token.Error()
	}
	return nil
}

func (s *MQTTService) RegisterHandler(topic string, handler func(string, []byte)) {
	s.handlers[topic] = handler
}

func (s *MQTTService) handle(client mqtt.Client, msg mqtt.Message) {
	if handler, ok := s.handlers[msg.Topic()]; ok {
		handler(msg.Topic(), msg.Payload())
	}
}
