FROM golang:1.22.0 AS builder
WORKDIR /app
COPY go.mod go.sum .
RUN go mod download
COPY GeoIP2-City.mmdb GeoIP2-ISP.mmdb .
COPY *.go .
COPY internal internal
RUN go build .

FROM gcr.io/distroless/base-debian12
WORKDIR /app
COPY --from=builder /app/zmono .
ENTRYPOINT ["./zmono"]
