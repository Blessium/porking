FROM golang:buster

WORKDIR /usr/src/porking

RUN apt-get update -y 

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -v -o /usr/local/bin/porking ./

EXPOSE 1234

CMD ["porking"]
