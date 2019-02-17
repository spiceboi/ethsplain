FROM golang:latest
WORKDIR /go/src/github.com/spiceboi/ethsplain/
COPY . .
RUN go get -d
RUN CGO_ENABLED=1 GOOS=linux go build -a -installsuffix cgo -o app .

FROM alpine:latest  
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=0 /go/src/github.com/spiceboi/ethsplain/app .
ENTRYPOINT [ "./app" ]