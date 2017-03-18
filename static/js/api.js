'use strict';

var API = {

    BASE: "/api",

    /**
     * Sends a GET request to the API.
     *
     * @param path {String}
     * @param cb {function(Error, Object)}
     */
    _get: function(path, cb) {
        cb = cb || function() {};
        var $this = API;
        path = $this.BASE + path;

        var xmlhttp = new XMLHttpRequest();
        xmlhttp.onreadystatechange = function() {
            if (xmlhttp.readyState != XMLHttpRequest.DONE) {
                return;
            }

            if (xmlhttp.status == 200) {
                cb(null, JSON.parse(xmlhttp.responseText));
            } else if (xmlhttp.status == 400) {
                cb(new error(xmlhttp.status + ": " + xmlhttp.responseText));
            }
        };

        console.log("GET", path);
        xmlhttp.open("GET", path, true);
        xmlhttp.send();
    },

    /**
     * Loads the root package list.
     * 
     * @param cb {function(Error, Object)}
     */
    loadPkgList: function(cb) {
        return API._get("/pkg/list", cb);
    },

    /**
     * Sends a request to open the browser dev tools.
     */
    openDevTools: function() {
        return API._get("/debug");
    }

};