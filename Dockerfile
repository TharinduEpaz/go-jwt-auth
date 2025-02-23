FROM golang:1.23.4-apline

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod tidy

COPY . .

RUN go build -o ./main.go

RUN chmod +x main

EXPOSE 8000

CMD [ "./main" ]