# server-control-api

Exposes Systemd, Podman as a REST-Api with key-based authentication.

## dev

### starting container locally 

> DOES NOT WORK IN DEV CONTAINER!!

```sh
docker rm -f sca && docker run --privileged=true --security-opt seccomp=unconfined --cap-add=SYS_ADMIN -d -p 3003:3000 -v /run/dbus/system_bus_socket:/run/dbus/system_bus_socket --name sca -e 'ENVIRONMENT=dev' --env DBUS_SESSION_BUS_ADDRESS="unix:path=/run/dbus/system_bus_socket" -e 'DEFAULT_API_KEY=1234' sca
```

confirmed to work in:
* docker
* podman
* podman-systemd (quadlets)

> **ATTENTION**: The developer is not liable for anything that happens when using this software.
> This is potentially insecure to set-up a container this way and to expose system-service as a REST-Api to the outside world.

## api

### authorization

with api key. 

defined as env variable `DEFAULT_API_KEY`

only needed if `ENVIRONMENT` is `prod`

You need to provide the Header in one of two forms:

```
Authorization: <Key>
Authorization: Bearer <Key>
```

**generating an api key**

use the following command

```sh
openssl rand -hex 32
```

### docker

**pull image**

url:
```
POST /docker/images/pull
```

params:

```
image: Param in form of "docker.io/library/busybox" with optional tag (default is latest). With tag "docker.io/library/busybox:latest" 
```

responses:

```
200: Successfully Pulled Image
500: Some Error with Docker Tool occurred

body structure: {message: "", status: "error | success"}
```

**restart container**

url:
```
POST /docker/containers/restart
```

params:

```
name: name or id of container
```

responses:

```
200: Successfully restarted container
500: Some Error with Docker Tool occurred

body structure: {message: "", status: "error | success"}
```

### systemd

**restart service**

url:
```
POST /systemd/restart
```

params:

```
name: name of unit
```

responses:

```
200: Successfully restarted service
500: Error occurred with systemd

body structure: {message: "", status: "error | success"}
```
