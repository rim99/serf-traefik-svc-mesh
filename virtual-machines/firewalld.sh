# For RHEL-based Linux distributions
firewall-cmd --add-port 7946/tcp --permanent
firewall-cmd --add-port 18080/tcp --permanent
firewall-cmd --reload

