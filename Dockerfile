FROM golang:1.15-alpine AS build

WORKDIR /build
RUN apk add gcc musl-dev
COPY cmd ./cmd
COPY internal ./internal
COPY go.mod .
COPY go.sum .
RUN go mod download
RUN ls -la
RUN go build -o backend ./cmd/app/main.go


# -------------------------

FROM alpine:3 as final
COPY --from=build /build/backend ./backend

ENTRYPOINT ["./backend"]
CMD [ "-c" "config.json"]