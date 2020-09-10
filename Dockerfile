FROM golang:latest AS builder

WORKDIR /src
COPY src/main.go .
RUN CGO_ENABLED=0 GOOS=linux go build -o 301-redirect -ldflags '-extldflags "-static"' .

FROM scratch

EXPOSE 8080

COPY --from=builder /src/301-redirect 301-redirect

CMD ["./301-redirect"]