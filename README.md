aws-k8s
===

# Intro

This project is just a wrapper around AWS STS, so most of the things can be done via shell.
I, however, find it more comfortable to have testable and maintainable Go code instead.
`aws-k8s` also does check for key expiration itself and prompts for re-authentication.


# Usage

## Installation

`go get github.com/tb0hdan/aws-k8s`

## Initial setup

`aws configure bla bla`
`aws-k8s configure`

## Logging in

`aws-k8s refresh --token=123456`

after that you can do

`eval $(aws-k8s keys --print)`

and get AWS sessions key in your console.

## Assuming role

In most cases one would not need AWS session and all this shell malarkey, so one can just do:

`aws-k8s assume --token=123456`

and then just

`aws-k8s wrap get pods`

# Contributing

- Open PR
- Wait for it get approved or rejected
