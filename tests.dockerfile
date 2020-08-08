FROM golang:1.14

ENV GO111MODULE=on
RUN go get github.com/cucumber/godog/cmd/godog@v0.10.0

CMD godog