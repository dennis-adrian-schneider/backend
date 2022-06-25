FROM golang:1.18.3-alpine
RUN apk add build-base
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN go build -o main .
CMD ["/app/main"]