# Beagle

## Setup

Write latest beaglebone image to micro SD card

  curl -O https://rcn-ee.com/rootfs/bb.org/testing/2018-03-25/stretch-lxqt/BBB-blank-debian-9.4-lxqt-armhf-2018-03-25-4gb.img.xz
  xzcat BBB-blank-debian-9.4-lxqt-armhf-2018-03-25-4gb.img.xz | sudo dd of=/dev/disk2

and flash the beaglebones internal memory by inserting the SD card into the
unpowered beaglebone, press the S2 button and power on.
