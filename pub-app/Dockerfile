FROM golang:1.18

WORKDIR /pub_app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./

RUN go build -o /pubapp .

EXPOSE 8080

CMD [ "/pubapp" ]