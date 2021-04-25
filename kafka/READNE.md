# Kafka Quick Start

## Create New Topic
```
kafka-topics.sh \
--create \
--topic test-16-3 \
--partitions 16 \
--replication-factor 3 \
--config retention.ms=86400000 \
--config min.insync.replicas=2 \
--bootstrap-server kafka1:9092,kafka2:9092,kafka3:9092
```

## Test the Producer Performance (WPS)
```
kafka-producer-perf-test.sh \
--topic test-16-3 \
--throughput -1 \
--num-records 1000000 \
--record-size 1024 \
--producer-props acks=all bootstrap.servers=kafka1:9092,kafka2:9092,kafka3:9092
```

## Test the Consumer Performance (RPS)
```
kafka-consumer-perf-test.sh \
--topic test-16-3 \
--broker-list kafka1:9092,kafka2:9092,kafka3:9092 \
--messages 1000000 | \
jq -R .|jq -sr 'map(./",")|transpose|map(join(": "))[]'
```

## Test the End-To-End Latency
> USAGE: java kafka.tools.EndToEndLatency$ broker_list topic num_messages producer_acks message_size_bytes [optional] properties_file

```
kafka-run-class.sh kafka.tools.EndToEndLatency \
kafka1:9092,kafka2:9092,kafka3:9092 \
test-16-3 10000 1 1024
```

## Refrance
- [https://kafka.apache.org/quickstart](https://kafka.apache.org/quickstart)
- [https://medium.com/metrosystemsro/apache-kafka-how-to-test-performance-for-clients-configured-with-ssl-encryption-3356d3a0d52b](https://medium.com/metrosystemsro/apache-kafka-how-to-test-performance-for-clients-configured-with-ssl-encryption-3356d3a0d52b)

- [https://www.confluent.io/blog/configure-kafka-to-minimize-latency/](https://www.confluent.io/blog/configure-kafka-to-minimize-latency/)
