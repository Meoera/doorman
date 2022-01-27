FROM golang:1.17-alpine AS build

WORKDIR /build
# RUN apk add gcc=11.2.1_git20220117-r0 musl-dev
COPY cmd ./cmd
COPY internal ./internal
COPY pkg ./pkg
COPY config.json go.mod go.sum /build/
RUN go mod download
RUN go build -o doorman ./cmd/server/main.go


# -------------------------

FROM alpine:3 as final

WORKDIR /app
COPY --from=build /build/doorman ./doorman

ENTRYPOINT ["./backend"]
CMD [ "-c" "config.json"]