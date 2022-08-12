FROM golang:1.18

WORKDIR /sub_app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./

RUN go build -o /subapp .

EXPOSE 6001

CMD [ "/subapp" ]