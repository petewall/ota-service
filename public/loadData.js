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
        allFirmware.forEach(addFirmware)
    }, "json")
    $("#firmwareTable .loading").remove()
}

function addDevice(device) {
    $("#deviceTable tbody").append(
        $("<tr>").append(
            $("<td>", { text: device.mac }),
            $("<td>", { text: device.firmwareType }),
            $("<td>", { text: device.firmwareVersion }),
            $("<td>").append(
                $("<input>", { type: "text", val: device.sensorId })
            )
        )
    )
}

function loadDevices() {
    $.get("/api/devices", devices => {
        for (let mac in devices) {
            addDevice(devices[mac])
        }
    }, "json")
    $("#deviceTable .loading").remove()
}

function updateFirmwareCounts() {
    
}

$(document).ready(() => {
    loadFirmware()
    loadDevices()
    updateFirmwareCounts()
})
