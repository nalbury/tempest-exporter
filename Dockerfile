FROM golang:1.16-buster

ADD ./ /src/

WORKDIR /src/

RUN go build -o ./tempest-exporter ./

FROM debian:buster-slim

RUN mkdir /tempest-exporter/
COPY --from=0 /src/tempest-exporter /bin/tempest-exporter

ENV tempest-exporter_PATH=/tempest-exporter/
ENTRYPOINT ["/bin/tempest-exporter"]
