# Houston

Houston connects your lab equipment.

## Usage

### CLI

The command line interface is defined in `cmd/dds/main.go` and offers DDS
control without any graphical or network interface.

See `dds --help` for details.

### HTTP

The http interface is defined in `cmd/http/main.go` and offers DDS
control over RESTful HTTP interface.

See `http --help` for details.

For a web application as frontend you can use
[houston-app](http://github.com/bodokaiser/houston-app).

### HTTP Dev

The `httpdev` command is intended for development use only. It uses driver
mockups to check if the HTTP interface uses the driver interface as expected
without the need to have the hardware present.

## Installation

We assume you have a working Go environment set up on your workstation and

    go get -u github.com/gobuffalo/packr/...

executes successful. You can clone the source code via

    export $REPO=github.com/bodokaiser/Houston

    git clone https://$REPO $GOPATH/src/$REPO

and ensure dependencies by calling

    dep ensure

inside the project directory. Install [Go Dep](https://golang.github.io/dep/)
if required.

From here you can easily cross compile binaries for ARM by running `make` and
copy files to the target device:

    scp -r bin debian@beaglebone.local:~/

## Development

Code documentation is available [online](https://godoc.org/github.com/bodokaiser/houston)
or can be hosted locally with `godoc -http=:6060`.

## Appendix

### How to enable SPI pins

Some Beaglebone pins are multi purpose and have to be configured accordingly.

In our case we need to set `P9.17`, `P9.21`, `P9.18` and `P9.22` to SPI mode,
other pins can be left to be set to GPIO mode.

With recent releases (4.x linux kernel) you can easily configure the pins via

    config-pin p9.17 spi_cs
    config-pin p9.21 spi
    config-pin p9.18 spi
    config-pin p9.22 spi_sclk

or permanent if you ammend `/boot/uEnv.txt` to

    dtb_overlay=/lib/firmware/BB-SPIDEV0-00A0.dtbo
    disable_uboot_overlay_video=1
    disable_uboot_overlay_audio=1
    disable_uboot_overlay_wireless=1

in which the bootloader will enable `SPI0` on startup.

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
