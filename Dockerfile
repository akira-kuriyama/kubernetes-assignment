FROM golang:1.22.2-bookworm as builder

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY src/* /app/
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/main

FROM gcr.io/distroless/static

COPY --from=builder app ./
USER 1000
CMD ["./main"]
