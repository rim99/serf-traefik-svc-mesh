# For local services to access remote services
http:
  routers:
    to-remote-{{ .Svc }}:
      rule: "Host(`{{ .Svc }}`)"
      entryPoints:
        - "local"
      service: "{{ .Svc }}-remote"

  services:
    {{ .Svc }}-remote:
      loadBalancer:
        servers:
        {{- range .Endpoints }}
        - url: http://{{ . }}:18080/
        {{- end }}

