var lines = new Array();
var currentLine = 0;
function handleTextInput(key) {
    if (key == 13) {
        lines.push(document.getElementById('command-line').value);
        drawCommand('$ ' + document.getElementById('command-line').value );
        currentLine++;
        parseCommand();
        document.getElementById('command-line').value = '';
    }
}

function drawCommand(text) {
    document.getElementById('output').innerHTML += text + "<br/>";
}

function parseCommand() {
    var parts = lines[currentLine-1].split(' ');
    switch(parts[0]) {
        default:
            drawCommand(parts[0] + ' is not a valid command.');
    }
}
document.addEventListener("DOMContentLoaded", function () {
    document.getElementById('command-line').addEventListener('keypress', function (e) {
        handleTextInput(e.keyCode);
    });
    document.getElementById('command-line').focus();
});

