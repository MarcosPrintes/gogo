FROM golang:1.12.5


WORKDIR $GOPATH/src/github.com.br/MarcosPrintes/nms

COPY . .

RUN go get -d -v ./...

RUN go install -v ./...

EXPOSE 8083

ENTRYPOINT ["./nms"]

