'use strict';

var State = {

    PkgList: "pkg-list",
    PkgDetails: "pkg-details",
    Preferences: "preferences",

    set: function(state, data) {
        var $this = State;
        console.log("State.set(" + state + ", " + JSON.stringify(data) + ")");

        var controller;
        switch (state) {
            case State.PkgList:
                controller = PkgListController;
                break;
            case State.PkgDetails:
                controller = PkgDetailsController;
                break;
            case State.Preferences:
                controller = PreferencesController;
                break;
        }

        controller.activate(data);
    },

};

State.set(State.PkgList);