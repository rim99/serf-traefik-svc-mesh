# README
## Context

This config assumes 2 virtual machines with IP address:

VM | IP Address
:-: |:-
vm1 | 192.168.122.128
vm2 | 192.168.122.116

## Test

Send request to vm2 for app: `webapp-2`

```
$> curl "192.168.122.116:18080/ping?remote=webapp-1" -H "host: webapp-2"
Hello, from webapp-1!
```

