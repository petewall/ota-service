# OTA Service

## Upgrade algorithm

1. request comes in
2. Get mac, current type, and firmware
3. Get the device record from device service
4. is current type and firmware (reported by the request) different from the one from the device service?
  If yes, sent an update
5. is firmware type set?
  if yes, is firmware version set?
    if yes, 




if assigned type and version, but it is not the same, send update
if assigned type, but not version, get latest from firmware service, if it's newer, send update
if no assigned type, do nothing
