'use strict';

var Toast = {

    $t: document.getElementById("t-toast"),
    $el: document.getElementById("toasts"),

    show: function(message) {
        var $this = Toast;

        var id = "toast-" + new Date().getTime();
        $this.$el.innerHTML += Template.apply($this.$t, {
            id: id,
            message: message
        });

        setTimeout(function() {
            document.getElementById(id).parentNode.classList.add('fade');
        }, 5000);
    }

};