FROM golang:1.19-alpine AS build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .
ARG GOOS=linux
ARG GOARCH=amd64
RUN go build -o /fiber-mongo

##
## Deploy
##
#34.7MB
#FROM gcr.io/distroless/base-debian10
#24.4MB
#19MB
#FROM alpine:latest
#13.8MB
#FROM scratch AS final
FROM alpine:latest AS final
WORKDIR /

COPY --from=build /fiber-mongo /app
#COPY ./fiber-mongo /app

EXPOSE 8080

#USER nonroot:nonroot

ENTRYPOINT ["/app"]