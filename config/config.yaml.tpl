---
logger:
  logFile:
  debug: true
daemon:
  port: 7777
alertmanager:
    host: "0.0.0.0"
    uri: /alertmanager
    prefix: "" # optional
    name: "" # optional
rabbitmq:
    amqp: "amqp://guest:guest@localhost:5672/"
    fail_timeout: 15 # seconds
    schema: "" # optional
    queue_name: "alertmanager_queue"
    exchange_name: "alertmanager_exchange"
