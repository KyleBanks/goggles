'use strict';

var PkgDetailsController = {

    $el: document.getElementById("pkg-details"),
    $t: document.getElementById("t-pkg-details"),
    headingTemplates: [
        document.getElementById("t-pkg-details-heading"),
        document.getElementById("t-pkg-details-subheading")
    ],
    $functionT: document.getElementById("t-pkg-details-function"),
    $badgeT: document.getElementById("t-pkg-details-badge"),
    $actionT: document.getElementById("t-pkg-details-action"),

    _converter: new showdown.Converter({
        simplifiedAutoLink: true
    }),

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

    _prefs: null,
    _pkg: null,

    activate: function(data) {
        var $this = PkgDetailsController;
        $this.load(data.name);
    },

    openFileExplorer: function(name) {
        return API.openFileExplorer(name);
    },

    openTerminal: function(name) {
        return API.openTerminal(name);
    },

    openUrl: function(url) {
        return API.openUrl(url);
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

        API.getPreferences(function(err, prefs) {
            if (err) {
                console.error(err);
                return;
            }

            $this._prefs = prefs;
            API.getPkg(name, $this._onLoad);
        });
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
            fullName: res.name,
            repository: res.docs.repository,
            import: res.docs.import,
            content: $this._renderPkg(0, res.docs),
            badges: $this._renderBadges(res),
            actions: $this._renderActions(res)
        });

        // Bind events
        $this._bindCollapsableEvents();

        // Prettyprint code
        var pres = $this.$el.getElementsByTagName("pre");
        for (var i = 0; i < pres.length; i++) {
            pres[i].classList.add("prettyprint");
        }
        PR.prettyPrint();
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
                    content: $this._sections[s].code ?
                        "<pre><code>" + c + "</code></pre>" : $this._converter.makeHtml(c)
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

    /**
     * Renders the README-style badges for the package/repository.
     *
     * @param pkg {Object}
     */
    _renderBadges: function(pkg) {
        var $this = PkgDetailsController,
            badges = [];

        badges.push(
            Template.apply($this.$badgeT, {
                url: "https://godoc.org/" + pkg.name,
                image: "https://godoc.org/" + pkg.name + "?status.svg",
                width: 109,
                height: 20
            })
        );

        badges.push(
            Template.apply($this.$badgeT, {
                url: "https://goreportcard.com/report/" + pkg.docs.repository,
                image: "https://goreportcard.com/badge/" + pkg.docs.repository,
                width: 88,
                height: 20
            })
        );

        if (pkg.docs.hasTravis) {
            var user = pkg.name.split("/")[1],
                repo = pkg.name.split("/")[2];

            badges.push(
                Template.apply($this.$badgeT, {
                    url: "https://travis-ci.org/" + user + "/" + repo,
                    image: "https://travis-ci.org/" + user + "/" + repo + ".svg?branch=master",
                    width: 90,
                    height: 20
                })
            );
        }

        return badges.join("");
    },

    /**
     * Renders the action buttons for a package.
     * 
     * @param pkg {Object}
     */
    _renderActions: function(pkg) {
        var $this = PkgDetailsController,
            actions = [];

        actions.push(
            Template.apply($this.$actionT, {
                method: "openFileExplorer",
                param: pkg.name,
                text: "Folder"
            })
        );

        if ($this._prefs.canOpenTerminal) {
            actions.push(
                Template.apply($this.$actionT, {
                    method: "openTerminal",
                    param: pkg.name,
                    text: "Terminal"
                })
            );
        }

        actions.push(
            Template.apply($this.$actionT, {
                method: "openUrl",
                param: pkg.docs.repository,
                text: "Repo"
            })
        );

        return actions.join("");
    },

    /**
     * Binds click events for all the "collapsable" sections to open/close
     * content when the header is clicked.
     */
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

    /**
     * Toggles visibility of a "collapsable" section.
     * 
     * @param collapsable {Element}
     */
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