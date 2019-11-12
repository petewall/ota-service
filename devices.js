class Devices {
    constructor() {
        this.devices = {}
    }

    registerDevice(mac, firmwareType, firmwareVersion) {
        this.devices[mac] = {
            mac,
            firmwareType,
            firmwareVersion
        }
    }

    get(mac) {
        return this.devices[mac]
    }

    getAll() {
        return this.devices
    }
}

module.exports = Devices