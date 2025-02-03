$env:ROLE="consumer"
$env:TOPIC="test"
$env:KAFKA_BROKER="localhost:9092"
$env:PARTITIONS="4"
go run main.go
