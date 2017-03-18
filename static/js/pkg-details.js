'use strict';

var PkgDetailsController = {

    _converter: new showdown.Converter(),

    $el: document.getElementById("pkg-details"),
    $t: document.getElementById("t-pkg-details"),
    $typeT: document.getElementById("t-pkg-details-type"),

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

        var types = [];
        if (res.docs.types) {
            for (var t = 0; t < res.docs.types.length; t++) {
                var ty = res.docs.types[t];
                types.push(
                    Template.apply($this.$typeT, {
                        name: ty.name,
                        header: ty.header,
                        constants: $this._converter.makeHtml(ty.constants),
                        variables: $this._converter.makeHtml(ty.variables),
                        functions: $this._converter.makeHtml(ty.functions)
                    })
                );
            }
        }

        $this.$el.innerHTML = Template.apply($this.$t, {
            name: res.docs.name,
            import: res.docs.import,
            package: $this._converter.makeHtml(res.docs.package),
            constants: $this._converter.makeHtml(res.docs.constants),
            variables: $this._converter.makeHtml(res.docs.variables),
            functions: $this._converter.makeHtml(res.docs.functions),
            types: types.join("")
        });

        var pres = $this.$el.getElementsByTagName("pre");
        for (var i = 0; i < pres.length; i++) {
            pres[i].classList.add("prettyprint");
        }
        PR.prettyPrint();

        Loader.hide();
    },

};