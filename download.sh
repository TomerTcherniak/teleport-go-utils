#!/bin/bash -e
FAIL=
args=("$@")
if [ -z "${args[4]}" ]; then
   NODENAME="$(hostname)"
else
   NODENAME="${args[4]}"
fi
echo $# arguments passed ${args[0]} ${args[1]} ${args[2]} ${args[3]}
if [ $# -eq 0 ]; then
  echo "Number of Args is zero , please add all arguments" && exit 1;
fi
export AUTH_SERVER="${args[0]}"; export JOIN_TOKEN="${args[1]}" ; export CA_PIN_HASHES="${args[2]}" ;export INSTALL_FILE="${args[3]}" ;

if [ -z "$JOIN_TOKEN" ] ||  [ -z "$AUTH_SERVER" ] || [ -z "$CA_PIN_HASHES" ] || [ -z "$INSTALL_FILE" ]; then

   echo "please rerun he script with var set , ( AUTH_SERVER , JOIN_TOKEN , CA_PIN_HASHES , INSTALL_FILE )"
   FAIL=1
fi
if [ ! -z "$FAIL" ]; then
   echo exiting...
   exit 1
fi
wget -c https://get.gravitational.com/$INSTALL_FILE --no-check-certificate ;  rpm -ivh --force $INSTALL_FILE ;

cat > /etc/teleport.yaml <<EOF
teleport:
  nodename: ${NODENAME}
  auth_token: ${JOIN_TOKEN}
  ca_pin:
  - ${CA_PIN_HASHES}
  auth_servers:
  - ${AUTH_SERVER}:443
  log:
    output: stderr
    severity: DEBUG
auth_service:
  enabled: no
ssh_service:
  enabled: yes
proxy_service:
  enabled: no
EOF

cat > /usr/lib/systemd/system/teleport.service <<EOF
[Unit]
Description=Teleport SSH Service
After=network.target

[Service]
Type=simple
Restart=on-failure
EnvironmentFile=-/etc/default/teleport
ExecStart=/usr/local/bin/teleport start --insecure --config=/etc/teleport.yaml --pid-file=/run/teleport.pid --debug
ExecReload=/bin/kill -HUP $MAINPID
PIDFile=/run/teleport.pid
LimitNOFILE=8192

[Install]
WantedBy=multi-user.target
EOF

systemctl daemon-reload
systemctl enable teleport
systemctl restart teleport
systemctl status teleport
