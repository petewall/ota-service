$(document).ready(() => {
  $(".firmwareSelector").change(function () {
    let mac = $(this).parents("tr").attr("id")
    let type = $(this).val()
    $.post(`/api/assign?mac=${mac}&firmware=${type}`)
  })

  $(".deleteFirmware.button").click(function () {
    let row = $(this).parents("tr")
    let filename = row.data("filename")
    let type = row.data("type")
    let version = row.data("version")
    $.ajax({
      url: `/api/firmware/${type}/${version}/${filename}`,
      type: "DELETE",
      success: () => {
        row.remove()
      }
    })
  })
})

