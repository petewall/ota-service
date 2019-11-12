const chokidar = require("chokidar")
const glob = require("glob")
const semver = require("semver")
const path = require("path")

class Firmware {
    constructor(dataPath) {
        this.glob = require("glob")
        this.firmwareGlob = path.join(dataPath, "firmware", "*", "*", "*.bin")
        this.loadFromPath(dataPath)
        let watcher = chokidar.watch(this.firmwareGlob)

        watcher.on("add", (file) => {
            console.log(`[Firmware] Binary file added: ${file}`)
            this.loadFromPath()
        })
        watcher.on("unlink", (file) => {
            console.log(`[Firmware] Binary file removed: ${file}`)
            this.loadFromPath()
        })
    }

    loadFromPath() {
        let newFirmware = []
        console.log("[Firmware] Loading firmware from the data directory...")
        glob(this.firmwareGlob, (err, files) => {
            if (err) {
                console.error("[Firmware] failed to find firmware files: ", err)
                process.exit(1)
            }

            for (let file of files) {
                let parts = file.split(path.sep)
                let version = parts[parts.length - 2]
                if (!semver.valid(version)) {
                    console.error(`[Firmware] Invalid version for file: ${file}`)
                } else {
                    console.log(`[Firmware]    ${file}`)
                    newFirmware.push({
                        file,
                        filename: parts[parts.length - 1],
                        type: parts[parts.length - 3],
                        version
                    })
                }
            }
            this.firmware = newFirmware
            console.log(`[Firmware] Firmware loaded: ${this.firmware.length} binaries`)
        })
    }

    getAll() {
        return this.firmware
    }

    // getFirmwareTypes() {
    //     let set = new Set(this.firmware.map(firmware => firmware.type))
    //     return Array.from(set.values())
    // }

    getAllForType(firmwareType) {
        return this.firmware.filter(firmware => firmware.type == firmwareType)
    }

    getLatestForType(firmwareType) {
        return this.getAllForType(firmwareType).reduce((latestFirmware, firmware) => {
            if (!latestFirmware) {
                return firmware;
            }
            if (semver.gt(firmware.version, latestFirmware.version)) {
                return firmware;
            }
            return latestFirmware;
        }, null)
    }
}

module.exports = Firmware