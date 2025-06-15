#! /bin/bash
WORKSPACE=${WORKSPACE:-/tmp}

# it connects to webapp-2
serf members -format=json | remote-conf-engine -services webapp-2 \
                                                -template $WORKSPACE/traefik-conf/remote-svc.yaml.tpl \
                                                -output-dir $WORKSPACE/traefik.conf.d
