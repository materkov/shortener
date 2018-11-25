FROM golang:1.11.1-alpine3.8
RUN apk add --no-cache git gcc musl-dev
WORKDIR /build
COPY . /build
RUN go build -o /app cmd/main.go

FROM alpine:3.8
EXPOSE 8000
COPY --from=0 /app /shortener-analytics
CMD ["/shortener-analytics"]
