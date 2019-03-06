function SendTimeRow(){
    var ValueFromRowField = $('.timeRowField').val();

    var sendMessage = new object();
    sendMessage.array = [];
    var count = 0;
    var number = "";

    for (var i = 0; i < ValueFromRowField.length; i++) {
        var ch = ValueFromRowField.charAt(i);
        if(ch != " "){
            number = number + ch;
        } else {
            if(number != ""){
                sendMessage.array[count] = number;
                count++;
                var number = "";
            }
        }
    }
    if(number != ""){
        sendMessage.array[count] = number;
    } 
    $('.timeRowField').val("");
  
    sendMessage.len = sendMessage.array.length;
    var jsonArray = JSON.stringify(sendMessage);
    console.log(jsonArray);

    $.ajax({
       url: "http://localhost:80/average",
        method: "POST",
        data : { sendedData: jsonArray}
    });

    return false
}