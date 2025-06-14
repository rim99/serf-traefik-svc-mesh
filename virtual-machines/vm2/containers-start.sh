
podman network create svc-mesh --subnet=192.168.101.0/24

# Run proxy with static IP: 192.168.101.2
# Should render templates first
podman run -d \
  --name traefik \
  --network svc-mesh \
  --ip 192.168.101.2 \
  -p 18080:18080 \
  -p 28080:28080 \
  -v ./traefik.conf.d:/etc/traefik/conf.d:ro,z \
  docker.io/library/traefik:v3.4 --configfile=/etc/traefik/conf.d/static.yaml

# Invoke webapp-1 via the proxy
podman run -d \
  --name webapp-2 \
  --network svc-mesh \
  -p 8080:8080 \
  -e APP_NAME=webapp-2 \
  --add-host webapp-1:192.168.101.2 \
  webapp:latest

