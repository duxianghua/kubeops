## Running OTel collecot
```sh 
docker run --rm -it -v $(pwd)/collector.yaml:/etc/otelcol/config.yaml -p 4318:4318 -p 9115:9115 otel/opentelemetry-collector:latest-arm64
```