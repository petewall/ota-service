const { Given, Then, When } = require("cucumber")
const assert = require("assert")
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

Then("I receive an empty hash", function () {
  assert.equal(this.requestResult.body, "{}")
})

Then("I receive a hash with {} entr{}", function (size, dummy) {
  this.result = JSON.parse(this.requestResult.body)
  assert.equal(Object.keys(this.result).length, size)
})

Then("it contains a device with mac {} running {} version {}", function (mac, type, version) {
  assert.deepEqual(this.result[mac], {
    mac,
    firmwareType: type,
    firmwareVersion: version
  })
})

Then("the service responds with no update", function () {
  assert.equal(this.requestResult.err, null)
  assert.equal(this.requestResult.response.statusCode, status.NOT_MODIFIED)
})