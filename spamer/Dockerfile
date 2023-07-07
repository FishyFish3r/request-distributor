FROM golang:latest

WORKDIR /spamer

COPY go.mod go.sum ./

RUN go mod download

COPY . . 

RUN go build -o spamer

CMD [ "./spamer" ]