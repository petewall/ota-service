devices = []
firmware = []

function firmwareDropdownList() {
    let set = new Set(this.firmware.map(firmware => firmware.type))
    let types = Array.from(set.values()).sort()
    let dropdown = $("<select>")
    for (let type of types) {
        dropdown.append(
            $("<option>", {
                text: type,
                val: type
            })
        )
    }
    return dropdown
}

function addDevice(device) {
    $("#deviceTable tbody").append(
        $("<tr>").append(
            $("<td>", { text: device.mac }),
            $("<td>", { text: device.firmwareType }),
            $("<td>", { text: device.firmwareVersion }),
            $("<td>", { text: "no" }),
            $("<td>", { text: new Date().toLocaleString() }),
            firmwareDropdownList()
        )
    )
}

function loadDevices() {
    $.get("/api/devices", devices => {
        devices = allDevices
        for (let mac in allDevices) {
            addDevice(allDevices[mac])
        }
    }, "json")
    $("#deviceTable .loading").remove()
}

function addFirmware(firmware) {
    $("#firmwareTable tbody").append(
        $("<tr>").append(
            $("<td>", { text: firmware.type }),
            $("<td>", { text: firmware.version }),
            $("<td>", { text: firmware.filename }),
            $("<td>", { text: "---" })
        )
    )
}

function loadFirmware() {
    $.get("/api/firmware", (allFirmware) => {
        firmware = allFirmware
        allFirmware.forEach(addFirmware)
    }, "json")
    $("#firmwareTable .loading").remove()
}

$(document).ready(() => {
    loadFirmware()
    loadDevices()
    updateFirmwareCounts()
})
