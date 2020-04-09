FROM golang:1.13-buster

WORKDIR /go/src/gen-cover

COPY . .

RUN make release

CMD "release/gen-cover/bin"
