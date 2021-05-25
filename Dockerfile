FROM golang:alpine AS builder

WORKDIR /build

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o go-webhook-job

CMD [ "sh" ]

# Final Image
FROM scratch

COPY --from=builder /build/go-webhook-job     /bin/
COPY --from=builder /lib/ld-musl-x86_64.so.1  /lib/

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

ENTRYPOINT [ "/bin/go-webhook-job" ]