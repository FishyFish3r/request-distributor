FROM golang:latest

WORKDIR /dumpservice

COPY go.mod go.sum ./

RUN go mod download

COPY . . 

RUN go build -o serv

CMD [ "./serv" ]