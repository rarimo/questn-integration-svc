# questn-integration-svc

The QuestN integration API for certain actions’ verification.

## Specifics

Status Code of the endpoint's response is always 200 and optionally contains error field in the response body.

Example: 
```json
{
  "error": {
    "code": 400,
    "message": "Bad Request"
  },
  "data": {
    "result": false
  }
}
```

## Build

To build the service image locally, there is a shell script `build.sh` that can be used to build the image:

```bash
sh build.sh
```

It will build the image with the tag `questn-integration-svc:latest` which could be used to run the service locally via 
Docker or Docker-Compose.

## Run

To run the service locally you could use Docker-Compose. There is a `docker-compose.yml` file that could be used to run:

```yaml
version: "3.7"

services:
  questn-integration-svc:
    image: questn-integration-svc:latest
    restart: on-failure
    ports:
      - "8000:8000"
    volumes:
      - ./config.yaml:/config.yaml
    environment:
      - KV_VIPER_FILE=/config.yaml
    entrypoint: sh -c "questn-integration-svc run api"
```

**Tip:** 
You need to create config file for the service, [there is example](./config-example.yaml) of it inside the repository.
Inside the `docker-compose.yml` file you need to specify the path to the config file in the `volumes` section or use the
same as in the example if it equals.

To run the service via Docker-Compose you need to run the following command:

```bash
docker-compose up -d
```

The API of the service will be available on the port `8000` of your local machine.

## Testing

To test the service locally after it was started you could use `curl`:

```bash
curl 'http://localhost:8000/questn/v1/poh_status?address=0x457ccef368a14c4d02c4b2a607bfeafc7e06cd5b' \
  -H 'Content-Type: application/json' \
  --compressed
```
