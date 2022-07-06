# WiiSOAP
WiiSOAP is a server designed specifically to handle Wii Shop Channel SOAP - more specifically, that of the ECommerce library.
Ideally, one day this will become feature complete enough to handle other titles utilizing EC, such as DLCs or other purchases.

It aims to implement everything necessary to provide title tickets, manage authentication, and everything between.

## Setup
WiiSOAP operates on the assumption that you run a PostgreSQL database holding existing tickets.

1. Ensure that your PostgreSQL database contains the schema within `database.sql`.
2. Copy `config.example.xml` to `config.xml` and edit accordingly.
    - Similar to [WSC-Patcher](https://github.com/OpenShopChannel/WSC-Patcher), you may use a base URL of `a.taur.cloud` for localhost development, i.e. via Dolphin.
3. `go build` to create an executable.
4. Run the resulting executable, such as `./WiiSOAP`.

## Contributing
Ensure you have run `gofmt` on your changes.