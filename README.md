# server-control-api

Exposes Systemd, Podman as a REST-Api with key-based authentication.

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