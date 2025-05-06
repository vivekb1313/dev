FROM golang:1.20 as builder
WORKDIR /app
COPY . .
RUN go build -o url-checker
FROM alpine
WORKDIR /app
COPY --from=builder /app/url-checker .
CMD ["./url-checker"]