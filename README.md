# Ping Default Gateway command

A very simple command: Ping the default gateway

This command pings the default gateway. It gets the Default gateway and passes it to the ping command.

## Usage:

`ping-dg`

If your Default Gateway is, for example, `192.168.1.254` (a very common one for residential routers), then it should
be equivalent of calling `ping 192.168.1.254`. On unix it has a limit of 4 pings

## Install:

You can either install it by using `go get eacp.dev/ping-dg` or by downloading the binary file from the realeases section of this 
repository
