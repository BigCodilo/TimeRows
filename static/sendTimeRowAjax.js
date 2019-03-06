function SendTimeRow(){
    var ValueFromRowField = $('.timeRowField').val();

    var obj = new object
    obj.array = [];
    var count = 0;
    var number = "";

    for (var i = 0; i < ValueFromRowField.length; i++) {
        var ch = ValueFromRowField.charAt(i);
        if(ch != " "){
            number = number + ch;
        } else {
            if(number != ""){
                obj.array[count] = number;
                count++;
                var number = "";
            }
        }
    }
    if(number != ""){
        obj.array[count] = number;
    } 
    $('.timeRowField').val("");
  
    var jsonArray = JSON.parse(obj.array);
    console.log(jsonArray);

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