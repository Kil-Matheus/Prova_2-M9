package main

import (
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"database/sql"
	"log"
	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

type Message struct {
	msg *kafka.Message

}

// Inicializar banco de dados SQLite
func initDB() {
	var err error
	db, err = sql.Open("sqlite3", "kil-sqlite.db")
	if err != nil {
		log.Fatal(err)
	}

	// Criação da tabela para armazenar as mensagens
	sqlStmt := `
	CREATE TABLE IF NOT EXISTS messages (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		msg TEXT
	);
	`
	_, err = db.Exec(sqlStmt)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		return
	}
}

// Função para salvar mensagem no banco de dados
func saveMessage(message Message) error {
	stmt, err := db.Prepare("INSERT INTO messages(msg) VALUES(?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(message.msg)
	if err != nil {
		return err
	}

	return nil
}

func consumeMessages(consumer *kafka.Consumer, topic string) {
	consumer.SubscribeTopics([]string{topic}, nil)
	for {
		msg, err := consumer.ReadMessage(-1)
		if err == nil {
			fmt.Printf("Received message: %s\n", string(msg.Value))
			// Salvar mensagem no banco de dados
			//saveMessage(msg)
		} else {
			fmt.Printf("Consumer error: %v (%v)\n", err, msg)
			break
		}
	}
}

func main() {
	// Inicializar banco de dados
	initDB()
	// Configurações do consumidor
	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:29092,localhost:39092",
		"group.id":          "go-consumer-group",
		"auto.offset.reset": "earliest",
	})
	if err != nil {
		panic(err)
	}
	defer consumer.Close()

	topic := "qualidadeAr"

	// Iniciar consumidor em uma goroutine separada
	go consumeMessages(consumer, topic)

	// Manter o programa em execução
	select {}
}
