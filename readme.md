# Mini Wallet Backend using Golang

This is an implementation MiniWallet service in Golang with Clean Architecture Pattern approached.

### Feature

- [x] Initialize Wallet
- [x] Get Balance
- [x] Enable Wallet
- [x] Disable Wallet
- [x] Deposit/Top Up
- [x] Withdrawal

All data stored in no-relation db [sqlite](https://www.sqlite.org/).

## Command

- ### How to Install

  - Install node modules

  ```sh
  $ go get .  
  ```

  - Build application

  ```sh
  $ go build -o main 
  ```

  - Run application

  ```sh
  $ go run main.go || ./main
  ```

<p>Default will run in port <b>:8181</b>, you may change port in .env</p>
