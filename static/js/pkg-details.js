'use strict';

var PkgDetailsController = {

    $el: document.getElementById("pkg-details"),
    $t: document.getElementById("t-pkg-details"),
    headingTemplates: [
        document.getElementById("t-pkg-details-heading"),
        document.getElementById("t-pkg-details-subheading")
    ],
    $functionT: document.getElementById("t-pkg-details-function"),

    _converter: new showdown.Converter(),
    _sections: [{
            title: "",
            key: "declaration",
            code: true
        },
        {
            title: "",
            key: "usage",
            noCollapse: true
        },
        {
            title: "Constants",
            key: "constants",
            code: true
        },
        {
            title: "Variables",
            key: "variables",
            code: true
        }
    ],
    _pkg: null,

    activate: function(data) {
        var $this = PkgDetailsController;

        $this.load(data.name);
    },

    /**
     * Loads package details by the full package name (ex. github.com/foo/bar).
     * 
     * @param name {String}
     */
    load: function(name) {
        var $this = PkgDetailsController;

        if (name === $this._pkg) {
            return;
        }
        $this._pkg = name;

        Loader.show();
        API.getPkg(name, $this._onLoad);
    },

    _onLoad: function(err, res) {
        var $this = PkgDetailsController;
        if (err) {
            console.error(err);
            return;
        } else if ($this._pkg !== res.name) {
            return;
        }

        // Render package docs
        $this.$el.innerHTML = Template.apply($this.$t, {
            name: res.docs.name,
            import: res.docs.import,
            content: $this._renderPkg(0, res.docs)
        });

        // Bind events
        $this._bindCollapsableEvents();

        // Prettyprint code
        var pres = $this.$el.getElementsByTagName("pre");
        for (var i = 0; i < pres.length; i++) {
            pres[i].classList.add("prettyprint");
        }
        PR.prettyPrint();

        Loader.hide();
    },

    _renderPkg: function(headingNum, pkg) {
        var $this = PkgDetailsController,
            content = [];

        for (var s = 0; s < $this._sections.length; s++) {
            var c = pkg[$this._sections[s].key];
            if (!c || c.length === 0) {
                continue;
            }

            content.push(
                $this._renderSection(headingNum, {
                    noCollapse: $this._sections[s].noCollapse,
                    title: $this._sections[s].title,
                    content: $this._sections[s].code ? "<pre><code>" + c + "</code></pre>" : $this._converter.makeHtml(c)
                })
            );
        }

        content.push(
            $this._renderContent(headingNum, pkg['content'])
        );

        return content.join("");
    },

    _renderContent: function(headingNum, content) {
        if (!content) {
            return "";
        }

        var $this = PkgDetailsController,
            res = [];

        // content.sort(function(a, b) {
        //     return (a.name > b.name) ? 1 : -1;
        // });

        for (var c = 0; c < content.length; c++) {
            var isType = content[c].type === "TYPE";
            res.push(
                $this._renderSection(headingNum, {
                    title: content[c].header,
                    content: isType ? $this._renderPkg(headingNum + 1, content[c]) : Template.apply($this.$functionT, content[c])
                })
            )
        }

        return res.join("");
    },

    /**
     * Renders an individual section.
     *
     * @param headingNum {Integer}
     * @param data {Object}
     * @param data.noCollapse {Boolean}
     * @param data.title {String}
     * @param data.content {String}
     */
    _renderSection: function(headingNum, data) {
        var $this = PkgDetailsController;

        return Template.apply($this.headingTemplates[headingNum], {
            class: data.noCollapse ? "" : "collapsable",
            titleClass: data.title.length ? "" : "hide",
            title: data.title,
            content: data.content
        });
    },

    _bindCollapsableEvents: function() {
        var $this = PkgDetailsController;

        var sections = $this.$el.getElementsByClassName("collapsable");
        for (var i = 0; i < sections.length; i++) {
            var section = sections[i],
                trigger = section.getElementsByClassName("trigger")[0],
                content = section.getElementsByClassName("content")[0];

            trigger.onclick = $this._toggleCollapsableVisibility.bind($this, section);
        }
    },

    _toggleCollapsableVisibility: function(collapsable) {
        var $this = PkgDetailsController,
            content = collapsable.getElementsByClassName("content")[0];

        if (collapsable.classList.contains("open")) {
            collapsable.classList.remove("open");
            content.style.maxHeight = "0px";
            return;
        }

        collapsable.classList.add("open");
        content.style.maxHeight = content.scrollHeight + "px";
    }

};