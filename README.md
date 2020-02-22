# About

[![Travis-CI](https://api.travis-ci.org/outofcoffee/base36.svg)](https://travis-ci.org/outofcoffee/base36)

CLI that implements Base36 encoding and decoding.

## Examples

```shell script
# build docker image
make docker

# encode to base36
docker run --rm -it outofcoffee/base36 --encode hello
5PZCSZU7

# decode from base36
docker run --rm -it outofcoffee/base36 --decode 5PZCSZU7
hello
```

## License

[MIT](LICENSE)
