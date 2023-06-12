/opt/bitnami/kafka/bin/kafka-topics.sh --create --zookeeper es-kafka-zookeeper:2181 --replication-factor 1 --partitions 1 --topic test

/opt/bitnami/kafka/bin/kafka-topics.sh --create --bootstrap-server localhost:9097 --replication-factor 1 --partitions 1 --topic test1

kafka-topics.sh --bootstrap-server 127.0.0.1:9092 --topic test --create --partitions 1 --replication-factor 1
kafka-topics.sh --bootstrap-server 127.0.0.1:9092 --list


/opt/bitnami/kafka/bin/kafka-console-producer.sh --broker-list localhost:9092 --producer.config /opt/bitnami/kafka/config/producer.properties --topic test
/opt/bitnami/kafka/bin/kafka-console-consumer.sh --bootstrap-server localhost:9092 --topic test --consumer.config /opt/bitnami/kafka/config/consumer.properties --from-beginning



/opt/bitnami/kafka/bin/kafka-console-consumer.sh --bootstrap-server internal-aa252ff75d6814b92a7f600693dedb01-2078365334.ap-southeast-1.elb.amazonaws.com:9094 --topic test --consumer.config /opt/bitnami/kafka/config/consumer.properties --from-beginning


/opt/bitnami/kafka/bin/kafka-console-producer.sh --bootstrap-server es-kafka.kafka.svc.cluster.local:9092 --producer.config /opt/bitnami/kafka/config/producer.properties --topic test
/opt/bitnami/kafka/bin/kafka-console-consumer.sh --bootstrap-server es-kafka.kafka.svc.cluster.local:9092 --topic test --consumer.config /opt/bitnami/kafka/config/consumer.properties --from-beginning


/opt/bitnami/kafka/bin/kafka-console-producer.sh --broker-list es-kafka-0.es-kafka-headless.kafka.svc.cluster.local:9092 --producer.config /opt/bitnami/kafka/config/producer.properties --topic test
/opt/bitnami/kafka/bin/kafka-console-consumer.sh --bootstrap-server es-kafka-0.es-kafka-headless.kafka.svc.cluster.local:9092 --topic test --consumer.config /opt/bitnami/kafka/config/consumer.properties --from-beginning


/opt/bitnami/kafka/bin/kafka-console-consumer.sh --bootstrap-server es-kafka-0.es-kafka-headless.kafka.svc.cluster.local:9092 --topic dev-eks-console-01 --consumer.config /opt/bitnami/kafka/config/consumer.properties --from-beginning

