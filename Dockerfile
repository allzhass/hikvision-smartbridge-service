FROM golang:1.23.6-alpine as build
LABEL authors="allzhass"

WORKDIR /build
COPY . /build/

RUN go mod download
RUN apk add bash git gcc gettext musl-dev

# Build your Go application
RUN go build -o app ./cmd/main.go

FROM alpine

ENV TZ=Asia/Almaty
WORKDIR /

COPY --from=build build/app /app

ENV SERVER_PORT=8081
#ENV SMART_BRIDGE_URL=http://10.245.12.102:80/bip-sync-wss-gost/
ENV SMART_BRIDGE_URL=http://10.61.40.133/shep/bip-sync-wss-gost/

EXPOSE 8081

RUN chmod +x /app
CMD ["/app"]
