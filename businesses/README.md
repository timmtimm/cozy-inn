# How to build mocks

## Prerequisites

Install mockery

```sh
go install github.com/vektra/mockery/v2@latest
```

## Generate mocks

Change to directory that contains interface you want to mock and run the following command to genereate mocks

```sh
mockery --all --keeptree
```