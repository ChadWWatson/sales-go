var commandLine = (function () {
    var control = {};
    
    return {
        init: init,
    }
    
    function init(control) {
        console.log('this.control',this.control);
        console.log('control',control);
        
        control.addEventListener('keypress', function(e) {
           handleTextInput(e.key); 
        });
    }
    
    
    function handleTextInput(key) {
        console.log(key);
    }
})