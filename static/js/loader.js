'use strict';

var Loader = {

    $el: document.getElementById("loader"),

    _images: ["/static/img/loader-1.png", "/static/img/loader-2.png", "/static/img/loader-3.png"],
    _imgIdx: 0,

    _showCount: 0,

    _init: function() {
        var $this = Loader;
        setInterval($this._swapImg, 400);
    },

    show: function() {
        var $this = Loader;
        $this._showCount++;
        $this._update();
    },

    hide: function() {
        var $this = Loader;
        $this._showCount--;
        $this._update();
    },

    _update: function() {
        var $this = Loader;

        if ($this._showCount > 0) {
            $this.$el.classList.remove("hide");
        } else {
            $this.$el.classList.add("hide");
        }
    },

    /**
     * Moves the loader image to the next in the list.
     */
    _swapImg: function() {
        var $this = Loader;

        $this._imgIdx++;
        if ($this._imgIdx >= $this._images.length) {
            $this._imgIdx = 0;
        }

        $this.$el.getElementsByTagName('img')[0]
            .setAttribute("src", $this._images[$this._imgIdx]);
    }
};
Loader._init();