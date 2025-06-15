# For RHEL-based Linux distributions
firewall-cmd --add-port 7946/tcp --permanent # for serf
firewall-cmd --add-port 18080/tcp --permanent # for traefik
firewall-cmd --reload

