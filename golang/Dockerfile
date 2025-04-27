FROM golang
WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . ./

RUN go build -o /docker-gs-ping

EXPOSE 8080

CMD [ "/docker-gs-ping" ]
