#! /bin/bash

export WORKSPACE=$(pwd) # For "event-handler.sh", the parent folder contains traefik templates and configs

serf agent --node=vm2 \
  --event-handler=event-handler.sh \
  --tag zone=b \
  --tag services=webapp-2 \
  --join 192.168.122.128:7946 # ip-of-vm1
