# Build Golang App
FROM golang:1.21 as build

WORKDIR /go/src/app
COPY . .

RUN go mod download

# RUN go install github.com/swaggo/swag/cmd/swag@latest && \
#     swag init # gen docs before building

RUN CGO_ENABLED=0 go build -ldflags "-X main.version=$(git tag --sort=taggerdate | tail -1)" -buildvcs=false -o /go/bin/app

# Now copy it into our base image.
#FROM gcr.io/distroless/static-debian11
FROM ubuntu:latest
RUN apt update && apt install -y dbus 

COPY --from=build /go/bin/app /

ENV ENVIRONMENT="prod"

CMD ["/app"]
