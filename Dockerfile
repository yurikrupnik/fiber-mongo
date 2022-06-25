#FROM golang:1.18-buster AS build
#
#WORKDIR /app
#
#COPY go.mod ./
#COPY go.sum ./
#RUN go mod download
#
#COPY . .
#
#RUN go build -o /docker-gs-ping

##
## Deploy
##
#34.7MB
#FROM gcr.io/distroless/base-debian10
#24.4MB
FROM alpine:latest
WORKDIR /

#COPY --from=build /docker-gs-ping /docker-gs-ping
COPY ./fiber-mongo /docker-gs-ping

EXPOSE 8080

#USER nonroot:nonroot

ENTRYPOINT ["/docker-gs-ping"]