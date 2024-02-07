# Kafka-ELK Data Pipeline

This project talks about configuring logstash for input from Kafka topics and modify the data.   

#### Components :
Filebeat – collects logs and forwards them to a Kafka topic.  
Kafka – brokers the data flow and queues it.  
Logstash – aggregates the data from the Kafka topic, processes it and ships to Elasticsearch.  
Elasticsearch – indexes the data.  
Kibana – for visualizing and analyzing the data.   

All applications run within Docker containers configured using YAML files.

