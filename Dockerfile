FROM golang:alpine

WORKDIR /usr/src/porking

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -v -o /usr/local/bin/porking ./

EXPOSE 1234

CMD ["porking"]
