# FileTransferGo

## Libraries used

* Gin (*run http server*)
* Cobra (*CLI commands*)

## Features 

* Client 
  * Send Files to another computer using command line
  
* Server
  * Listen for files from other computers across the internet using Ngrok secure tunnels
  
* On every subcommand there is -h flag to show description of all commands
  
## How to use


1.Set account for ngrok **```$ ngrok authtoken <authtoken> ```** 

2.Run CLI from Go **```$ go run . <Commands>```** 

or

2.Run CLI from .exe file **```$ FileTransfer.exe <Commands>```**  (If u did the build)

## Build

**```$ go build```** (on project directory)

## Commands

To see all commands run :

**```$ go run .```** 

or

**```$ FileTransfer.exe```** (If u did the build)

