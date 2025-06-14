# For remote services to access local services
http:
  routers:
    to-local-{{ .Svc }}:
      rule: "Host(`{{ .Svc }}`)"
      entryPoints:
        - "remote"
      service: "{{ .Svc }}-local"

  services:
    {{ .Svc }}-local:
      loadBalancer:
        servers:
        {{- range .Endpoints }}
        - url: http://{{ . }}:8080/
        {{- end }}

