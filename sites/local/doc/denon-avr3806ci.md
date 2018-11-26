---
title: Denon AVR 3806 CI
description: >
 A pretty high end surround sound reciever that makes that
 music in my man cave.
date: 2018-11-21
categories: [ audio, video, automation ]
tags: denon, serial, control
refs:
  - TODO pointer to all denon manuals (place these in CDN)
  - TODO denon website
params:
  mac: 00:05:cd:13:d8:ca
  ip: 10.24.2.11
connections:
  hdmi-out:	 vizio
  hdmi-dvd:	 roku1
  hdmi-cab1: rpi1
  hdmi-cab2: rpi2
  eth: net1024
  ser0: bbone.ttyUSB0
---

I was able to get the Denon Firmware upgraded and responding to both
Ethernet and Serial command prompts.  However, I seem to be sending
data to the telnet or serial port.

Need to look closer to get the controls working correctly.
