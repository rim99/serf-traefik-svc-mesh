
cp traefik-conf/static.yaml traefik.conf.d/static.yaml

./engine -svc webapp-1 \
         -endpoints webapp-1 \
         -template traefik-conf/local-svc.yaml.tpl \
         -output-dir ./traefik.conf.d \
         -output-file local-webapp-1.yaml

./engine -svc webapp-2 \
         -endpoints "192.168.122.116" \
         -template traefik-conf/remote-svc.yaml.tpl \
         -output-dir ./traefik.conf.d \
         -output-file remote-webapp-2.yaml

