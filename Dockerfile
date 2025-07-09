FROM golang:1.22.5 AS builder
WORKDIR /app
COPY go.mod ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o cloudrun

FROM gcr.io/distroless/static-debian11
WORKDIR /
COPY --from=builder /app/cloudrun /
CMD ["/cloudrun"]