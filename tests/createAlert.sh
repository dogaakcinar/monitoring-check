#!/bin/bash

# Define the Alertmanager API URL
ALERTMANAGER_API="http://localhost:9093/api/v1/alerts"

# Define the alert payload
ALERT_PAYLOAD='[
  {
    "labels": {
       "alertname": "infoinhibitor",
       "service": "my_service",
       "severity":"critical",
       "instance": "my_instance"
     },
     "annotations": {
        "summary": "High request latency"
     }
  }
]'

# Send the alert
curl -XPOST -H "Content-Type: application/json" -d "${ALERT_PAYLOAD}" ${ALERTMANAGER_API}
