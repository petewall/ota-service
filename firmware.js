const chokidar = require("chokidar")
const fs = require("fs").promises
const util = require("util")
const glob = util.promisify(require("glob"))
const mkdir = util.promisify(require("mkdirp"))
const path = require("path")
const semver = require("semver")

class Firmware {
  constructor(dataPath) {
    this.firmware = []
    this.firmwarePath = path.join(dataPath, "firmware")
    this.firmwareGlob = path.join(dataPath, "firmware", "*", "*", "*.bin")
    this.loadFromPath()

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

  async loadFromPath() {
    let newFirmware = []
    console.log("[Firmware] Loading firmware from the data directory...")
    let firmwarePaths = await glob(this.firmwareGlob)

    for (let firmwarePath of firmwarePaths) {
      let parts = firmwarePath.split(path.sep)
      let version = parts[parts.length - 2]
      if (!semver.valid(version)) {
        console.error(`[Firmware] Invalid version for file: ${firmwarePath}`)
      } else {
        console.log(`[Firmware]    ${firmwarePath}`)

        let stats = await fs.stat(firmwarePath)
        newFirmware.push({
          file: firmwarePath,
          filename: parts[parts.length - 1],
          size: stats.size,
          type: parts[parts.length - 3],
          version
        })
      }
    }
    this.firmware = newFirmware
    console.log(`[Firmware] Firmware loaded: ${this.firmware.length} binaries`)
  }

  async addBinary(firmwareType, version, data) {
    let directory = path.join(this.firmwarePath, firmwareType, version)
    let firmware = path.join(directory, `${firmwareType}-${version}.bin`)

    console.log(`[Firmware] Writing firmware binary file: ${firmware}`)
    await mkdir(directory)
    await fs.writeFile(firmware, data)
  }

  getAll() {
    return this.firmware
  }

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