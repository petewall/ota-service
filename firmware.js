const glob = require("glob")
const semver = require("semver")
const path = require("path")
const watch = require("watch")

class Firmware {
    constructor(dataPath) {
        this.glob = require("glob")
        this.loadFromPath(dataPath)
        watch.watchTree(path.join(dataPath, "firmware"), (a) => {
            console.log("[Firmware] Detected firmware changes")
            console.log(dataPath)
            console.log(a)
            this.loadFromPath(dataPath)
        })
    }

    loadFromPath(dataPath) {
        let newFirmware = []
        console.log("[Firmware] Loading firmware from the data directory...")
        glob(path.join(dataPath, "firmware", "*", "*", "*.bin"), (err, files) => {
            if (err) {
                console.error("[Firmware] failed to find firmware files: ", err)
                process.exit(1)
            }

            console.log(files)

            for (let file of files) {
                let parts = file.split(path.sep)
                let version = parts[parts.length - 2]
                if (!semver.valid(version)) {
                    console.error(`Invalid version for file: ${file}`)
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