FROM golang:1.14-buster AS prepare

ARG APP_PORT

WORKDIR /go/cache

ADD go.mod .
ADD go.sum .
RUN go mod download

WORKDIR /go/build
COPY . .

FROM golang:1.14-buster AS builder

COPY --from=prepare /go/cache /go/cache
COPY --from=prepare /go/build /go/build

WORKDIR /go/build
RUN go build -o app app.go

# https://docs.docker.com/develop/develop-images/multistage-build/
# https://github.com/phusion/baseimage-docker
FROM gcr.io/distroless/base-debian10

WORKDIR /app/
COPY --from=builder /go/build/app .
COPY --from=builder /go/build/config.json .

EXPOSE ${APP_PORT}
CMD ["./app"]
