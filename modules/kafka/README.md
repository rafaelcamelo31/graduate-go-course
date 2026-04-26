# Apache Kafka

## What is Apache Kafka

Apache Kafka is first introduced by Linkedin, when they had to scale their platform to microservices. Today, Kafka is open-source publish/subscribe queue (pub-sub) solution, widely used in software development industry.

<img src="image/kafka_architecture.png" height="500">

Producer publishes messages to Kafka. Kafka stores these messages as commit log files in its filesystem. Any number of consumers can then subscribe to receive these messages.

Unlike other queueing technologies, messages are persistent and are not deleted once read. This allows consumers to read all the previous messages, which helps recovering from failures.

## Topics

Events and messages are organized and durably stored in Topics. We can understand Topics as a folder in a filesystem, and the events are the files in that folder.

Each event registered inside a topic is called "offset", from first record to Nth record. Consumers are able to read these offsets simultaniously.

<img src="image/topics.png" height="500">

Each offset has a data structure as shown in the figure below.
- Headers
- Key
- Value (JSON)
- Timestamp

<img src="image/offset.png" height="500">

## Partitions

A topic is subdivided into Partitions. This enables event messages to be read and stored in distributed setting, giving much more resiliency and throughput.

Distributed setting is configured by Replicator Factor. Each broker can be configured replicator factor to make brokers contain numbers of partitions necessary to guarantee resiliency.

<div style="display: flex; justify-content: space-around; margin: 25px">
    <img src="image/partitions2.png" height="300">
    <img src="image/partitions.png" height="300">
</div>

<div style="text-align: center">
    <img src="image/distributed_partitions.png" height="500">
</div>