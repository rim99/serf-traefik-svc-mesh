#! /bin/bash
cp traefik-conf/static.yaml traefik.conf.d/static.yaml

# Maybe this needs to be executed after first round of local healthcheck
./engine -svc webapp-1 \
         -endpoints webapp-1 \
         -template traefik-conf/local-svc.yaml.tpl \
         -output-dir ./traefik.conf.d \
         -output-file local-webapp-1.yaml


