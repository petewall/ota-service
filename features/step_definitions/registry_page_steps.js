const { After, Then, When } = require("cucumber")
const { Builder, By, until } = require("selenium-webdriver")
const { Options } = require("selenium-webdriver/chrome")
const { assert, expect } = require('chai')
const debug = require("debug")
const semver = require("semver")

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

async function findFirmwareInTable(driver, type, version) {
  let rows = await driver.findElements(By.css("#firmwareTable tr"))
  for (let row of rows) {
    let cells = await row.findElements(By.tagName("td"))
    if (cells.length > 0) {
      let rowType = await cells[0].getText()
      let rowVersion = await cells[1].getText()
      if (type == rowType && version == rowVersion) {
        return row
      }
    }
  }

  return null
}

After(async function () {
  if (this.driver) {
    debug("browser")(await this.driver.manage().logs().get("browser"))
    await this.driver.quit();
    this.driver = null
  }
})

When("I view the registry page", {timeout: 60 * 1000}, async function () {
  if (!this.driver) {
    let builder = new Builder().forBrowser("chrome")
    if (process.env.CI == "true") {
      let options = new Options()
      options.headless()
      options.addArguments("--disable-dev-shm-usage")
      options.addArguments("--disable-extensions")
      options.addArguments("--disable-gpu")
      options.addArguments("--no-sandbox")
      builder.setChromeOptions(options)
    }
    this.driver = await builder.build()  
  }

  await this.driver.get(`http://localhost:${this.port}`)
})

When("I select {} from the dropdown of firmware for {} on the registry page", async function (type, mac) {
  let row = await findDeviceInTable(this.driver, mac)
  expect(row, `Device ${mac} not found in device list`).to.not.be.undefined

  await row.findElement(By.tagName("select")).sendKeys(type)
})

When("I set the firmware type for {} to unassigned", async function (mac) {
  let row = await findDeviceInTable(this.driver, mac)
  expect(row, `Device ${mac} not found in device list`).to.not.be.undefined

  await row.findElement(By.tagName("select")).sendKeys("Unassigned")
})

When("I click the delete button for {} with a version of {}", async function (type, version) {
  let row = await findFirmwareInTable(this.driver, type, version)
  expect(row, `Firmware ${type} ${version} not found in firmware list`).to.not.be.undefined

  await row.findElement(By.className("deleteFirmware")).click()
  await this.driver.wait(until.alertIsPresent())

  let alert = await this.driver.switchTo().alert()
  let alertText = await alert.getText()
  
  expect(alertText).to.equal(`Are you sure you want to delete the firmare ${type} ${version}?\n\nThis cannot be undone.`)
  await alert.accept()
})

Then("the device list is empty", async function () {
  let rows = await this.driver.findElements(By.css("#deviceTable tr"))
  expect(rows).to.have.length(1) // includes header row
})

Then("the device list has {} entr{}", async function (count, dummy) {
  let rows = await this.driver.findElements(By.css("#deviceTable tr"))
  expect(rows).to.have.length(parseInt(count) + 1) // includes header row
})

Then("the device list has a device with mac {} running {} version {}", async function (mac, type, version) {
  let row = await findDeviceInTable(this.driver, mac)
  if (!row) {
    assert.fail(`Device ${mac} not found in device list`)
  }
 
  let firmwareTypeCell = await row.findElement(By.className("firmwareType"))
  let firmwareVersionCell = await row.findElement(By.className("firmwareVersion"))
  expect(await firmwareTypeCell.getText()).to.equal(type)
  expect(await firmwareVersionCell.getText()).to.equal(version)
})

Then("the registry page shows that the state of device {} is {}", async function (mac, state) {
  let row = await findDeviceInTable(this.driver, mac)
  expect(row, `Device ${mac} not found in device list`).to.not.be.undefined

  let cell = await row.findElement(By.className("state"))
  expect(await cell.getText()).to.equal(state)
})

Then("the firmware list is empty", async function () {
  let rows = await this.driver.findElements(By.css("#firmwareTable tr"))
  expect(rows).to.have.length(1) // includes header row
})

Then("the firmware list has {} entr{}", async function (count, dummy) {
  let rows = await this.driver.findElements(By.css("#firmwareTable tr"))
  expect(rows).to.have.length(parseInt(count) + 1) // includes header row
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

  expect(found, `Firmware ${type}:${version} not found in firmware list`).to.be.true
})

Then("the device list is sorted by mac", async function () {
  let rows = await this.driver.findElements(By.css("#deviceTable tr"))
  let lastMAC = ""

  for (let row of rows) {
    let cells = await row.findElements(By.tagName("td"))
    if (cells.length > 0) {
      let mac = await cells[0].getText()
      expect(lastMAC < mac, `List not properly sorted: ${lastMAC} < ${mac}`).to.be.true
      lastMAC = mac
    }
  }
})

Then("the firmware list is sorted by firmware then by version", async function () {
  let rows = await this.driver.findElements(By.css("#firmwareTable tr"))
  let lastType = ""
  let lastVersion = ""

  for (let row of rows) {
    let cells = await row.findElements(By.tagName("td"))
    if (cells.length > 0) {
      let type = await cells[0].getText()
      let version = await cells[1].getText()
      expect(lastType <= type, `List not properly sorted: ${lastType} <= ${type}`).to.be.true
      if (lastType == type) {
        expect(semver.gt(lastVersion, version), `List not properly sorted: ${lastVersion} > ${version}`).to.be.true
      }
      lastType = type
      lastVersion = version
    }
  }
})
