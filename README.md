# WiiSOAP
WiiSOAP is a server designed specifically to handle Wii Shop Channel SOAP - more specifically, that of the ECommerce library.
Ideally, one day this will become feature complete enough to handle other titles utilizing EC, such as DLCs or other purchases.

It aims to implement everything necessary to provide title tickets, manage authentication, and everything between.

This repository has been modified to meet the requirements for WiiMart. This includes functions to handle gifted titles (sending and receiving), handling points (addition and removal), a new database schema to support this, and more. 

## Building
To build WiiSOAP, `git clone` this repository, and do `go build` in the directory it was cloned into.
Make sure you have Go installed, or you won't be able to build it. Get it [here](https://go.dev/learn/)

## Contributing
Ensure you have run `gofmt` on your changes.
