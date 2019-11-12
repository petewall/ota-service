#!/usr/local/bin/node

const express = require("express")
const app = express()
const path = require("path")
const semver = require("semver")
const status  = require("http-status")

const Devices = require("./devices.js")
const Firmware = require("./firmware.js")

if (!process.env.PORT) {
    console.error("No port defined.")
    process.exit(1)
}

if (!process.env.DATA_DIR) {
    console.error("No data path defined.")
    process.exit(1)
}
if (!path.isAbsolute(process.env.DATA_DIR)) {
    console.error("DATA_DIR must be an absolute path.")
    process.exit(1)
}

let devices = new Devices()
let firmwareLibrary = new Firmware(process.env.DATA_DIR)

app.get("/api/firmware", (req, res) => {
    res.json(firmwareLibrary.getAll())
})

app.get("/api/devices", (req, res) => {
    res.json(devices.getAll())
})

app.get("/api/update/", (req, res) => {
    console.log(req.headers)
    let mac = req.get("x-esp8266-sta-mac")
    let currentType = req.query.firmware
    let currentVersion = req.query.version
    console.log(`New request from ${mac}: type: ${currentType} version: ${currentVersion}`)

    let device = devices.get(mac)
    if (!device) {
        devices.registerDevice(mac, currentType, currentVersion)
        device = devices.get(mac)
    }

    let latestFirmware = firmwareLibrary.getLatestForType(device.firmwareType)
    if (!latestFirmware) {
        console.log(`No firmware found for ${device.firmwareType}`)
        return res.sendStatus(status.NOT_MODIFIED)
    }

    if (semver.gt(latestFirmware.version, currentVersion)) {
        console.log("Sending new firmware: ", latestFirmware)
        return res.sendFile(latestFirmware.file)
    }

    console.log(`No firmware to send`)
    res.sendStatus(status.NOT_MODIFIED)
})

app.use(express.static("public"))

app.listen(process.env.PORT, () => console.log(`OTA Service listening on port ${process.env.PORT}!`))
