#!/bin/bash

qemu-system-x86_64 \
-m 8G \
-smp 6 \
-cdrom /Volumes/Samsung_T5/iso/Fedora-Workstation-Live-x86_64-35-1.2.iso \
-drive file=mydisk.qcow2,if=virtio \
-vga virtio \
-display default,show-cursor=on \
-usb \
-device usb-tablet \
-cpu host \
-machine type=q35,accel=hvf \

