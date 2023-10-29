$(document).ready(function () {
  // our clickme button action
  $("#clickme").click(function () {
    $("#PayloadCutup").removeClass("required_red");

    $.ajax({
      url: "message",
      dataType: "json",
      type: "post",
      contentType: "application/json",
      data: JSON.stringify({
        cutup: $("#cutup").val(),
      }),
      processData: false,
      success(data) {
        console.log("data: ", data);
        var returnedData = data;
        $.each(returnedData, function (index, item) {
          // check if field cutup in data
          if (index == "cutup") {
            console.log("cutup is in data");
            $("textarea#cutup").val(item);
          } else {
            console.log("index: " + index);
            console.log(index, item, item["FailedField"].split(".")[0]);
            var FailedField =
              "#" +
              item["FailedField"].split(".")[0] +
              item["FailedField"].split(".")[1];
            console.log("FailedField: " + FailedField);
            $(FailedField).addClass("required_red");
          }
        });
      },
      error: function (jqXhr, textStatus, errorThrown) {
        console.log(errorThrown);
      },
    });
  });
});
