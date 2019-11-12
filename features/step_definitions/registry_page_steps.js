const { When } = require("cucumber")
const {Builder, By, Key, until} = require('selenium-webdriver');

When("I view the registry page", async function () {
    this.driver = await new Builder().forBrowser("chrome").build();
    await this.driver.get(`http://localhost:${this.port}`);
    
})


