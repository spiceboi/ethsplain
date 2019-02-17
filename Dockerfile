FROM golang:latest
WORKDIR /go/src/github.com/spiceboi/ethsplain/
COPY . .
RUN go get -d
CMD go run main.go