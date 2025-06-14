# This the seed of the cluster, no need to specifiy "join" parameter
serf agent --node=vm1 \
  --tag az=a \
  --tag service=webapp-1
