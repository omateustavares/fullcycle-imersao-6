package main

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"github.com/omateustavares/fullcycle-imersao-6/email"
	"github.com/omateustavares/fullcycle-imersao-6/kafka"
	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	gomail "gopkg.in/mail.v2"
)

func main() {
	var emailCh = make(chan email.Email)
	var msgChan = make(chan *ckafka.Message)

	// port, _ := strconv.Atoi(os.Getenv("587"))
	
	d := gomail.NewDialer(
		"smtp.mailgun.org",
		587,
		"teste@sandbox0270ddb825c64e9bbf11ad65b642d419.mailgun.org",
		"82fa5f618f3eba099810b8cfcd5e93e7-787e6567-1ac85fb5",
	)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	es := email.NewMailSender()
	es.From = "mateustm17@gmail.com"
	es.Dialer = d

	go es.Send(emailCh)

	configMapConsumer := &ckafka.ConfigMap{
		"bootstrap.servers": "kafka:9094",
		// "bootstrap.servers": os.Getenv("BOOTSTRAP_SERVERS"),
		// "security.protocol": os.Getenv("SECURITY_PROTOCOL"),
		// "sasl.mechanisms":   os.Getenv("SASL_MECHANISMS"),
		// "sasl.username":     os.Getenv("SASL_USERNAME"),
		// "sasl.password":     os.Getenv("SASL_PASSWORD"),
		"client.id": "emailapp",
		"group.id":  "emailapp",
		// "session.timeout.ms": 45000,
	}
	topics := []string{"emails"}
	consumer := kafka.NewConsumer(configMapConsumer, topics)
	go consumer.Consume(msgChan)

	fmt.Println("Consumindo msgs")
	for msg := range msgChan {
		var input email.Email
		json.Unmarshal(msg.Value, &input)
		fmt.Println("Recebendo mensagem")
		emailCh <- input
	}
}