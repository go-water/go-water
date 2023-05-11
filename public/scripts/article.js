$("#publish").click(function(){
    let formData = new FormData($("#publish").parents("form")[0]);
    $.ajax({
        async: true,
        cache: false,
        contentType: false,
        processData: false,
        type: "POST",
        dataType: "JSON",
        url: "/admin/add",
        data: formData,
        error: function (err) {
            alert(err.responseJSON.msg);
        },
        success: function (data) {
            if (data.result) {
                location.href="/admin/list";
            }
        }
    });
});

$("#update").click(function(){
    let formData = new FormData($("#update").parents("form")[0]);
    $.ajax({
        async: true,
        cache: false,
        contentType: false,
        processData: false,
        type: "POST",
        dataType: "JSON",
        url: "/admin/update",
        data: formData,
        error: function (err) {
            alert(err.responseJSON.msg);
        },
        success: function (data) {
            if (data.result) {
                location.href="/admin/list";
            }
        }
    });
});