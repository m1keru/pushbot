#!/bin/bash
curl -XPOST -H "Content-Type: application/json" -d @json.reqest 127.0.0.1:7777/alertmanager
