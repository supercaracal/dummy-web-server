![Deploy](https://github.com/supercaracal/dummy-web-server/actions/workflows/deploy.yaml/badge.svg)

Dummy Web Server
===============================================================================

This is a web application for connectivity test of infrastructure at first deploy.

### Server
```
$ docker run --rm -p 3000:3000 ghcr.io/supercaracal/dummy-web-server
2023/09/23 08:08:34 listen: 0.0.0.0:3000
```

You can specify the port number to listen via an environment variable or a command option.
```
$ docker run --rm -p 3000:3001 -e PORT=3001 ghcr.io/supercaracal/dummy-web-server
2023/09/22 23:27:43 listen: 0.0.0.0:3001

$ docker run --rm -p 3000:3002 ghcr.io/supercaracal/dummy-web-server -port 3002
2023/09/22 23:29:53 listen: 0.0.0.0:3002
```

### Client
Clients can access any paths and get fixed response as a JSON.
```
$ curl http://127.0.0.1:3000/foo/bar/baz
{}

$ curl http://127.0.0.1:3000/ping
{}

$ curl http://127.0.0.1:3000/health
{}

$ curl -IXGET http://127.0.0.1:3000/
HTTP/1.1 200 OK
Content-Type: application/json;charset=UTF-8
Date: Fri, 22 Sep 2023 23:40:56 GMT
Content-Length: 3
```

### Sub command for the health check
You can use the health check feature by the sub command.
```yaml
---
# An example for Docker compose
services:
  server:
    image: ghcr.io/supercaracal/dummy-web-server
    restart: always
    healthcheck:
      test: ["CMD", "/usr/local/bin/dummy-web-server", "health"]
      interval: "10s"
      timeout: "5s"
      retries: 1
    ports:
      - "3000:3000"
```

```
$ go run main.go health
2023/09/23 08:09:30 OK

$ go run main.go health
2023/09/23 08:12:05 NG: Get "http://127.0.0.1:3000/": dial tcp 127.0.0.1:3000: connect: connection refused
exit status 1
```
