'use strict';

var API = {

    BASE: "/api",

    /**
     * Loads the root package list.
     * 
     * @param cb {function(Error, Object)}
     */
    loadPkgList: function(cb) {
        return API._get("/pkg/list", cb);
    },

    /**
     * Loads the full details of a package.
     *
     * @param name {String}
     * @param cb {function(Error, Object)}
     */
    getPkg: function(name, cb) {
        return API._get("/pkg/details?name=" + encodeURIComponent(name), cb);
    },

    /**
     * Opens the system file explorer to the package name provided.
     *
     * @param name {String}
     * @param cb {function(Error, Object)}
     */
    openFileExplorer: function(name, cb) {
        return API._get("/open/file-explorer?name=" + encodeURIComponent(name), cb);
    },

    /**
     * Opens the system terminal to the package name provided.
     *
     * @param name {String}
     * @param cb {function(Error, Object)}
     */
    openTerminal: function(name, cb) {
        return API._get("/open/terminal?name=" + encodeURIComponent(name), cb);
    },

    /**
     * Opens the default browser to the specified URL.
     *
     * @param url {String}
     * @param cb {function(Error, Object)}
     */
    openUrl: function(url, cb) {
        return API._get("/open/url?url=" + encodeURIComponent(url), cb);
    },

    /**
     * Returns the Goggles preferences.
     * 
     * @param cb {function(Error, Object)}
     */
    getPreferences: function(cb) {
        return API._get("/preferences/", cb);
    },

    /**
     * Updates Goggles preferences.
     * 
     * @param prefs {Object}
     * @param cb {function(Error, Object)}
     */
    setPreferences: function(prefs, cb) {
        return API._get("/preferences/update?gopath=" + encodeURIComponent(prefs.gopath), cb);
    },

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

        Loader.show();

        var xmlhttp = new XMLHttpRequest();
        xmlhttp.onreadystatechange = function() {
            if (xmlhttp.readyState != XMLHttpRequest.DONE) {
                return;
            }

            console.log("(" + xmlhttp.status + ") " + xmlhttp.responseText);
            if (xmlhttp.status == 200) {
                cb(null, JSON.parse(xmlhttp.responseText));
            } else if (xmlhttp.status == 400) {
                cb(new error(xmlhttp.status + ": " + xmlhttp.responseText));
            }

            Loader.hide();
        };

        console.log("GET: " + path);
        xmlhttp.open("GET", path, true);
        xmlhttp.send();
    },


};