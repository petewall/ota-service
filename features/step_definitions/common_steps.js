const assert = require("assert")
const { Before, After, Given, When, Then } = require("cucumber")
const { spawn } = require("child_process")
const fs = require("fs").promises
const path = require("path")
const util = require("util")
const rimraf = util.promisify(require("rimraf"))
const getPort = require("get-port")
const debug = require("debug")

Before(async function () {
  debug.enable("otaService:*")
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
      if (!started && data.indexOf("OTA Service listening on port") >= 0) {
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
  assert.equal(this.requestResult.response.statusCode, 200)
})
