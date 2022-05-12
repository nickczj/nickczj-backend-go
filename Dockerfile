#build stage
FROM golang:1.16-alpine AS builder
WORKDIR /go/src/app
COPY . .
RUN go build -o /go/bin/app

#final stage
FROM alpine:3.14
RUN apk --no-cache add ca-certificates tzdata
COPY --from=builder /go/bin/app /app
COPY --from=builder /go/src/app/config /config/
ENTRYPOINT /app
LABEL Name=api Version=0.0.1
EXPOSE 8888
