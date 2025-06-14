
serf agent --node=vm2 \
  --tag az=b \
  --tag service=webap-2 \
  --join 192.168.122.128:7946 # ip-of-vm1
