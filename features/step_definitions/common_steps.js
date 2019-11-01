const assert = require("assert")
const { Before, After, Given, When, Then } = require("cucumber")
const { spawn } = require("child_process")
const fs = require("fs").promises
const path = require("path")
const util = require("util")
const request = require("request")
const rimraf = util.promisify(require("rimraf"))
const getPort = require("get-port")
const debug = require("debug")

Before(async function () {
  debug.enable("*")
  this.tmpDir = await fs.mkdtemp("tmp-features-")
  this.binDir = path.join(this.tmpDir, "binaries")
  await fs.mkdir(this.binDir)
})

After(async function () {
  this.otaService.kill()
  await rimraf(this.tmpDir)
})

Given("an empty binary directory", function () {})

Given("a binary directory with one binary", async function () {
  await fs.mkdir(path.join(this.binDir, "WEMOS_OFFICE"))
  await fs.writeFile(path.join(this.binDir, "WEMOS_OFFICE", "WEMOS_OFFICE_1.2.3.bin"))
})

Given("a binary directory with binaries for multiple devices", async function () {
  await fs.mkdir(path.join(this.binDir, "WEMOS_OFFICE"))
  await fs.writeFile(path.join(this.binDir, "WEMOS_OFFICE", "WEMOS_OFFICE_1.2.3.bin"))
  await fs.mkdir(path.join(this.binDir, "WEMOS_GARAGE"))
  await fs.writeFile(path.join(this.binDir, "WEMOS_GARAGE", "WEMOS_GARAGE_1.2.3.bin"))
})

Given("the OTA Service is running", function (done) {
  getPort().then((port) => {
    this.port = port

    let env = process.env
    env.PORT = port
    env.DATA_DIR = this.binDir
  
    let started = false
    let stdout = debug("otaService:stdout")
    let stderr = debug("otaService:stderr")
    this.otaService = spawn("node", ["index.js"], { env })
    this.otaService.on("error", (err) => {
      assert.fail("The OTA service failed to start: ", err)
    })
    this.otaService.stdout.on("data", (data) => {
      stdout(data.toString())
      if (!started && data.indexOf("OTA Service listening on port") >= 0) {
        started = true
        done()
      }
    })
    this.otaService.stderr.on("data", (data) => { stderr(data.toString()) })
  })
})

When("I ask for the list of binaries", function (done) {
  request.get(`http://localhost:${this.port}/binaries`, (err, response, body) => {
    this.requestResult = { err, response, body }
    done()
  })
})

Then("the request is successful", function () {
  assert.equal(this.requestResult.err, null)
  assert.equal(this.requestResult.response.statusCode, 200)
})

Then("I receive an empty hash", function () {
  assert.equal(this.requestResult.body, "{}")
})

Then("I receive a hash with a binary in a single device", function () {
  assert.deepEqual(JSON.parse(this.requestResult.body), {
    WEMOS_OFFICE: [
      "WEMOS_OFFICE_1.2.3.bin"
    ]
  })
})

Then("I receive a hash with multiple devices", function () {
  assert.deepEqual(JSON.parse(this.requestResult.body), {
    WEMOS_GARAGE: [
      "WEMOS_GARAGE_1.2.3.bin"
    ],
    WEMOS_OFFICE: [
      "WEMOS_OFFICE_1.2.3.bin"
    ]
  })
})
