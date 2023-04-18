FROM golang:1.19-alpine

WORKDIR /app

COPY . .

RUN go mod download
RUN apk update && apk add bash && bash setup-env.sh docker
RUN go build -o todo-app

EXPOSE 3030

CMD [ "./todo-app" ]