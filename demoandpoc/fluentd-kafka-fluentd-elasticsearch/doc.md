# Install kafka cluster 
``` bash
helm repo add bitnami https://charts.bitnami.com/bitnami
helm upgrade  --install --namespace=kafka  es-kafka  bitnami/kafka -f value.yaml
```

# Test this kafka
### create topic
``` bash
kafka-topics.sh --bootstrap-server 127.0.0.1:9092 --topic test --create --partitions 1 --replication-factor 1
kafka-topics.sh --bootstrap-server 127.0.0.1:9092 --list
```

### Test
```bash
/opt/bitnami/kafka/bin/kafka-console-producer.sh --broker-list localhost:9092 --producer.config /opt/bitnami/kafka/config/producer.properties --topic test
/opt/bitnami/kafka/bin/kafka-console-consumer.sh --bootstrap-server localhost:9092 --topic test --consumer.config /opt/bitnami/kafka/config/consumer.properties --from-beginning
```