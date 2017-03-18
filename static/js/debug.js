'use strict';

var Debug = {

    _init: function() {
        var $this = Debug;

        // Listen for CMD+d
        document.onkeydown = function(e) {
            if (e.metaKey && e.keyCode === 68) {
                $this._open();
            }
        };
    },

    _open: function() {
        API.openDevTools();
    }

};
Debug._init();