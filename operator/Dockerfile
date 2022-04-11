FROM golang:1.17 AS builder

ENV GO111MODULE=on \
      CGO_ENABLED=0 \
      GOOS=${TARGETOS} \
      GOARCH=${TARGETARCH} \
      GOGC="" \
      GOPROXY=https://proxy.golang.org\
      XDG_CACHE_HOME=/go/cache \
      GOBIN=/dist

WORKDIR /app

# COPY go.mod /
# COPY go.sum /

# RUN ls -lah
# RUN pwd

# RUN go mod download

COPY . .

RUN ls -lah
RUN pwd

RUN go build -ldflags "-s -w" -o trident_orchestrator


RUN ls -lah
RUN pwd

FROM alpine:3.12

ARG PORT=8000
ENV PORT $PORT
EXPOSE $PORT
ARG K8S=""
ENV K8S $K8S
ENV TRIDENT_IP localhost
ENV TRIDENT_SERVER 127.0.0.1:$PORT

COPY --from=builder /app/trident_orchestrator /


ENTRYPOINT ["/trident_orchestrator"]