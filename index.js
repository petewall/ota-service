#!/usr/local/bin/node

const express = require("express")
const app = express()
const glob = require("glob")
const path = require("path")

let devices = {}
let binaries = {}

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

function readBinariesFromDisk() {
    glob(path.join(process.env.DATA_DIR, "*", "/*"), function (err, files) {
        binaries = {}
        for (let file of files) {
            let parts = file.split(path.sep)
            let device = parts[parts.length - 2]
            let binary = parts[parts.length - 1]
            binaries[device] = [{
                version: binary.split("-")[1],
                filename: binary,
            }]
        }
    })
}
readBinariesFromDisk()

app.get("/devices", (req, res) => {
    res.json(devices)
})

app.get("/binary", (req, res) => {
    const device = req.header("HTTP_X_ESP8266_VERSION").split("-")[0]
    const version = req.header("HTTP_X_ESP8266_VERSION").split("-")[1]
    devices[device] = { device, version }

    // console.log(`device: ${device}`)
    // console.log(`version: ${version}`)
    // console.log("binaries: ")
    // console.log(binaries)
    if (binaries[device] && binaries[device][0].version > version) {
        let binaryPath = path.join(process.env.DATA_DIR, device, binaries[device][0].filename)
        // console.log(`sending ${binaryPath}`)
        res.sendFile(binaryPath)
        devices[device] = { device, version: binaries[device][0].version }
    } else {
        res.sendStatus(304)
    }
})

app.get("/binaries", (req, res) => {
    res.json(binaries)
})

app.listen(process.env.PORT, () => console.log(`OTA Service listening on port ${process.env.PORT}!`))
