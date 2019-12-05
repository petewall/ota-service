const util = require("util")
const { Before, After, Given, Then } = require("cucumber")
const assert = require("assert")
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
  assert.equal(this.requestResult.err, null)
  assert.equal(this.requestResult.response.statusCode, status.OK)
})

Then("the request returns no content", function () {
  assert.equal(this.requestResult.err, null)
  assert.equal(this.requestResult.response.statusCode, status.NO_CONTENT)
})

Then("the request returns not found", function () {
  assert.equal(this.requestResult.err, null)
  assert.equal(this.requestResult.response.statusCode, status.NOT_FOUND)
})

Then("I receive the value {}", function (value) {
  assert.equal(this.requestResult.body, value)
})

Then("I receive an empty list", function () {
  assert.equal(this.requestResult.body, "[]")
})

Then("I receive a list with {} entr{}", function (size, dummy) {
  try {
    this.result = JSON.parse(this.requestResult.body)
  } catch (e) {
    assert.fail(`request body could not be parsed: ${this.requestResult.body}`)
  }
  assert.equal(typeof this.result, "object")
  assert.equal(this.result.length, size)
})

Then("I receive a hash", function () {
  try {
    this.result = JSON.parse(this.requestResult.body)
  } catch (e) {
    assert.fail(`request body could not be parsed: ${this.requestResult.body}`)
  }
  assert.equal(typeof this.result, "object")
})

Then("the result has a{} {} of {}", function (dummy, key, value) {
  assert.notEqual(typeof(this.result[key]), "undefined")
  assert.equal(this.result[key], value)
})

Then("the result has no {}", function (key) {
  assert.equal(Boolean(this.result[key]), false)
})
