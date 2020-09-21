---
daemon:
  logFile:
  port: 7777
  debug: true
alertmanager:
    host: "0.0.0.0"
    uri: /alertmanager
    prefix: "" # optional
    name: "" # optional
databases:
  - type: rabbitmq
    host: "localhost"
    port: 5672
    fail_timeout: 15 # seconds
    schema: "" # optional
