$(document).ready(() => {
  $(".firmwareSelector").change(function () {
    let mac = $(this).parents("tr").attr("id")
    let type = $(this).val()
    $.ajax({
      url: `/api/device/${mac}?firmware=${type}`,
      type: "PATCH",
    })
  })

  $(".deleteFirmware.button").click(function () {
    let row = $(this).parents("tr")
    let filename = row.data("filename")
    let type = row.data("type")
    let version = row.data("version")

    if (confirm(`Are you sure you want to delete the firmare ${type} ${version}?\n\nThis cannot be undone.`)) {
      $.ajax({
        url: `/api/firmware/${type}/${version}/${filename}`,
        type: "DELETE",
        success: () => {
          row.remove()
        }
      })
    }
  })

  $(".deviceIdTextField").keyup(function (e) {
    if (e.keyCode == 13) {
      let mac = $(this).parents("tr").attr("id")
      let id = $(this).val()
      $.ajax({
        url: `/api/device/${mac}?id=${id}`,
        type: "PATCH"
      })
    }
  });
})

