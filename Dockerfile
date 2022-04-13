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

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN mkdir bin

RUN go build -ldflags "-s -w" -o ./bin/chwrap chwrap/chwrap.go
# RUN apt update && apt install uuid-runtime
# RUN chwrap/make-tarball.sh ./bin/chwrap chwrap.tar

RUN go build -ldflags "-s -w" -o trident_orchestrator
RUN go build -ldflags "-s -w" -o tridentctl

RUN ls -lah
RUN pwd

#### from makefile
FROM ubuntu:latest AS builder2

COPY . .
COPY --from=builder /app/bin/chwrap /bin/chwrap
RUN apt update && apt install -y uuid-runtime util-linux open-iscsi lsscsi multipath-tools procps e2fsprogs rpcbind mount xfsprogs

RUN chwrap/make-tarball.sh /bin/chwrap chwrap.tar
####

FROM gcr.io/distroless/static

ARG PORT=8000
ENV PORT $PORT
EXPOSE $PORT
ARG K8S=""
ENV K8S $K8S
ENV TRIDENT_IP localhost
ENV TRIDENT_SERVER 127.0.0.1:$PORT

COPY --from=builder /app/tridentctl /bin/
COPY --from=builder /app/trident_orchestrator /
COPY --from=builder2 ./chwrap.tar /


ENTRYPOINT ["/bin/tridentctl"]
CMD ["version"]
