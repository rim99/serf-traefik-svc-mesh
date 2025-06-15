# This the seed of the cluster, no need to specifiy "join" parameter

export WORKSPACE=$(pwd) # For "event-handler.sh", the parent folder contains traefik templates and configs

serf agent --node=vm1 \
  --event-handler=event-handler.sh \
  --tag zone=a \
  --tag services=webapp-1

