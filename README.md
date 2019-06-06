# acc

[![CircleCI](https://circleci.com/gh/smith-30/acc/tree/master.svg?style=svg)](https://circleci.com/gh/smith-30/acc/tree/master)
[![Coverage Status](https://coveralls.io/repos/github/smith-30/acc/badge.svg?branch=master)](https://coveralls.io/github/smith-30/acc?branch=master)
[![GoDoc](https://godoc.org/github.com/smith-30/acc?status.svg)](https://godoc.org/github.com/smith-30/acc)
[![license](https://img.shields.io/badge/license-MIT-4183c4.svg)](https://github.com/smith-30/acc/blob/master/LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/smith-30/acc)](https://goreportcard.com/report/github.com/smith-30/acc)

**a**ll **c**lear **c**heck

## Install

```
$ go get -u github.com/smith-30/acc
```

## Usage

### atcoder

```
$ acc atcoder -u https://atcoder.jp/contests/<problem path> -c "go run path/to/main.go"
```

after above command, cache db is created at `$GOPATH/bin/acc.db`. this is to ensure that acc does not load the site.

**example**

```
$ acc atcoder -u https://atcoder.jp/contests/abc119/tasks/abc119_b -c "go run ./main.go"
start atcoder test v0.0.0-e0ce3ec (built with go1.12.1)

Case [0] exp: 48000.0
	ok!
Case [1] exp: 138000000.0038
	ok!
```

## Contribution

### Add other contest

clone this repository and add command.

```bash
$ cobra add <contest> # ex. aoj 
```

write the logic as `contest.go` is created in the cmd directory.