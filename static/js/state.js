'use strict';

var State = {

    $controllers: document.getElementById("controllers"),

    PkgList: "pkg-list",
    PkgDetails: "pkg-details",


    set: function(state, data) {
        var $this = State;
        console.log("State.set(" + state + ", " + JSON.stringify(data) + ")");

        $this.$controllers.classList.remove("hide");

        var controller;
        switch (state) {
            case State.PkgList:
                controller = PkgListController;
                break;
            case State.PkgDetails:
                controller = PkgDetailsController;
                break;
        }
        controller.activate(data);
    },

};

State.set(State.PkgList);