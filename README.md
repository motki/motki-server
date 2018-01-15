# motki-server

The MOTKI application server.

[![GoDoc](https://godoc.org/github.com/motki/motki-server?status.svg)](https://godoc.org/github.com/motki/motki-server)

## Getting started

This repository contains the `motkid` source code, including dependency source code (stored in `vendor/`).

#### Pre-requisites for building

* [A recent Go compiler](https://golang.org)
* [go-bindata](https://github.com/jteeuwen/go-bindata)
  ```bash
  go get -u github.com/jteeuwen/go-bindata/...
  ```
* postgres9 client (psql and pg_restore)
* protoc and proto-gen-go
* cat, awk, sed, grep, curl, bunzip2, unzip, git

##### Requirements for running

* An available PostgreSQL database.
* A SMTP provider.
* A machine to run motkid on.


## Install using `make`

[Download a copy of this repository](https://github.com/motki/motki-server/archive/master.zip) and verify that the Makefile works.

```bash
curl -L -o motkid-master.tar.gz https://github.com/motki/motki-server/archive/master.tar.gz
tar xzf motkid-master.tar.gz
cd motkid-master
make debug
```

Assuming you haven't copied `config.toml.dist` to `config.toml`, you will be greeted with an error.

```
Makefile:61: *** config.toml does not exist. Copy config.toml.dist and edit appropriately, then try again..  Stop.
```

Once you've copied `config.toml`, you can actually build the program. The simplest way is `make`.


### Makefile reference

|  Target       | Description 
|-----------    |---------------------------------------------------
| build         | Build the `motkid` binary.
| install       | Installs database schemas and EVE static dump data.
| uninstall     | Drop created database schemas and EVE static dump data.
| clean         | Delete all build files.
| generate      | Runs `go generate`.
| matrix        | Build a matrix of arches and OSes, see below.
| download      | Download EVE static dump data.
| assets        | Installs EVE static dump data.
| db            | Install the required schema definitions into the database.
| schema_evesde | Installs the EVE static dump schema definitions.
| schema_app    | Installs the app schema definitions.

> Note that Makefile targets that deal with the database connect to the server defined in `config.toml`

#### Notes

Build for a specific OS and arch
```bash
make build GOOS=linux GOARCH=arm7
```

Cross-compile the binaries for many platforms at once
```bash
make matrix ARCHES="amd64 arm6 arm7 386" OSES="windows linux darwin"
```


## Manual Installation

Clone or [download the repository](https://github.com/motki/motki-server/archive/master.zip).

This application does not rely on `$GOPATH`, but if you are planning on making changes, it may help to put it there.

Below is an example of one way to get the code.

> This assumes you have a simple `$GOPATH` with only one value (and no colons in it)

```bash
mkdir -p $GOPATH/src/github.com/motki
git clone git@github.com:motki/motki-server $GOPATH/src/github.com/motki/motki-server
cd $GOPATH/src/github.com/motki/motki-server
```


### Install resources

Load the data in the `resources` folder.

1. Un-bzip the `evesde-*-postgres.dmp.bz2`.
2. Use `pg_restore` to load the EVE static dump.
   > Warnings abouts a missing "yaml" role can be ignored.
3. Extract the Icons and Types zips to `public/images` (creating `public/images/Icons` and `public/images/Types`)


## Configuration

Copy `config.toml.dist` to `config.toml` and edit appropriately.

### Configuring the EVE API

To use the EVE API you need to set up an Application at the [EVE Developer Portal](https://developers.eveonline.com/applications).  You'll need to select appropriate roles (*all* of them is fine) and then set a correct Return URL for your setup.

> Note: the Return URL can include a port specification.

Once you have created your application on the developer portal, put the Client ID, Secret, and Return URL in the corresponding section in `config.toml`.


### Configuring SSL

You need to generate a certificate and private key to properly set up SSL. During development, a self-signed certificate is recommended. For production deployments, the process is made simpler by using Let's Encrypt to automatically generate a valid certificate.

> Using [Let's Encrypt](https://letsencrypt.org/) is highly recommended. [See this section for more information](https://github.com/motki/motki-server#generating-a-cert-with-lets-encrypt) on setting up Let's Encrypt.

#### Generating a cert with Let's Encrypt

1. Configure the SSL section in `config.toml`
    1. Set `autocert=true` in config.toml.
    2. Set `certfile=""` and `keyfile=""` in config.toml
    3. Set the SSL `listen` parameter to a valid public hostname.
2. ...
3. Profit

##### Generating a self-signed cert

1. Copy the source code from this stdlib utility: [generate_cert.go](https://golang.org/src/crypto/tls/generate_cert.go).
2. Put it inside its own package (something like `./cmd/gencert/` in the project directory).
3. Compile and run it: 
   `go run ./cmd/gencert/generate_cert.go --host localhost`
4. There should now be a `key.pem` and `cert.pem` file in the current working directory. Update `config.toml` with the path to these.
5. Start motkid

> Don't commit the `gencert` utility or the generated keys to the source code repository.


### Running `motkid`

Build the `motkid` command using make, then run the resulting executable.

```
$ make clean build
$ build/motkid -h
Usage of build/motkid:
  -conf string
    	Path to configuration file. (default "config.toml")
  -version
    	Display the application version.
```
