'use strict';

var PkgDetailsController = {

    _converter: new showdown.Converter(),
    _sections: [{
            title: "",
            key: "declaration",
            code: true
        },
        {
            title: "",
            key: "usage"
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

    $el: document.getElementById("pkg-details"),
    $t: document.getElementById("t-pkg-details"),
    $headingT: document.getElementById("t-pkg-details-heading"),
    $functionT: document.getElementById("t-pkg-details-function"),

    activate: function(data) {
        var $this = PkgDetailsController;

        $this.load(data.name);
    },

    load: function(pkg) {
        var $this = PkgDetailsController;

        Loader.show();
        API.getPkg(pkg, $this._onLoad);
    },

    _onLoad: function(err, res) {
        var $this = PkgDetailsController;
        if (err) {
            console.error(err);
            return;
        }

        $this.$el.innerHTML = Template.apply($this.$t, {
            name: res.docs.name,
            import: res.docs.import,
            content: $this._renderPkg(2, res.docs)
        });

        var pres = $this.$el.getElementsByTagName("pre");
        for (var i = 0; i < pres.length; i++) {
            pres[i].classList.add("prettyprint");
        }
        PR.prettyPrint();

        Loader.hide();
    },

    _renderPkg: function(headingNum, t) {
        var $this = PkgDetailsController,
            content = [];

        for (var s = 0; s < $this._sections.length; s++) {
            var c = t[$this._sections[s].key];
            if (!c || c.length === 0) {
                continue;
            }

            content.push(
                Template.apply($this.$headingT, {
                    headingnum: headingNum,
                    title: $this._sections[s].title,
                    content: $this._sections[s].code ? "<pre><code>" + c + "</code></pre>" : $this._converter.makeHtml(c)
                })
            );
        }
        content.push(
            $this._renderContent(headingNum, t['content'])
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
            var isType = content[c].contentType === "TYPE";
            res.push(
                Template.apply($this.$headingT, {
                    headingnum: headingNum,
                    title: content[c].header,
                    content: isType ? $this._renderPkg(headingNum + 1, content[c]) : Template.apply($this.$functionT, content[c])
                })
            )
        }

        return res.join("");
    },


};