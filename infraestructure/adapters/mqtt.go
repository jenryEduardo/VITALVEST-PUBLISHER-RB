package adapters

import (
	"encoding/json"
	"log"
	"publisher/domain"

	"github.com/streadway/amqp"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}


type ConnAMQP struct{
	conn *amqp.Connection
}


func NewConn()*ConnAMQP{
	conn, err := amqp.Dial("amqp://admin:tu_password_muy_segura@44.205.97.30:5672/")
	failOnError(err, "No se pudo conectar a RabbitMQ")
	return &ConnAMQP{conn: conn}
}

func (r *ConnAMQP)DataMqtt(data domain.DatosSensor) error {
	// Conectar a RabbitMQ


	ch, err := r.conn.Channel()
	failOnError(err, "No se pudo abrir el canal")
	defer ch.Close()

	// Declarar la cola
	q, err := ch.QueueDeclare(
		"usuarios_cola", // nombre de la cola
		false,           // durable
		false,           // delete when unused
		false,           // exclusive
		false,           // no-wait
		nil,             // argumentos
	)
	failOnError(err, "No se pudo declarar la cola")

	// Crear un objeto	
					  
	// Convertir a JSON
	jsonData, err := json.Marshal(data)
	failOnError(err, "No se pudo convertir el objeto a JSON")

	// Publicar el JSON
	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key (nombre de la cola)
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        jsonData,
		})
	failOnError(err, "No se pudo publicar el mensaje")

	log.Printf("📤 Enviado objeto JSON: %s", jsonData)

	return err
}