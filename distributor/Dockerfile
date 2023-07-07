FROM golang:latest

WORKDIR /distributor

COPY go.mod go.sum ./

RUN go mod download

COPY . . 

RUN go build -o dist

CMD [ "./dist" ]