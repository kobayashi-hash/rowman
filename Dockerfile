FROM golang:1-bullseye AS builder
WORKDIR /work
COPY . .
RUN go build -o rowman cmd/main/rowman.go

FROM debian:bullseye-slim
ARG VERSION=1.0.0

LABEL org.opencontainers.image.source="https://github.com/kobayashi-hash/rowman" \
      org.opencontainers.image.version="${VERSION}" \
      org.opencontainers.image.title="rowman" \
      org.opencontainers.image.description="A CLI tool to process and filter CSV data."

RUN useradd -m -d /workdir nonroot
WORKDIR /workdir
COPY --from=builder /work/rowman /opt/rowman/rowman

USER nonroot
ENTRYPOINT ["/opt/rowman/rowman"]