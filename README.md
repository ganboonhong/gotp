# GOTP - The Golang One-Time Password Library

[![workflow][build-status]](https://github.com/ganboonhong/gotp/actions)
![MIT License][license-badge]

GOTP aims to be the CLI version of Google Authenticator that we use daily to secure our accounts.
As a Vim lover, I enjoy having (most of) my work done without leaving keyboard.


## Why using it
**More efficient**
- In order to get an OTP from your phone when working on computer, we need:
    - Grab a phone
    - Unlock it
    - Find google authenticator app
    - Read the OTP
    - Type it on computer
- Although these action seems can be done less than 30 seconds, repeating it while working on a computer (like what I did) is a pain
- As a developer, it would be better if I can just get my OTP without leaving a computer (even keyboard)
- You can feel the efficiency especially when you need to get different OTPs from different issuers frequently
- Your OTP can be available in 3 seconds by utilizing this application (OTP is copied to clipboard immediately after generated)

**Easy to Backup**
- All your data are stored in single database file (sqlite file), you just need to keep the file in somewhere secure (cloud, hard disk, ...)
	- When you lost your phone, your secret is still available on cloud
	- The [secret](https://github.com/google/google-authenticator/wiki/Key-Uri-Format#secret) of your OTP is encryped in database
	- Even you lost the file, your OTP still won't be accessible by others
- Portable backup file means you can make it avialble in any computers you want (working computer, home computer, need to switch to another computer, ...)

*Caveat:\
Some people might think it goes against the purpose of OTP to store OTP on the same device with the application you’re using.\
Use it if it brings more advantages to you while this doesn't bother you.*

## Prerequisite
`go version ^1.14`

## Installation
```
# Clone source code
$ git clone git@github.com:ganboonhong/gotp.git

# Change to package directory
$ cd gotp

# Get package dependencies
$ go get ./...

# Compile and install package
$ go install

# Setup database
$ gotp app init
```

## Usage


| Create Account    | Delete Account    | Generate OTP 	   |
| ----------------- | ----------------- | ---------------- |
| `gotp otp create` | `gotp otp delete` | `gotp [otp] gen` |

Since generate OTP command is likely to be used frequently, there is a shorthand `gotp gen` for `gotp otp gen`.

## License

GOTP is licensed under the [MIT License][License]


[build-status]: https://github.com/ganboonhong/gotp/actions/workflows/go.yml/badge.svg
[license-badge]:   https://img.shields.io/badge/license-MIT-000000.svg
[License]: https://github.com/ganboonhong/gotp/blob/master/LICENSE
