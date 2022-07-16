FROM golang:1.18-buster AS build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN go build -o /fiber-mongo

##
## Deploy
##
#34.7MB
#FROM gcr.io/distroless/base-debian10
#24.4MB
FROM alpine:latest
WORKDIR /

COPY --from=build /app/fiber-mongo /fiber-mongo
#COPY ./fiber-mongo /fiber-mongo

EXPOSE 8080

#USER nonroot:nonroot

ENTRYPOINT ["/fiber-mongo"]