FROM golang:alpine3.12 AS builder

ENV GO111MODULE=on \
      CGO_ENABLED=0 \
      GOOS=${TARGETOS} \
      GOARCH=${TARGETARCH} \
      GOBIN=/dist

WORKDIR /app

RUN ls
COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go get -ldflags="-w -s" .

FROM alpine:3.12

ARG PORT=8000
ENV PORT $PORT
EXPOSE $PORT
ARG K8S=""
ENV K8S $K8S
ENV TRIDENT_IP localhost
ENV TRIDENT_SERVER 127.0.0.1:$PORT

COPY --from=builder /dist/trident_orchestrator /
COPY --from=builder /dist/tridentctl /bin/
# ADD chwrap.tar /

ENTRYPOINT ["/bin/tridentctl"]
