'use strict';

var PkgDetailsController = {

    $el: document.getElementById("pkg-details"),

    activate: function(data) {
        var $this = PkgDetailsController;

        $this.load(data.name);
    },

    load: function(pkg) {
        var $this = PkgDetailsController;

        Loader.show();
        $this.$el.innerHTML = pkg;
        // TODO
    }

};