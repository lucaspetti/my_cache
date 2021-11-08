FROM golang:latest as builder
ADD . /app
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags "-w" -a -o ./my_cache .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
COPY --from=builder /app/my_cache ./
COPY --from=builder /app/.env ./
RUN chmod +x ./my_cache
EXPOSE 5000
