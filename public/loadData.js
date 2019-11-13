devices = []
firmware = []

function firmwareDropdownList(mac, selected) {
  let set = new Set(this.firmware.map(firmware => firmware.type))
  let types = Array.from(set.values()).sort()
  let dropdown = $("<select>").append($("<option>"))
  for (let type of types) {
    dropdown.append(
        $("<option>", {
        selected: (type == selected ? "selected" : null),
        text: type,
        val: type
      })
    )
  }

  $(dropdown).change(() => {
    $.post(`/api/assign?mac=${mac}&firmware=${$(dropdown).val()}`)
  })

  return dropdown
}

function addDevice(device) {
  $("#deviceTable tbody").append(
    $("<tr>", {
      id: device.mac
    }).append(
      $("<td>", { text: device.mac }),
      $("<td>", { text: device.firmwareType }),
      $("<td>", { text: device.firmwareVersion }),
      $("<td>", { text: device.lastUpdated }),
      $("<td>", { text: device.state }),
      $("<td>", { text: new Date().toLocaleString() }),
      firmwareDropdownList(device.mac, device.assignedFirmware)
    )
  )
}

function loadDevices() {
  $("#deviceTable").addClass("loading")
  return $.get("/api/devices", allDevices => {
    devices = allDevices
    for (let mac in allDevices) {
      addDevice(allDevices[mac])
    }
    $("#deviceTable").removeClass("loading")
  }, "json")
}

function addFirmware(firmware) {
  $("#firmwareTable tbody").append(
    $("<tr>").append(
      $("<td>", { text: firmware.type }),
      $("<td>", { text: firmware.version }),
      $("<td>", { text: firmware.filename }),
    )
  )
}

function loadFirmware() {
  $("#firmwareTable").addClass("loading")
  return $.get("/api/firmware", allFirmware => {
    firmware = allFirmware
    allFirmware.forEach(addFirmware)
    $("#firmwareTable").removeClass("loading")
  }, "json")
}

$(document).ready(async () => {
  await loadFirmware()
  await loadDevices()
})
