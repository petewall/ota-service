#!/usr/local/bin/node

const express = require("express")
const app = express()
const bodyParser = require("body-parser")
const morgan = require("morgan")
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

app.use(morgan("combined"))

app.get("/api/firmware", (req, res) => {
  res.json(firmwareLibrary.getAll())
})

app.put("/api/firmware/:type/:version([0-9a-zA-Z-._]+)", bodyParser.raw({ limit: "5mb" }), async (req, res) => {
  await firmwareLibrary.addBinary(req.params.type, req.params.version, req.body)
  res.sendStatus(status.OK)
})

app.delete("/api/firmware/:type/:version([0-9a-zA-Z-._]+)/:filename([0-9a-zA-Z-._]+.bin)", async (req, res) => {
  try {
    await firmwareLibrary.deleteBinary(req.params.type, req.params.version, req.params.filename)
    res.sendStatus(status.OK)
  } catch (err) {
    if (err.code == "ENOENT") {
      res.sendStatus(status.NOT_FOUND)
    } else {
      res.sendStatus(status.INTERNAL_SERVER_ERROR)
    }
  }
})

app.get("/api/devices", (req, res) => {
    res.json(devices.getAll())
})

app.get("/api/update", (req, res) => {
    let mac = req.get("x-esp8266-sta-mac")
    let currentType = req.query.firmware
    let currentVersion = req.query.version
    console.log(`New request from ${mac}: type: ${currentType} version: ${currentVersion}`)

    let device = devices.updateDevice(mac, currentType, currentVersion)
    let latestFirmware = firmwareLibrary.getLatestForType(device.assignedFirmware)
    if (!latestFirmware) {
        console.log(`No firmware found for ${device.firmwareType}`)
        devices.setState(mac, "up to date")
        return res.sendStatus(status.NOT_MODIFIED)
    }

    if (device.firmwareType != device.assignedFirmware) {
        console.log("Firmware type changed.  Sending new firmware: ", latestFirmware)
        devices.setState(mac, "updating")
        return res.sendFile(latestFirmware.file)
    }

    if (semver.gt(latestFirmware.version, currentVersion)) {
        console.log("Newer version available.  Sending new firmware: ", latestFirmware)
        devices.setState(mac, "updating")
        return res.sendFile(latestFirmware.file)
    }

    console.log("Device is up to date")
    devices.setState(mac, "up to date")
    res.sendStatus(status.NOT_MODIFIED)
})

app.post("/api/assign", (req, res) => {
    let type = req.query.firmware
    let mac = req.query.mac
    devices.assignFirmware(mac, type)
    res.sendStatus(status.OK)
})

app.use(express.static("public"))

app.set("view engine", "ejs");
app.get("/", (req, res) => {
  res.render("index", {
    devices: devices.getAll(true),
    allFirmware: firmwareLibrary.getAll(true),
    firmwareTypes: firmwareLibrary.getAllTypes()
  });
});

app.listen(process.env.PORT, () => console.log(`OTA Service listening on port ${process.env.PORT}!`))
