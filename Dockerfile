FROM golang:1.19-alpine

WORKDIR /app

COPY . .

RUN go mod download
RUN apk update
RUN go build -o run-app

EXPOSE 4080

CMD [ "./run-app" ]