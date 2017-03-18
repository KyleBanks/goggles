'use strict';

var Template = {

    /**
     * Given a template element, applies the data provided and returns a copy.
     *
     * @param template {Element}
     * @param data {String}
     */
    apply: function(template, data) {
        var t = template.cloneNode(true),
            tmp = document.createElement("div");

        t.removeAttribute("id");
        tmp.appendChild(t);

        var contents = tmp.innerHTML,
            keys = Object.keys(data);

        for (var i = 0; i < keys.length; i++) {
            var prop = keys[i],
                value = data[prop],
                regex = new RegExp("{{ ?" + prop + " ?}}", 'g');

            contents = contents.replace(regex, value);
        }

        return contents;
    }

};