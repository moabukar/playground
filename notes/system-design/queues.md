# Queues vs PubSub

## Queuing architecture

- System A > Queue > System B
- System A will send a message to the queue, and System B will poll the queue and retrieve the message
- Known as a pull-based architecture
- Only one consumer can process the message
- We have SQS & Amazon MQ

## PubSub architecture

- System A > Topic > System B/System C/System D
- System A will send a message to the topic, and System B will subscribe to the topic and retrieve the message
- Known as a push-based architecture
- Multiple consumers can process the message

## Real world example

- You will get a weather alert on your phone
- The weather alert is sent to a topic
- Your phone is subscribed to the topic
- Your phone will receive the weather alert
- Under the hood, it's a pubsub architecture
