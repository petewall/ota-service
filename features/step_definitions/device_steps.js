const { Given, Then, When } = require("cucumber")
const assert = require("assert")
const eventually = require("./eventually.js")
const request = require("request")
const status  = require("http-status")

Given("an update request comes from {} running {} version {}", function (mac, type, version, done) {
  request.get(`http://localhost:${this.port}/api/update?firmware=${type}&version=${version}`, {
    headers: {
      "x-esp8266-sta-mac": mac
    }
  }, (err, response, body) => {
    this.requestResult = { err, response, body }
    done()
  })
})

When("I ask for the list of devices", function (done) {
  request.get(`http://localhost:${this.port}/api/devices`, (err, response, body) => {
    this.requestResult = { err, response, body }
    done()
  })
})

When("I ask for the device properties for {}", function (mac, done) {
  request.get(`http://localhost:${this.port}/api/device/${mac}`, (err, response, body) => {
    this.requestResult = { err, response, body }
    done()
  })
})

Given("there is a device {} with an assigned firmware type {}", function (mac, type, done) {
  request.patch(`http://localhost:${this.port}/api/device/${mac}?firmware=${type}`, (err, response, body) => {
    this.requestResult = { err, response, body }
    done()
  })
})

When("I assign a device id of {} to {}", function (id, mac, done) {
  request.patch(`http://localhost:${this.port}/api/device/${mac}?id=${id}`, (err, response, body) => {
    this.requestResult = { err, response, body }
    done()
  })
})

When("I assign a firmware type of {} to {}", function (type, mac, done) {
  request.patch(`http://localhost:${this.port}/api/device/${mac}?firmware=${type}`, (err, response, body) => {
    this.requestResult = { err, response, body }
    done()
  })
})

Then("the server updates {} with device id of {}", async function (mac, type) {
  await eventually(() => this.serviceStdout.indexOf(`[Device] Setting ${mac} to device id ${type}`) >= 0)
})

Then("the firmware {} is assigned to {}", async function (type, mac) {
  await eventually(() => this.serviceStdout.indexOf(`[Device] Setting ${mac} to firmware ${type}`) >= 0)
})

Then("it contains a device with mac {} running {} version {}", function (mac, type, version) {
  for (let device of this.result) {
    if (device.mac == mac) {
      assert.equal(device.firmwareType, type)
      assert.equal(device.firmwareVersion, version)
      assert(Date.now() - new Date(device.lastUpdated).getTime() < 1000)
      return
    }
  }
  assert.fail("Device not found")
})

Then("the service responds with no update", function () {
  assert.equal(this.requestResult.err, null)
  assert.equal(this.requestResult.response.statusCode, status.NOT_MODIFIED)
})