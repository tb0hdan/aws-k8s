aws-k8s
===

# Intro

This project is just a wrapper around AWS STS, so most of the things can be done via shell.
I, however, find it more comfortable to have testable and maintainable Go code instead.
`aws-k8s` also does check for a key expiration itself and prompts for re-authentication.


# Usage

## Installation

`go install github.com/tb0hdan/aws-k8s/cmd/aws-k8s@latest`

## Initial setup

- Configure AWS CLI: `aws configure bla bla`
- Configure AWS K8S: `aws-k8s configure`

## Logging in using MFA token

`aws-k8s refresh --token=123456`

after that you can do

`eval $(aws-k8s keys --print)`

and get AWS sessions key in your console.

## Assuming role using MFA token

In most cases one would not need AWS session and all this shell malarkey, so one can just do:

`aws-k8s assume --token=123456`

then add alias to shell of your choice (e.g. bash): `alias kdev="aws-k8s wrap`, open new terminal window and then

`kdev get pods`

## Getting assume role keys

If by any chance you're going to need assume role keys then following command will make them available

`eval $(aws-k8s assume-keys --print)`

## Debugging

`aws-k8s --debug {cmd} parameters...`

e.g.

`aws-k8s --debug keys --print`

# Contributing

- Open PR
- Wait for it get approved or rejected
