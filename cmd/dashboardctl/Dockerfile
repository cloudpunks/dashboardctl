FROM --platform=$BUILDPLATFORM ghcr.io/jdx/mise:2026.11.12-debian@sha256:9ab8d63389bf7d787f0b762e02f7b8fd5ada3d076d3e808a1f818ece155b6fe7 AS builder
WORKDIR /src

COPY mise.toml ./
RUN --mount=type=cache,target=$MISE_CACHE_DIR \
   mise trust && mise install

COPY . .
RUN --mount=type=cache,target=/root/go/pkg \
    go mod download -x

ARG TARGETOS
ARG TARGETARCH

RUN --mount=type=cache,target=/go/pkg \
    --mount=type=cache,target=/root/.cache/go-build \
    task generate build GOOS=${TARGETOS} GOARCH=${TARGETARCH}

FROM alpine:3.24@sha256:f59abe656b692d901e184be3445e084786602108599d2a51d229c0f5a48c71bf

RUN apk add --no-cache ca-certificates mailcap

RUN apk add --no-cache ca-certificates mailcap && \
    addgroup -g 1337 dashboardctl && \
    adduser -D -u 1337 -h /var/lib/dashboardctl -G dashboardctl dashboardctl

ENTRYPOINT ["/usr/bin/dashboardctl"]
CMD ["help"]

COPY --from=builder /src/bin/dashboardctl /usr/bin/dashboardctl
WORKDIR /var/lib/dashboardctl
USER dashboardctl
