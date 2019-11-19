$(document).ready(() => {
  $(".firmwareSelector").change(function () {
    let mac = $(this).parents("tr").attr("id")
    let type = $(this).val()
    $.post(`/api/assign?mac=${mac}&firmware=${type}`)
  })
})