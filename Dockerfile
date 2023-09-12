FROM golang:latest AS builder
WORKDIR /src
COPY ./ ./
RUN go build -o biling cmd/biling/main.go


FROM golang:latest
WORKDIR /src
COPY --from=builder /src/biling ./
COPY --from=builder /src/json-file /opt/json-file
CMD ["./biling"]