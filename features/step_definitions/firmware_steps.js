const { Given, Then, When } = require("cucumber")
const assert = require("assert")
const fs = require("fs").promises
const request = require("request")
const path = require("path")

function eventually(check) {
  return new Promise((resolve, reject) => {
    let count = 0;
    let checkerId = setInterval(() => {
      if (check()) {
        resolve()
        clearInterval(checkerId)
      } else {
        count += 1
        if (count >= 10) {
          reject()
          clearInterval(checkerId)
        }
      }
    }, 100)
  })
}

Given("an empty firmware directory", function () {})

async function addBinary(tmpDir, type, version) {
  let firmwareTypeDir = path.join(tmpDir, "firmware", type)
  let firmwareVersionDir = path.join(tmpDir, "firmware", type, version)
  let firmwareBinaryPath = path.join(tmpDir, "firmware", type, version, `${type}-${version}.bin`)
  try {
    await fs.stat(firmwareTypeDir)
  } catch (e) {
    assert.equal(e.code, "ENOENT")
    await fs.mkdir(firmwareTypeDir)
  }
  await fs.mkdir(firmwareVersionDir)
  await fs.writeFile(firmwareBinaryPath, `data-for-${type}-${version}`)
}

Given("there is a firmware binary for {} with a version of {}", async function (type, version) {
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

Then("I receive an empty list", function () {
  assert.equal(this.requestResult.body, "[]")
})

Then("I receive a list with {} entr{}", function (size, dummy) {
  this.result = JSON.parse(this.requestResult.body)
  assert.equal(this.result.length, size)
})

Then("it contains a firmware for {} with a version of {}", function (type, version) {
  let found = false
  for (let entry of this.result) {
    if (entry.type == type && entry.version == version) {
      found = true;
    }
  }

  assert(found, `Firmware for ${type}:${version} not found in result list`)
})

Then("the service detects {} binar{}", async function (count, dummy) {
  await eventually(() => this.serviceStdout.indexOf(`[Firmware] Firmware loaded: ${count} binaries`) >= 0)
})

Then("the service sends the firmware binary for {} with version {}", function (type, version) {
  assert.equal(this.requestResult.body, `data-for-${type}-${version}`)
})
