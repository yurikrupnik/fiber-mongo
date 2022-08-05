FROM golang:1.18-buster AS build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .
ARG GOOS=linux
ARG GOARCH=amd64
RUN go build -o /app

##
## Deploy
##
#34.7MB
#FROM gcr.io/distroless/base-debian10
#24.4MB
FROM scratch
WORKDIR /

COPY --from=build /app /app
#COPY ./fiber-mongo /fiber-mongo

EXPOSE 8080

#USER nonroot:nonroot

ENTRYPOINT ["/app"]