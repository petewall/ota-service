const util = require("util")
const { Before, After, Given, Then } = require("cucumber")
const { assert, expect } = require('chai')
const debug = require("debug")
const fs = require("fs").promises
const getPort = require("get-port")
const path = require("path")
const rimraf = util.promisify(require("rimraf"))
const spawn = require("child_process").spawn
const status  = require("http-status")

Before(async function () {
  this.tmpDir = await fs.mkdtemp("tmp-features-")
  await fs.mkdir(path.join(this.tmpDir, "firmware"))
})

After(async function () {
  if (this.otaService) {
    this.otaService.kill()
  }
  await rimraf(this.tmpDir)
})

Given("the OTA service is running", function (done) {
  getPort().then((port) => {
    this.port = port

    let env = process.env
    env.PORT = port
    env.DATA_DIR = path.join(process.cwd(), this.tmpDir)
  
    let started = false
    let stdout = debug("otaService:stdout")
    let stderr = debug("otaService:stderr")
    this.otaService = spawn("node", ["index.js"], { env })
    this.otaService.on("error", (err) => {
      assert.fail("The OTA service failed to start: ", err)
    })
    this.otaService.stdout.on("data", (data) => {
      stdout(data.toString())
      if (!started && data.indexOf("[Firmware] Firmware loaded") >= 0) {
        started = true
        done()
      }
      this.serviceStdout += data.toString()
    })
    this.otaService.stderr.on("data", (data) => { stderr(data.toString()) })
  })
})

Then("the request is successful", function () {
  expect(this.requestResult.err).to.be.null
  expect(this.requestResult.response.statusCode).to.equal(status.OK)
})

Then("the request returns no content", function () {
  expect(this.requestResult.err).to.be.null
  expect(this.requestResult.response.statusCode).to.equal(status.NO_CONTENT)
})

Then("the request returns not found", function () {
  expect(this.requestResult.err).to.be.null
  expect(this.requestResult.response.statusCode).to.equal(status.NOT_FOUND)
})

Then("I receive the value {}", function (value) {
  expect(this.requestResult.body).to.equal(value)
})

Then("I receive an empty list", function () {
  expect(this.requestResult.body).to.equal("[]")
})

Then("I receive a list with {} entr{}", function (size, dummy) {
  try {
    this.result = JSON.parse(this.requestResult.body)
  } catch (e) {
    assert.fail(`request body could not be parsed: ${this.requestResult.body}`)
  }
  expect(this.result).to.be.an("array")
  expect(this.result).to.have.lengthOf(size);
})

Then("I receive a hash", function () {
  try {
    this.result = JSON.parse(this.requestResult.body)
  } catch (e) {
    assert.fail(`request body could not be parsed: ${this.requestResult.body}`)
  }
  expect(this.result).to.be.an("object")
})

Then("the result has a{} {} of {}", function (dummy, key, value) {
  expect(this.result).to.be.an("object")
  expect(this.result).to.include.all.keys(key)
  expect(this.result[key]).to.equal(value)
})

Then("the result has no {}", function (key) {
  expect(this.result).to.not.include(key)
})
