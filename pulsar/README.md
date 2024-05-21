# Apache Pulsar learning notes

## Pulsar in a container

```bash

docker run -it \
-p 6650:6650 \
-p 8080:8080 \
--mount source=pulsardata,target=/pulsar/data \
--mount source=pulsarconf,target=/pulsar/conf \
apachepulsar/pulsar:3.2.3 \
bin/pulsar standalone

--platform=linux/amd64

## If issues with metadata store:

-e PULSAR_STANDALONE_USE_ZOOKEEPER=1 \

pulsar://localhost:6650
http://localhost:8080

Using pulsar client

pip3 install pulsar-client

Run the consumer `python3 consumer.py`
Send some messages to be consumed by the producer `python3 producer.py`


## Get topic stats

`curl http://localhost:8080/admin/v2/persistent/public/default/my-topic/stats | python -m json.tool`

```

## Pulsar in docker compose

## Pulsar raw

## Pulsar in K8s
