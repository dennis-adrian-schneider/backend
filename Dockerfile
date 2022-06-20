FROM golang:1.18.3-alpine
WORKDIR /app
COPY Auth ./
COPY Database ./
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY *.go ./
RUN go build -o /go-backend
EXPOSE 4500
CMD [ "/go-backend" ]