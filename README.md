# teleport-go-utils

teleport-go-utils

# author

Tomer Tcherniak

# project

service has been created for https://goteleport.com/teleport-connect-2022/

# info

```
The reason behind this service is to supply the end users script that can be installed if needed
Script will be built using export AUTH_SERVER , JOIN_TOKEN , CA_PIN_HASHES , INSTALL_FILE parameters;

The Go Service has 3 endpoints:

1. download script file
2. get auth server
3. get ca server

JOIN TOKEN endpoint has to be created
```

# env variables exammple

```
export AUTH_SERVER="teleport.server.com"
export PORT=9999
export CA_PIN_HASHES=dGVsZXBvcnQuc2VydmVyLmNvbQo= # echo teleport.server.com | base64
```

# env variables exammple

```
curl -s --header "Teleportfile: teleport2022" http://localhost:9999/download
curl -s --header "Teleporturl: teleport2022" http://localhost:9999/auth_server | base64 -d
curl -s --header "Teleportca: teleport2022" http://localhost:9999/ca_pin_hashes | base64 -d
```

# create service in the local machine
run build.sh and services.sh

## How teleport-go-utils works
![How teleport-go-utils works image](teleport-go-utils.png)
