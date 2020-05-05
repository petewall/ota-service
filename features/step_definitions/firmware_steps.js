const { Given, Then, When } = require("cucumber")
const { expect } = require('chai')
const eventually = require("./eventually.js")
const fs = require("fs").promises
const util = require("util")
const glob = util.promisify(require("glob"))
const mkdirp = require("mkdirp")
const request = require("request")
const path = require("path")

Given("an empty firmware directory", function () {})

async function addBinary(tmpDir, type, version) {
  let directory = path.join(tmpDir, "firmware", type, version)
  let firmwarePath = path.join(directory, `${type}-${version}.bin`)
  await mkdirp(directory)
  await fs.writeFile(firmwarePath, `data-for-${type}-${version}`)
}

Given("a firmware binary with type {} and version {}", async function (type, version) {
  await addBinary(this.tmpDir, type, version)
})

When("I ask for the list of firmware binaries", function (done) {
  request.get(`http://localhost:${this.port}/api/firmware`, (err, response, body) => {
    this.requestResult = { err, response, body }
    done()
  })
})

When("a firmware binary for {} with a version of {} is added", async function (type, version) {
  await addBinary(this.tmpDir, type, version)
})

When("I send a binary file for {} with a version of {}", function (type, version, done) {
  request.put(`http://localhost:${this.port}/api/firmware/${type}/${version}`, {
    body: "my-firmware-data",
    headers: {
      "Content-type": "application/octet-stream"
    }
  }, (err, response, body) => {
    this.requestResult = { err, response, body }
    done()
  })
})

When("I send a delete request for {} with a version of {}", function (type, version, done) {
  request.delete(`http://localhost:${this.port}/api/firmware/${type}/${version}/${type}-${version}.bin`, (err, response, body) => {
    this.requestResult = { err, response, body }
    done()
  })
})

Then("it contains a firmware for {} with a version of {}", function (type, version) {
  let found = false
  for (let entry of this.result) {
    if (entry.type == type && entry.version == version) {
      found = true;
    }
  }

  expect(found, `Firmware for ${type}:${version} not found in result list`).to.be.true
})

Then("the service detects {} binar{}", {timeout: 60 * 1000}, async function (count, dummy) {
  await eventually(() => this.serviceStdout.indexOf(`[Firmware] Firmware loaded: ${count} binaries`) >= 0, 50 * 1000)
})

Then("the service sends the firmware binary for {} with version {}", function (type, version) {
  expect(this.requestResult.body).to.equal(`data-for-${type}-${version}`)
})

Then("the a binary file for {} with a version of {} exists in the firmware directory", async function (type, version) {
  let files = await glob(path.join(this.tmpDir, "firmware", type, version, "*.bin"))
  expect(files).to.have.length(1)
})
