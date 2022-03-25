#!/bin/bash

qemu-system-x86_64 \
-m 8G \
-smp 6 \
-drive file=mydisk.qcow2,if=virtio \
-vga virtio \
-display default,show-cursor=on \
-usb \
-device usb-tablet \
-cpu host \
-machine type=q35,accel=hvf

