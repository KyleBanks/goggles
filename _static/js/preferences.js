'use strict';

var PreferencesController = {

    $el: document.getElementById("preferences"),
    $txtGopath: document.getElementById("txt-prefs-gopath"),

    activate: function() {
        var $this = PreferencesController;
        $this.load();
    },

    deactivate: function() {
        var $this = PreferencesController;

        $this.$el.classList.add("hide");
    },

    load: function() {
        var $this = PreferencesController;

        $this.$el.classList.remove("hide");
        API.getPreferences($this._onLoad);
    },

    _onLoad: function(err, prefs) {
        var $this = PreferencesController;
        if (err) {
            console.error(err);
            return;
        }

        $this.$txtGopath.value = prefs.gopath;
    },

    save: function() {
        var $this = PreferencesController;

        API.setPreferences({
            gopath: $this.$txtGopath.value
        }, function() {
            $this.deactivate();
            State.set(State.PkgList);
        });
    }
};