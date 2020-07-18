FROM golang:1.14.4-alpine
RUN apk add --no-cache make ca-certificates
WORKDIR /app/awanku
COPY go.mod go.sum /app/awanku/
RUN go mod download
COPY . /app/awanku/
RUN make build

FROM alpine:3
WORKDIR /app/awanku
COPY --from=0 /app/awanku/dist/core-api .
CMD /app/awanku/core-api
