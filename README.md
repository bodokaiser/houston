# Houston

Houston provides an interface with your lab equipment.

## Preface

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

Also check the **Appendix** section before!

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

1. Devices

The `/devices/dds` resource supports a `GET` request which will return the
configured state of all available devices. Further a `PATCH` to
`/devices/dds/<name>` can be used to amend device parameters i.e. frequency.

```shell
curl -X PATCH \
     -H "Content-Type: application/json" \
     -d '{"singleTone":{"frequency":100000000}}'\
     http://localhost:8000/devices/dds/DDS0
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

### How to render docs

You can serve the documentation of the Go packages with

    godoc -http=:6060

and then visiting [http://localhost:6060](http://localhost:6060).

### How to enable SPI pins

Some Beaglebone pins are multi purpose and have to be configured accordingly.

In our case we need to set `P9.17`, `P9.21`, `P9.18` and `P9.22` to SPI mode,
other pins can be left to be set to GPIO mode.

With recent releases (4.x linux kernel) you can easily configure the pins via

    config-pin p9.17 spi_cs
    config-pin p9.21 spi
    config-pin p9.18 spi
    config-pin p9.22 spi_sclk

but if you are a Bilal then you can also ammend the `/boot/uEnv.txt` with

    dtb_overlay=/lib/firmware/BB-SPIDEV0-00A0.dtbo
    disable_uboot_overlay_video=1
    disable_uboot_overlay_audio=1
    disable_uboot_overlay_wireless=1

then the bootloader will enable `SPI0` on startup.

### How to connect to the Beaglebone

The preferable access to the Beaglebone is via SSH. The Beaglebone should
broadcast itself as `beaglebone.local` yet there may be name conflicts if
your network compromises multiple beaglebones.

For development it may be preferable to have internet access from the
Beaglebone. In case your target network does not have internet access (like
the notorius Labornetz) then we recommend to share the wifi connection over
ethernet. For macOS you can enable this feature in the sharing settings, note
that `eduroam` cannot be shared because too enterprise. On Windows we were
not successful in getting this setup to work.

For low-level debugging for example if a kernel update went wrong, the
bootloader was misconfigured or there is no network connection (no DHCP
available or the like) you can still connect to the Beaglebone through a
serial console, howbeit a serial (to usb) adapter will be necessary.

### How to flash the internal MMC

This is necessary if your Beaglebone was delivered with an older image as we
want to use Debian with a 4.x linux kernel.

*We were not able to initiate the flashing procedure with Rev. A Beaglebones!*

To update the internal MMC with the new image you need a micro SD card. You can
find the available images [here](https://rcn-ee.com/rootfs/bb.org/testing/).
We used the `2018-03-25` release with `stretch-lxqt`, however in retrospect
`stretch-console` without the user interface bloat should be better fit.

Post download you need to extract the image and write it to the micro SD card

    xzcat <bb-image>.img.xz | sudo dd of=/dev/disk2

and start the flash process by inserting the SD card into the unpowered
Beaglebone, power on and hold the S2 until the LEDs start to flash.
