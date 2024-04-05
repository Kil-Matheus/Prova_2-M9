# Prova_2-M9

# Desenvolvimento

## Tecnologia Usadas

- Kafka
- Docker
- Go (Linguagem)

## Requisitos e Desenvolvimento

O Kafka utilizado foi o local, dockerizado pelo professor.

1. Implementar um Producer (Produtor): Deve coletar dados simulados de sensores de qualidade do ar e publicá-los em um tópico do Kafka chamado qualidadeAr. Os dados devem incluir:

        Id do sensor, timestamp, tipo de poluente e nivel da medida.

Baseado nesse requisito, a seguinte função foi criada para demonstrar que no payload vai mostrar os dados conforme requisitado. Mediante tempo, os valores impostos neles são fixos, mas é possível alterar eles para uma aplicação mais robusta

``` go
func produceMessages(producer *kafka.Producer, topic string) {
    time_stamp := time.Now().Format(time.RFC850)
    message := "'idSensor': '001',\n 'valor': '30',\n 'timestamp': '" + fmt.Sprintf("%d" ,time_stamp) + "',\n 'tipoPoluentes': 'PM2.5',\n 'nivel': '35.2'"
    fmt.Println("Sensor 1 - Produzindo mensagens...")

    for {
        producer.Produce(&kafka.Message{
            TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
            Value:          []byte(message),
        }, nil)
        time.Sleep(1 * time.Second)
    }
}
```

2. Implementar um Consumer (Consumidor): Deve assinar o tópico qualidadeAr e processar os dados recebidos, exibindo-os em um formato legível, além de armazená-los para análise posterior (escolha a forma de armazenamento que achar mais adequada).

Seguindo com o requisito seguinte, foi consumidor.go, cuja sua principal função é consumir um tópico direto do Kafka (qualidadeAR) após sua conexão com o mesmo. No vídeo é possível perceber as permanências dos dados até o consumo


``` go
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
```

*Nota: A função de salvamento dos dados não funiconou, e mediante ao tempo, a mesma ficará comentada para o funcionamento do restante do código. Não foi possível debugar. Mas o código main() cria o banco já configurado.

3. Implementar testes de Integridade e Persistência: Criar testes automatizados que validem a integridade dos dados transmitidos (verificando se os dados recebidos são iguais aos enviados) e a persistência dos dados (assegurando que os dados continuem acessíveis para consulta futura, mesmo após terem sido consumidos).

Mediante tempo, não foi possível realizar os testes.

## Funcionamento

https://drive.google.com/file/d/1h5XppGWleJTiZcQ3zacLH8FqZ3-u0ZYA/view?usp=sharing