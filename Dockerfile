#デプロイ用のコンテナに含めるバイナリを作成コンテナ
FROM golang:1.18.2-bullseye as deploy-builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -trimpath -ldflags "-w -s" -o app

#----------------------------------------------
#デプロイ用のコンテナ
FROM debian:bullseye-slim as deploy

RUN apt-get update

COPY --from=deploy-builder /app/app .

CMD ["./app"]

#----------------------------------------------
#ローカル開発環境で利用するホットリロード環境

FROM golang:1.18.2 as dev
WORKDIR /app
RUN go install github.com/cosmtrek/air@latest
CMD ["air"]