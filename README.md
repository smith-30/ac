# acc

[![CircleCI](https://circleci.com/gh/smith-30/acc/tree/master.svg?style=svg)](https://circleci.com/gh/smith-30/acc/tree/master)
[![Coverage Status](https://coveralls.io/repos/github/smith-30/acc/badge.svg?branch=master)](https://coveralls.io/github/smith-30/acc?branch=master)
[![GoDoc](https://godoc.org/github.com/smith-30/acc?status.svg)](https://godoc.org/github.com/smith-30/acc)
[![license](https://img.shields.io/badge/license-MIT-4183c4.svg)](https://github.com/smith-30/acc/blob/master/LICENSE)

**a**ll **c**lear **c**heck

## Install

```
$ go get -u github.com/smith-30/acc
```

## Contribution

### Add other contest

clone this repository and add command.

```bash
$ cobra add <contest> # ex. aoj 
```

write the logic as `contest.go` is created in the cmd directory.