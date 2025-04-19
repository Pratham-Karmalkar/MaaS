# 🚀 Kafka-ELK Data Pipeline

This project sets up a real-time log processing pipeline using **Kafka** and the **ELK Stack** — fully containerized with **Docker** for easy deployment.

---

## 🔧 Components

- **Filebeat** – Ships log data to Kafka topics.  
- **Kafka** – Buffers and queues log data.  
- **Logstash** – Pulls data from Kafka, processes it, and forwards it.  
- **Elasticsearch** – Indexes and stores structured logs.  
- **Kibana** – Visualizes logs through rich dashboards.

---

## 🔄 Data Flow

```
Logs → Filebeat → Kafka → Logstash → Elasticsearch → Kibana
```

---

## 🐳 Dockerized Setup

Each component runs in its own Docker container, orchestrated via YAML files — ensuring a portable, scalable, and ready-to-use pipeline.

---

Use this pipeline to monitor logs, detect anomalies, and gain real-time insights into your system.

---
