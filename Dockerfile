FROM golang:1.18.3-alpine
RUN apk add build-base
WORKDIR /app
COPY . .
RUN go build -o /go-backend
EXPOSE 4500
CMD [ "/go-backend" ]