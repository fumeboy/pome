iptables -t nat -A OUTPUT -o eth0 -p tcp -d 10.0.0.0/16 -j REDIRECT --to-ports 20011

/pome/sidecar & python3 ./service.py