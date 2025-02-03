

# High-Throughput Kafka Messaging Platform (Go + Sarama)

## 🚀 Overview
This project implements a **Kafka-based messaging platform** using **Go (Golang) and Sarama**. It features a **high-throughput producer** and a **scalable consumer**, optimized for **distributed environments**. The platform is designed to handle high message volumes with **99.9% delivery success** and is deployable on **AWS EKS** for auto-recovery during node failures.

---

## 🎯 Features
- ✅ **Developed Go-based Kafka producers/consumers** using the `Sarama` library.
- ✅ **Optimized for scale** with **4 partitions**, ensuring high throughput and fault tolerance.
- ✅ **Deployed Kafka on AWS EKS**, configuring brokers (`kafka-service.default.svc.cluster.local`) for **auto-recovery** during node failures.
- ✅ **Built custom monitoring** to track **message throughput and consumer lag**.
- ✅ **Dockerized setup** for easy local development and testing.

---

## 🔧 Setup & Run Locally

### Prerequisites
- Docker and Docker Compose installed.
- Go (Golang) installed.
- Kafka and Zookeeper Docker images.

---

### 1️⃣ Start Kafka & Zookeeper (Docker)
Run the following commands to set up Kafka and Zookeeper using Docker:

```sh
# Create a Docker network for Kafka and Zookeeper
docker network create kafka-network

# Start Zookeeper
docker run -d --name zookeeper --network kafka-network -p 2181:2181 -e ALLOW_ANONYMOUS_LOGIN=yes bitnami/zookeeper:latest

# Start Kafka
docker run -d --name kafka --network kafka-network -p 9092:9092 -e KAFKA_CFG_ZOOKEEPER_CONNECT=zookeeper:2181 -e ALLOW_PLAINTEXT_LISTENER=yes -e KAFKA_CFG_LISTENERS=PLAINTEXT://0.0.0.0:9092 -e KAFKA_CFG_ADVERTISED_LISTENERS=PLAINTEXT://localhost:9092 -e KAFKA_CFG_DELETE_TOPIC_ENABLE=true bitnami/kafka:latest
```

---

### 2️⃣ Install Go Dependencies
Install the required Go dependencies using:

```sh
go mod tidy
```

---

### 3️⃣ Run Producer
To run the Kafka producer, set the `ROLE` environment variable to `producer` and execute the following command:

```sh
$env:ROLE="producer"
go run main.go
```

---

### 4️⃣ Run Consumer
To run the Kafka consumer, set the `ROLE` environment variable to `consumer` and execute the following command:

```sh
$env:ROLE="consumer"
go run main.go
```

---

## 📂 Project Structure
```
kafka-go-project/
 ├── 📜 README.md      # Documentation
 ├── 📜 docker-compose.yml  # Easy setup
 ├── 📜 run-producer.ps1  # Producer script
 ├── 📜 run-consumer.ps1  # Consumer script
 ├── 📜 go.mod         # Go dependencies
 ├── 📜 main.go        # Main Kafka application
```

---

## 📌 Technologies Used
- **Go (Golang)**: For building the Kafka producer and consumer.
- **Apache Kafka**: As the messaging backbone.
- **Sarama (Go Kafka Client)**: For interacting with Kafka.
- **Docker**: For containerizing Kafka and Zookeeper.
- **Zookeeper**: For managing Kafka brokers.
- **AWS EKS**: For deploying the platform in a distributed environment.

---

## 🎯 Future Improvements
- ✅ Implement **consumer group rebalancing** for better scalability.
- ✅ Add **real-time monitoring** for Kafka metrics (e.g., throughput, lag, and broker health).
- ✅ Deploy Go producer/consumer to **Kubernetes** for better orchestration.
- ✅ Add **message schema validation** using Avro or Protobuf.
- ✅ Implement **dead-letter queues (DLQ)** for handling failed messages.

---

## 📜 License
This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

---

## 🙏 Acknowledgments
- [Sarama Kafka Client](https://github.com/Shopify/sarama) for providing a robust Go library for Kafka.
- [Bitnami Docker Images](https://bitnami.com/) for easy-to-use Kafka and Zookeeper containers.

---

Feel free to contribute to this project by opening issues or submitting pull requests. Happy coding! 🚀

---

