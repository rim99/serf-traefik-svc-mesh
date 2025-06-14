
cp traefik-conf/static.yaml traefik.conf.d/static.yaml

./engine -svc webapp-2 \
         -endpoints webapp-2 \
         -template traefik-conf/local-svc.yaml.tpl \
         -output-dir ./traefik.conf.d \
         -output-file local-webapp-2.yaml

./engine -svc webapp-1 \
         -endpoints "192.168.122.128" \
         -template traefik-conf/remote-svc.yaml.tpl \
         -output-dir ./traefik.conf.d \
         -output-file remote-webapp-1.yaml

