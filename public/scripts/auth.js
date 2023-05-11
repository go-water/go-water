$("#login").click(function(){
    let formData = new FormData($("#login").parents("form")[0]);
    $.ajax({
        async: true,
        cache: false,
        contentType: false,
        processData: false,
        type: "POST",
        dataType: "JSON",
        url: "/login",
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