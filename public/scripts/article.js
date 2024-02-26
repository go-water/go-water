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

$(document).ready(function(){
    let easyMDE = new EasyMDE({
        element: document.getElementById('MyID'),
        status: ["autosave", "lines", "words", "upload-image"],
        spellChecker: false,
        toolbar: ['undo', 'redo', '|', 'bold', 'italic', 'strikethrough', 'heading-1', 'heading-2', 'heading-3', '|', 'image', 'code', 'quote', 'link', 'clean-block', 'ordered-list', 'unordered-list', '|', 'upload-image', 'table','preview', 'side-by-side', 'fullscreen', '|', 'guide'],
        imageUploadEndpoint: "/api/upload",
        imagePathAbsolute: true,
        imageCSRFHeader: true,
        imageCSRFName: "Authorization",
        imageCSRFToken: 'Bearer ' + localStorage.getItem('token'),
        uploadImage: true,
    });

    console.log(document.getElementById('MyID'));

    easyMDE.codemirror.on("change", () => {
        $("#MyID").val(easyMDE.value());
    });
});