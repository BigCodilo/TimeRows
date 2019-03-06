function SendTimeRow(){
    var ValueFromRowField = $("#time-row-field").val();

    for (var i = 0; i < ValueFromRowField.length; i++) {
        console.console.log(ValueFromRowField.charAt(i));
    }

    return false
}