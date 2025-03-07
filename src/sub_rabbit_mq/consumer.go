package subrabbitmq

import (
    "consumerApi/src/models"
    "consumerApi/src/utils"
    "encoding/json"
    "fmt"
    "log"
    "os"
    "time"
    "github.com/joho/godotenv"
    amqp "github.com/rabbitmq/amqp091-go"
)

func failOnError(err error, msg string) {
    if err != nil {
        log.Panicf("%s: %s", msg, err)
    }
}

func SubToQueue() {
    // Cargar variables de entorno desde el archivo .env
    err := godotenv.Load()
    if err != nil {
        log.Fatalf("Error loading .env file")
    }

    rabbitmqUser := os.Getenv("RABBITMQ_USER")
    rabbitmqPass := os.Getenv("RABBITMQ_PASS")
    rabbitmqHost := os.Getenv("RABBITMQ_HOST")
    rabbitmqPort := os.Getenv("RABBITMQ_PORT")

    connStr := fmt.Sprintf("amqp://%s:%s@%s:%s/", rabbitmqUser, rabbitmqPass, rabbitmqHost, rabbitmqPort)
    conn, err := amqp.Dial(connStr)
    failOnError(err, "Failed to connect to RabbitMQ")
    defer conn.Close()

    ch, err := conn.Channel()
    failOnError(err, "Failed to open a channel")
    defer ch.Close()

    q, err := ch.QueueDeclare(
        "messages", // name
        true,       // durable
        false,      // delete when unused
        false,      // exclusive
        false,      // no-wait
        nil,        // arguments
    )
    failOnError(err, "Failed to declare a queue")

    msgs, err := ch.Consume(
        q.Name, // queue
        "",     // consumer
        true,   // auto-ack
        false,  // exclusive
        false,  // no-local
        false,  // no-wait
        nil,    // args
    )
    failOnError(err, "Failed to register a consumer")

    go func() {
        for d := range msgs {
            log.Printf("Received a message: %s", d.Body)

            // Transform received message to the format required
            var receivedMessage models.ReceivedMessage
            err := json.Unmarshal(d.Body, &receivedMessage)
            if err != nil {
                log.Printf("Error decoding received message: %s", err)
                continue
            }

            if receivedMessage.Id == 0 {
                log.Printf("Error: receivedMessage.Id is 0. Message: %s", d.Body)
                continue
            }

            time.Sleep(500 * time.Millisecond)
            message_welc := "Welcome user! " + receivedMessage.FullName 
            sendMessage := models.SendMessage{
                UserId:     receivedMessage.Id,
                FullName:   receivedMessage.FullName,
                Email:      receivedMessage.Email,
                Description: message_welc, 
            }
            fmt.Println("Message to send: ", sendMessage)

            sendMessageBody, err := json.Marshal(sendMessage)
            if err != nil {
                log.Printf("Error encoding send message: %s", err)
                continue
            }

            // Imprimir el cuerpo del mensaje JSON serializado
            fmt.Println("Serialized JSON message: ", string(sendMessageBody))

            utils.SendMessageToAPI(string(sendMessageBody))
        }
    }()

    log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
    select {}
}
