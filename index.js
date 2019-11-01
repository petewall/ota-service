#!/usr/local/bin/node

const express = require("express")
const app = express()
const glob = require("glob")
const path = require("path")

app.get('/binaries', (req, res) => {
    glob(path.join(process.env.DATA_DIR, "*", "/*.bin"), function (err, files) {
        let results = {}
        for (let file of files) {
            let parts = file.split(path.sep)
            let device = parts[parts.length - 2]
            let binary = parts[parts.length - 1]
            results[device] = [ binary ]
        }
        res.json(results)
    })
})

// app.get("/updates", (req, res) => {
//     // [HTTP_USER_AGENT] => ESP8266-http-Update
//     // [HTTP_X_ESP8266_STA_MAC] => 18:FE:AA:AA:AA:AA
//     // [HTTP_X_ESP8266_AP_MAC] => 1A:FE:AA:AA:AA:AA
//     // [HTTP_X_ESP8266_FREE_SPACE] => 671744
//     // [HTTP_X_ESP8266_SKETCH_SIZE] => 373940
//     // [HTTP_X_ESP8266_SKETCH_MD5] => a56f8ef78a0bebd812f62067daf1408a
//     // [HTTP_X_ESP8266_CHIP_SIZE] => 4194304
//     // [HTTP_X_ESP8266_SDK_VERSION] => 1.3.0
//     // [HTTP_X_ESP8266_VERSION] => DOOR-7-g14f53a19

//     // const currentVersion = req.header("HTTP_X_ESP8266_SDK_VERSION")
//     res.sendStatus(304)
// })

if (!process.env.PORT) {
    console.error("No port defined.")
    process.exit(1)
}
if (!process.env.DATA_DIR) {
    console.error("No data path defined.")
    process.exit(1)
}
app.listen(process.env.PORT, () => console.log(`OTA Service listening on port ${process.env.PORT}!`))


// EnvironmentVariable for:
// binary dir

// Endpoints for:
// get List types
// get List versions for type
// get List known devices and their versions
// put add binary


