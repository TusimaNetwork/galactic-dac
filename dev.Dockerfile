# CONTAINER FOR RUNNING BINARY
FROM cdk-near-base
WORKDIR /app
COPY ./cmd/cmd /app/cdk-data-availability
COPY ./db/migrations /app/migrations

EXPOSE 8444
CMD ["/bin/sh", "-c", "/app/cdk-data-availability run --cfg /app/config/dac-config.toml"]
