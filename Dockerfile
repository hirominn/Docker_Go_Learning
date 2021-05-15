FROM golang:1.16.4-alpine3.13

# アップデートとgitのインストール！！
RUN apk update && apk add git
# goappディレクトリの作成
RUN mkdir /go/src/goapp
# ワーキングディレクトリの設定
WORKDIR /go/src/goapp
# ホストのファイルをコンテナの作業ディレクトリに移行
ADD ./goapp /go/src/goapp