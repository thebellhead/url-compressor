FROM golang:1.22

WORKDIR /usr/src/app

COPY ./go.mod ./
COPY ./go.sum ./
RUN go mod download && go mod verify

COPY ./backend ./backend
RUN mkdir -p /usr/local/bin/
RUN go mod tidy
RUN go build -v -o /usr/local/bin/app ./backend/cmd

EXPOSE 1228

CMD ["app"]