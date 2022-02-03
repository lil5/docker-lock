# Example running in docker compose file

```yml
version: "3"

services:
  lock:
    image: safewaters/docker-lock:0.8.10
    volumes:
      - $PWD/docker-compose.yml:/run/docker-compose.yml:ro
      - $PWD/config/radicale/Dockerfile:/run/config/radicale/Dockerfile:ro
      - $PWD/docker-lock.json:/run/docker-lock.json
    environment:
      - PWD=/run/
    command: lock generate
```

# Explanation

The environment variable `PWD` is used to be able to add file volumes directly to the container.

Make sure to backup your `docker-lock.json` file before switching off your server.

If you need to start from the old versions run this first. If there are any changes, alter the version in the docker-compose file.

```sh
docker run \
-v $PWD:/run \
-v $PWD/docker-lock.json:/run/docker-lock.json \
-e PWD=/run/ \
safewaters/docker-lock:0.8.10 lock verify
```
