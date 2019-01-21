FROM golang:1.10 AS build

ENV WORKDIR=/go/src/github.com/3scale/3scale-istio-adapter

ADD . ${WORKDIR}
WORKDIR ${WORKDIR}

RUN go get -u github.com/golang/dep/cmd/dep && \
    dep ensure -v && \
    go build -o /tmp/3scale-istio-adapter cmd/main.go

FROM centos

WORKDIR /app
COPY --from=build /tmp/3scale-istio-adapter /app/
ENV THREESCALE_LISTEN_ADDR 3333
EXPOSE 3333
EXPOSE 8080
ENTRYPOINT ./3scale-istio-adapter

