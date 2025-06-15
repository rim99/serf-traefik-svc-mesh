#! /bin/bash
WORKSPACE=${WORKSPACE:-/tmp}

# It connects to webapp-1
serf members -format=json | remote-conf-engine -services webapp-1 \
                                                -template $WORKSPACE/traefik-conf/remote-svc.yaml.tpl \
                                                -output-dir $WORKSPACE/traefik.conf.d
