package inspect

var VCONSOLE = []byte(`<script src="https://unpkg.com/vconsole@latest/dist/vconsole.min.js"></script><script>var vConsole = new window.VConsole();</script>`)

var ERUDA = []byte(`<script src="//cdn.bootcdn.net/ajax/libs/eruda/2.3.3/eruda.js"></script><script>eruda.init();</script>`)

var MDEBUG = []byte(`<script>(function() {
    var scp = document.createElement('script');
    // 加载最新的mdebug版本
    scp.src = 'https://unpkg.com/mdebug@latest/dist/index.js';
    scp.async = true;
    scp.charset = 'utf-8';
    // 加载成功并初始化
    scp.onload = function() {
        mdebug.init();
    };
    // 加载状态切换回调
    scp.onreadystate = function() {};
    // 加载失败回调
    scp.onerror = function() {};
    document.getElementsByTagName('head')[0].appendChild(scp);
})();</script>`)
