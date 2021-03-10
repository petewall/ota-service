$(() => {
  $(".firmwareSelector").on("change", function () {
    const mac = $(this).parents("tr").attr("id")
    const type = $(this).val()
    $.ajax({
      url: `/api/device/${mac}?firmware=${type}`,
      type: "PATCH",
    })
  })

  $(".deleteFirmware.button").on("click", function () {
    const row = $(this).parents("tr")
    const filename = row.data("filename")
    const type = row.data("type")
    const version = row.data("version")

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
})
