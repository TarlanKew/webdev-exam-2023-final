FROM golang:1.19

WORKDIR /usr/src/app

COPY ./ ./
RUN go get .
RUN export GO111MODULE="on"
RUN export PGHOST=localhost
RUN go build -v -o app map.go

EXPOSE 8080

CMD [ "./app" ]
