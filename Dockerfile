FROM golang:1.17 AS builder

ENV GO111MODULE=on \
      CGO_ENABLED=0 \
      GOOS=${TARGETOS} \
      GOARCH=${TARGETARCH} 
# GOGC="" \
# GOPROXY=https://proxy.golang.org\
# XDG_CACHE_HOME=/go/cache

COPY . .

RUN ls

RUN go build -ldflags "-s -w" -o trident_orchestrator

RUN go build -ldflags "-s -w" -o tridentctl

RUN go build .

RUN chwrap/make-tarball.sh /chwrap chwrap.tar


FROM alpine:3.12

ARG PORT=8000
ENV PORT $PORT
EXPOSE $PORT
ARG K8S=""
ENV K8S $K8S
ENV TRIDENT_IP localhost
ENV TRIDENT_SERVER 127.0.0.1:$PORT

COPY --from=builder /trident_orchestrator /
COPY --from=builder /tridentctl /bin/tridentctl
COPY --from=builder /trident /
COPY --from=builder /chwrap.tar /

ENTRYPOINT ["/bin/tridentctl"]
