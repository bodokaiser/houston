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

## Appendix

### Flash Image

Update your Beaglebone image by writing a flasher image to the micro SD card

```shell
  curl -O https://rcn-ee.com/rootfs/bb.org/testing/2018-03-25/stretch-lxqt/BBB-blank-debian-9.4-lxqt-armhf-2018-03-25-4gb.img.xz
  xzcat BBB-blank-debian-9.4-lxqt-armhf-2018-03-25-4gb.img.xz | sudo dd of=/dev/disk2
```

and start the flash process by inserting the SD card into the
unpowered Beaglebone, press the S2 button and power on.
