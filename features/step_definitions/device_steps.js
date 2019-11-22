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

When("I assign a firmware type of {} to {}", function (type, mac, done) {
  request.post(`http://localhost:${this.port}/api/assign?firmware=${type}&mac=${mac}`, (err, response, body) => {
    this.requestResult = { err, response, body }
    done()
  })
})

Then("the firmware {} is assigned to {}", async function (type, mac) {
  await eventually(() => this.serviceStdout.indexOf(`[Device] Setting ${mac} to firmware ${type}`) >= 0)
})

Then("it contains a device with mac {} running {} version {}", function (mac, type, version) {
  for (let device of this.result) {
    if (device.mac == mac) {
      assert.equal(device.firmwareType, type)
      assert.equal(device.firmwareVersion, version)
      assert.equal(device.assignedFirmware, type)
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