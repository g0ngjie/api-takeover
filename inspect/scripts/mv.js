// v2.0.2 Fri Sep 17 2021 20:39:01 GMT+0800 (China Standard Time) by Mdebug Team. Released under the MIT License.

;(function() {
    if (typeof window === 'undefined' || window['___apiTakeover.inspect-is-inited']) {
      return;
    }
    window['___apiTakeover.inspect-is-inited'] = true;
    var exports = {};
    var module = { exports: exports };

    document.addEventListener('DOMContentLoaded',  module.exports.init);
  })();