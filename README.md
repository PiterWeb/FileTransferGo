# FileTransferGo

## Packages used :package:

* Go:
  * Gin (*http server*) :globe_with_meridians:
  * Cobra (*CLI command framework*) :snake:
* External :
  * Ngrok (*secure tunnels*) :lock:

## Features 

* :bust_in_silhouette: Client
  * Send Files to another computer using command line
  
* :computer: Server
  * Listen for files from other computers across the internet using Ngrok secure tunnels
  
*  :bookmark: On every subcommand there is -h flag to show description of all commands
  
## :books: How to use

*[Get yout Ngrok Authtoken](https://dashboard.ngrok.com/get-started/your-authtoken "Go to Ngrok")*

1. Set authtoken if u haven't use ngrok **```$ ngrok authtoken <authtoken> ```**  (run it on project folder)

2. Run CLI from Go **```$ go run . <Commands>```** 

- - - -

2. Run CLI from .exe file **```$ FileTransfer.exe <Commands>```**  (If u did the build)

## :construction_worker: Build

**```$ go build```** (on project directory)

## 💬 Commands

To see all commands run :

**```$ go run .```** 

- - - -

**```$ FileTransfer.exe```** (If u did the build)

