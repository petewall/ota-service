const { Before, After, Then, When } = require("cucumber")
const { Builder, By, until } = require("selenium-webdriver")
const { Options } = require("selenium-webdriver/chrome")
const assert = require("assert")

async function findDeviceInTable(driver, mac) {
  let rows = await driver.findElements(By.css("#deviceTable tr"))
  for (let row of rows) {
    let cells = await row.findElements(By.tagName("td"))
    if (cells.length > 0) {
      let rowMAC = await cells[0].getText()
      if (mac == rowMAC) {
        return row
      }
    }
  }

  return null
}

After(async function () {
  if (this.driver) {
    await this.driver.quit();
    this.driver = null
  }
})

When("I view the registry page", async function () {
  if (!this.driver) {
    let builder = new Builder().forBrowser("chrome")
    if (process.env.CI == "true") {
      console.log("running in headless mode")
      let options = new Options()
      options.headless()
      options.addArguments("--disable-gpu")
      options.addArguments("--no-sandboxex")
      builder.setChromeOptions(options)
    }
    this.driver = await builder.build()  
  }

  await this.driver.get(`http://localhost:${this.port}`)
  
  let deviceTableLoading = await this.driver.findElement(By.css("#deviceTable caption"))
  let firmwareTableLoading = await this.driver.findElement(By.css("#firmwareTable caption"))
  await this.driver.wait(until.elementIsNotVisible(deviceTableLoading), 1000);
  await this.driver.wait(until.elementIsNotVisible(firmwareTableLoading), 1000);
})

When("I select {} from the dropdown of firmware for {} on the registry page", async function (type, mac) {
  let row = await findDeviceInTable(this.driver, mac)
  if (!row) {
    assert.fail(`Device ${mac} not found in device list`)
  }

  await row.findElement(By.tagName("select")).sendKeys(type)
})

Then("the device list is empty", async function () {
  let rows = await this.driver.findElements(By.css("#deviceTable tr"))
  assert.equal(rows.length, 1) // includes header row
})

Then("the device list has {} entr{}", async function (count, dummy) {
  let rows = await this.driver.findElements(By.css("#deviceTable tr"))
  assert.equal(rows.length, parseInt(count) + 1) // includes header row
})

Then("the device list has a device with mac {} running {} version {}", async function (mac, type, version) {
  let row = await findDeviceInTable(this.driver, mac)
  if (!row) {
    assert.fail(`Device ${mac} not found in device list`)
  }
 
  let cells = await row.findElements(By.tagName("td"))
  assert.equal(await cells[1].getText(), type)
  assert.equal(await cells[2].getText(), version)
})

Then("the registry page shows that the state of device {} is {}", async function (mac, state) {
  let row = await findDeviceInTable(this.driver, mac)
  if (!row) {
    assert.fail(`Device ${mac} not found in device list`)
  }

  let cells = await row.findElements(By.tagName("td"))
  assert.equal(await cells[4].getText(), state)
})

Then("the firmware list is empty", async function () {
  let rows = await this.driver.findElements(By.css("#firmwareTable tr"))
  assert.equal(rows.length, 1) // includes header row
})

Then("the firmware list has {} entr{}", async function (count, dummy) {
  let rows = await this.driver.findElements(By.css("#firmwareTable tr"))
  assert.equal(rows.length, parseInt(count) + 1) // includes header row
})

Then("the firmware list contains a firmware for {} with a version of {}", async function (type, version) {
  let rows = await this.driver.findElements(By.css("#firmwareTable tr"))
  let found = false
  for (let row of rows) {
    let cells = await row.findElements(By.tagName("td"))
    if (cells.length > 0) {
      let rowType = await cells[0].getText()
      let rowVersion = await cells[1].getText()
      if (type == rowType && version == rowVersion) {
        found = true
      }
    }
  }

  assert(found, `Firmware ${type}:${version} not found in firmware list`)
})
