function SendTimeRow(){
    var ValueFromRowField = $('.timeRowField').val();

    var array = [];
    var count = 0;
    var number = "";

    for (var i = 0; i < ValueFromRowField.length; i++) {
        var ch = ValueFromRowField.charAt(i);
        if(ch != " "){
            number = number + ch;
        } else{
            array[count] = number;
            count++;
            var number = "";
        }
    }
    number = number + ch;
    console.log(array)

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