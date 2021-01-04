# Network Search

A very small project intended for learning Golang, based on [OtavioGallego/curso-golang](https://github.com/OtavioGallego/curso-golang). It is capable of looking up the IPs and Name Servers of a given host.

## Usage

The project expects one of two possible arguments, `ips` or `name-servers`, and both arguments expect the `host` flag. Below you will find usage examples:

```
$ ./network-search ips --host github.com
$ ./network-search server-names --host github.com
```

## Dependencies

Download all project dependencies with the `go get` command.

## Compiling and running

You may compile the project by running `go build` from the root directory, which will generate a binary file called `network-search`. 

It is also possible to invoke Go and run the project with `go run main.gp [ips|name-servers] --host github.com`.
