FROM golang:1.18-alpine as builder
# Define build env
ENV GOOS linux
ENV CGO_ENABLED 0
# Add a work directory
WORKDIR /app
# Cache and install dependencies
COPY go.mod go.sum ./
RUN go mod download
# Copy app files
COPY . .
# Build app
RUN go build -o dockeringo

FROM alpine:3.14 as production
# Add certificates
RUN apk add --no-cache ca-certificates
# Copy built binary from builder
COPY --from=builder /app/dockeringo .
# Expose port
EXPOSE 9808
# Exec built binary
RUN chmod +x ./dockeringo
CMD ["./dockeringo"]
