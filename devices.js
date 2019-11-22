class Devices {
  constructor() {
    this.devices = {}
  }

  updateDevice(mac, firmwareType, firmwareVersion) {
    if (!this.devices[mac]) {
      this.devices[mac] = {
        mac,
        assignedFirmware: firmwareType,
        state: "new"
      }
    }
    this.devices[mac].firmwareType = firmwareType
    this.devices[mac].firmwareVersion = firmwareVersion
    this.devices[mac].lastUpdated = new Date()
    return this.devices[mac]
  }

  get(mac) {
    return this.devices[mac]
  }

  getAll(sorted = false) {
    let array = Object.values(this.devices)
    if (sorted) {
      return array.sort((a, b) => {
        if (a.mac < b.mac) {
          return -1
        } else if (a.mac > b.mac) {
          return 1
        }
        return 0
      })
    }
    return array
  }

  setState(mac, state) {
    console.log(`[Device] Setting ${mac} to state ${state}`)
    this.devices[mac].state = state
  }

  assignFirmware(mac, type) {
    console.log(`[Device] Setting ${mac} to firmware ${type}`)
    this.devices[mac].assignedFirmware = type
  }
}

module.exports = Devices