# Beagle

This project is composed of a single page web application written in
javascript and a hardware abstraction backend written in go which communicate
over a HTTP REST interface. However it is possible to write your own scripts
directly from the go packages if you feel the desire to.

In case you are sceptic about the choice of using go as backend language,
let me give you a small sales pitch of why go seems to be a better fit
as for example python in my opinion.
Go is a typed language which is very opinionated about how to write your code.
For example go enforces you to document exported variables, automatically
formats your code and does not give you a lot of options in designing your
interfaces, i.e. you don't have the choice between functional and object
oriented programming as python often offers. This can be admittedly annoying
in the beginning, however it pays off when you consider an environment in
which different people from different backgrounds work together and where there
are only few guarantees that the person responsible for a piece of code is
still available any time later. It also prevent people who have not much
experience with software engineering from leaving back a mess and supports them
with a good development toolset integrated into the development process which
for example gives you hints or exports your documentation as HTML. Besides being
opinionated there is not much magic happening when you write go code.
The second argumentation branch I would like to mention originates from the
fact that go is a typed language which offers cross compile support, thus it
is possible to write code on your workstation and compile it to your target
ARM platform with much less potential for debugging endeavors. Finally I would
like you to give the language a chance.

## Usage

There are different ways to use this project. Starting from the highest level
we describe the most common usages.

### App

Using the app will give you a graphical user interface which you can access
from your local network. At the time being it supports control of the direct-
digital-synthesizer in single tone (constant frequency, amplitude and phase)
and (frequency) sweep mode.

![Screenshot](https://user-images.githubusercontent.com/1780466/38424115-4b91f7a2-39b0-11e8-87cd-ba9eb11f30d6.png)

### REST

The REST interface is a HTTP service which follows the representational state
transfer conventions common for HTTP APIs. The idea is to represent a family
of objects as a resource which we can interact with the corresponding HTTP
methods. At the time of writing serialization format used is JSON, however
XML support (built into go) can be added in the future by adding additional
content accept header handlers. We will now discuss the available resources.

1. Specs

The `/specs` resource only supports a `GET` request and will return the
registered device specifications. These can be extended in future to support
different device classes or in present if we want to validate user input
against the operational specifications. From the command line we can use the
`curl` tool available on most UNIX machines.

```shell
curl -X GET \
     -H "Accept: application/json" \
     http://localhost:8000/specs
```

2. Devices

The `/devices` resource supports a `GET` request which will return the
configured state of all available devices. Further a `PATCH` to `/devices/0`
can be used to amend device parameters i.e. frequency.

```shell
curl -X PATCH \
     -H "Content-Type: application/json" \
     -d '{"singleTone":{"frequency":100000000}}'\
     http://localhost:8000/devices/0
```

### Driver

In case the HTTP interface does not offer enough flexibility i.e. you would like
to operate the direct-digital-synthesizer devices in RAM mode or you just
don't want to have the overhead of a running HTTP server, you can also directly
use the driver. You should find the necessary information from the datasheet
and the documentation inside the device package.

## Development

If you want develope on this project or you want to have recompile certain
tools for a new release read the following section.

### Frontend

The app uses the [parcel](https://parceljs.org) bundler to transpile modern,
modularized javascript code to output known by modern web browsers. The bundler
runs on [node](https://nodejs.org), thus it is a prerequisite to have node
installed. From there on you can install all necessary packages for bundling
through `npm install` and start the development server with `npm run serve`.
The development server supports hot module replacement, hence you should be
able to edit your component and see instant changes in your web browser.

### Backend

For the backend you need to install [go](https://golang.org/doc/install) and
the package manager [dep](https://golang.github.io/dep/). You may need to
setup a `$GOPATH` in your environment variables. From that on a simple
`dep ensure` will fetch dependencies and a `go run cmd/http/main.go` executes
the command scripts.

## Appendix

### Flash Image

Update your Beaglebone image by writing a flasher image to the micro SD card

```shell
curl -O https://rcn-ee.com/rootfs/bb.org/testing/2018-03-25/stretch-lxqt/BBB-blank-debian-9.4-lxqt-armhf-2018-03-25-4gb.img.xz
xzcat BBB-blank-debian-9.4-lxqt-armhf-2018-03-25-4gb.img.xz | sudo dd of=/dev/disk2
```

and start the flash process by inserting the SD card into the
unpowered Beaglebone, press the S2 button and power on.
