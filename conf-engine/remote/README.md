# Usage

```
serf members -format=json | ./remote-conf-engine -services webapp1 -template traefik-conf/remote-svc.yaml.tpl -output-dir ./traefik.conf.d
```
