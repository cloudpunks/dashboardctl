FROM --platform=$BUILDPLATFORM golang:1.27.1-alpine AS builder

RUN apk add --no-cache -U git curl
RUN sh -c "$(curl --location https://taskfile.dev/install.sh)" -- -d -b /usr/local/bin

WORKDIR /go/src/dashboardctl
COPY . /go/src/dashboardctl/

RUN --mount=type=cache,target=/go/pkg \
    go mod download -x

ARG TARGETOS
ARG TARGETARCH

RUN --mount=type=cache,target=/go/pkg \
    --mount=type=cache,target=/root/.cache/go-build \
    task generate build GOOS=${TARGETOS} GOARCH=${TARGETARCH}

FROM alpine:3.24

RUN apk add --no-cache ca-certificates mailcap

ENTRYPOINT ["/usr/bin/dashboardctl"]
CMD ["help"]

COPY --from=builder /go/src/dashboardctl/bin/dashboardctl /usr/bin/dashboardctl
WORKDIR /
