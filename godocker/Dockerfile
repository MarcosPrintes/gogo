FROM golang:1.12.5

LABEL maintener="marcosprintes"

WORKDIR $GOPATH/src/github.com.br/MarcosPrintes/godocker 

COPY . .

RUN go get -d -v ./...

RUN go install -v ./...

EXPOSE 8080

ENTRYPOINT [ "godocker" ]