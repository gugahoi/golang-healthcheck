# golang-healthcheck

[![Build Status](https://travis-ci.org/gugahoi/golang-healthcheck.svg?branch=master)](https://travis-ci.org/gugahoi/golang-healthcheck)
[![Coverage Status](https://coveralls.io/repos/github/gugahoi/golang-healthcheck/badge.svg?branch=master)](https://coveralls.io/github/gugahoi/golang-healthcheck?branch=master)
[![Maintainability](https://api.codeclimate.com/v1/badges/6c6ee4069794a0141aa2/maintainability)](https://codeclimate.com/github/gugahoi/golang-healthcheck/maintainability)
[![Go Report Card](https://goreportcard.com/badge/github.com/gugahoi/golang-healthcheck)](https://goreportcard.com/report/github.com/gugahoi/golang-healthcheck)

Basic healthcheck application with a web api style endpoint.

It has the following implemented endpoint:

- A `/healthcheck` which returns basic information about your application: Applications Version, Description, Last Commit SHA.
