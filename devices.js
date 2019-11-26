class Devices {
  constructor() {
    this.devices = {}
  }

  updateDevice(mac, firmwareType, firmwareVersion) {
    let device = this.getOrCreate(mac, "new")
    device.firmwareType = firmwareType
    device.firmwareVersion = firmwareVersion
    device.lastUpdated = new Date()
    return device
  }

  getOrCreate(mac, initialState) {
    if (!this.devices[mac]) {
      this.devices[mac] = {
        mac,
        state: initialState
      }
    }
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
  
  setDeviceId(mac, id) {
    console.log(`[Device] Setting ${mac} to device id ${id}`)
    this.getOrCreate(mac, "prepared")
    this.devices[mac].id = id
  }

  setState(mac, state) {
    console.log(`[Device] Setting ${mac} to state ${state}`)
    this.getOrCreate(mac, "prepared")
    this.devices[mac].state = state
  }

  assignFirmware(mac, type) {
    console.log(`[Device] Setting ${mac} to firmware ${type}`)
    this.getOrCreate(mac, "prepared")
    this.devices[mac].assignedFirmware = type
  }
}

module.exports = Devices