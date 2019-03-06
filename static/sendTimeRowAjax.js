function SendTimeRow(){
    var ValueFromRowField = $('.timeRowField').val();

    for (var i = 0; i < ValueFromRowField.length; i++) {
        console.log(ValueFromRowField.charAt(i));
    }

    //$.ajax({
    //   url: "http://localhost:80/average",
    //    method: "POST",
    //    data : { sendedData: 'Hello'},
    //          success : function(data) {
    //              alert(data);
    //          },
    //});

    return false
}