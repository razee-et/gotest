# message-broker-sdk-go

This repository contains an SDK for producing and consuming messages from the
Message Broker.

### Spinning up an broker for local development

start broker
```bash
make amq-start
```
stop broker
```bash
make amq-stop
```


###Running unit and integration tests
Running unit tests
```bash
make unit-test
```
Running integration tests
```bash
make amq-start          # start broker
make integration-test   # run the tests
make amq-stop           # stop broker
```
