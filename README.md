# README

## 概要

Go Di demo

## 各パッケージについて

1. cmd
   1. エントリポイント
2. repository
   1. データの永久化などを行う（DB関係）
3. service
   1. データベースを利用する
4. model
   1. モデルを定義

## Test running

**You have to prepare Database**

1. clone this repository
   * `git clone git@github.com:aqyuki/go-di-demo.git`
2. change directory
   * `cd go-di-demo`
3. create `.env` file
   ```env
   dsn=your_database_dsn
   create=true
   ```
   * If you run first, you have to `create` is `true`. Else , `false`
4. run or build
   * `go run ./cmd/main.go`