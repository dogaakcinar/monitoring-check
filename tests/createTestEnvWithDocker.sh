#!/bin/bash

# Pull the prom/alertmanager docker image
docker pull prom/alertmanager

# Create a directory for Alertmanager data
mkdir /tmp/alertmanager-data

# Create a basic configuration file for Alertmanager
echo "
route:
  group_by: ['alertname', 'cluster', 'service']
  group_wait: 10s
  group_interval: 10s
  repeat_interval: 10s
  receiver: 'webhook_receiver'
receivers:
- name: 'webhook_receiver'
  webhook_configs:
  - url: 'http://host.docker.internal:8080/api/monitor'
" >/tmp/alertmanager-data/alertmanager.yml

# Run the Alertmanager container
docker run -d -p 9093:9093 -v /tmp/alertmanager-data:/alertmanager prom/alertmanager --config.file=/alertmanager/alertmanager.yml

docker run --name mattermost-preview -d --publish 8065:8065 mattermost/mattermost-preview
