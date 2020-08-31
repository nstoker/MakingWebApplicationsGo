# Making Web Applications With Go

Following a [Pluralsight course](https://app.pluralsight.com/library/courses/creating-web-applications-go-update/exercise-files).

Some changes.

## Environment

Copy the `example.env` to `.env` and put your port, password salt, and database details in. Apart from the port, don't check these values into your database.

## TLS

To enable local tls testing run

```bash
go run ~/go/src/crypto/tls/generate_cert.go -host localhost
``

and the `cert.pem` and `key.pem` files will be created. Do not check these into your repo.
