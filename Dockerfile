FROM golang:onbuild

WORKDIR /go/src/github.com/blankbook/writecontent
COPY . .

RUN go-wrapper download
RUN go-wrapper install

CMD ["go-wrapper", "run"]
