# RESAS-cli

# setup
## golang
it is developed in golang:1.20.1.
if you are using asdf, you can install it by `asdf install`, otherwise install it from official page.

## tools
install tools for development

```
make tools.get
```

## env vars
set the following environment variables

|name|value or description|
|-|-|
|RESAS_API_KEY| api key for RESAS api|
|RESAS_API_ENDPOINT|`https://opendata.resas-portal.go.jp`|

## ready to go
now you are ready ;)
you can run 
```
go run . population
```


# development
# linting
lint check
```
make lint
```
