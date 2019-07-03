# Tencent Cloud CLI
A CLI for managing virtual machines in Tencent Cloud.

## Summary
With this CLI you are able to create, update and delete Cloud Virtual Machine
(CVM) instances in Tencent Cloud.

## Installation
To build the CLI from source you need to have `make` installed on your machine.
You can then build the binaries by running this from from the root directory of
this repository:
```bash
$ make
```

## Configuration
A JSON config file is required to use the CLI, by default this file has to be
placed in `$HOME/.tcloud.json`. This can be overridden with the `-config`
flag.

An example config file is as follows:
```json
{
  "tencent_secret_id": "MY_SECRET_ID_GOES_HERE",
  "tencent_secret_key": "MY_SECRET_KEY_GOES_HERE"
}
```

Alternatively you can set environment variables for the configuration. These
variables must be named as the uppercase versions of those in the example json
above. For example `TENCENT_SECRET_ID=foobar` would set the secret ID to `foobar`.

## Usage
```bash
Usage:
  tcloud [command]

Available Commands:
  help        Help about any command
  images      Commands for interacting with Cloud Virtual Machine images.

Flags:
      --config string   config file (default is $HOME/.tcloud.json)
  -h, --help            help for tcloud
      --region string   The Tencent Cloud API region. See: https://intl.cloud.tencent.com/document/product/213/15692#Region-List (default "eu-frankfurt")
  -t, --toggle          Help message for toggle

Use "tcloud [command] --help" for more information about a command.
```
