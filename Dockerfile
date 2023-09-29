FROM golang:1.8
WORKDIR /go/src/github.com/viniciusventura29/rinha_compilers
COPY main.go .
RUN go build -o rinha_go .

CMD ./rinha_go