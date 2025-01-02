FROM golang:1.23

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod tidy

COPY . .
COPY .env .

RUN go build -o e-wallet-wallet
RUN chmod +x e-wallet-wallet
EXPOSE 8081
CMD [ "./e-wallet-wallet" ]