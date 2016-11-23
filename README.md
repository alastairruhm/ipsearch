[![Build Status](https://travis-ci.org/alastairruhm/ipsearch.svg?branch=master)](https://travis-ci.org/alastairruhm/ipsearch) [![Coverage Status](https://coveralls.io/repos/github/alastairruhm/ipsearch/badge.svg?branch=master)](https://coveralls.io/github/alastairruhm/ipsearch?branch=master) [![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](https://opensource.org/licenses/MIT)

# ipsearch

a CLI tool searching ip location based on api of https://www.ipip.net/

## Installation

Use [`go get`](https://golang.org/cmd/go/#hdr-Download_and_install_packages_and_dependencies) to install and update:

```sh
$ go get -u github.com/alastairruhm/ipsearch
```

## Usage

From the commandline, `ipsearch` provides the functions as follows.By default, it prints its output to `stdout`.

example: 

```sh

# IP
$ ipsearch 8.8.8.8
8.8.8.8 GOOGLE-GOOGLE-google.com

# host
$ ipsearch github.com
ipsearch github.com
lookup host: github.com
192.30.253.112 美国-弗吉尼亚州-阿什本
192.30.253.113 美国-弗吉尼亚州-阿什本

# URL
$ ipsearch https://github.com/alastairruhm/ipsearch
lookup host: github.com
192.30.253.112 美国-弗吉尼亚州-阿什本
192.30.253.113 美国-弗吉尼亚州-阿什本
```

## License

`ipsearch` is released under the [MIT License](https://opensource.org/licenses/MIT).