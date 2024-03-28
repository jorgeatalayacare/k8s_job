# Step 1: Builder stage
FROM golang:latest AS builder
WORKDIR /app
COPY go.mod ./
RUN go mod download
COPY . .
RUN go test ./pkg/...
RUN CGO_ENABLED=0 GOOS=linux go build -o file-reader ./cmd

FROM scratch
WORKDIR /myapp
COPY --from=builder /app/file-reader .
COPY --from=builder /app/pkg/file.txt .
ENTRYPOINT ["./file-reader"] 
