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

var TAG_SCRIPT = []byte(`<script>
    const wrap = document.createElement("div")
    wrap.style.cssText = '\
        width: 100px;\
        height: 18px;\
        position: fixed;\
        top: 0;\
        left: 0;\
        display: flex;\
        align-items: center;\
        justify-content: center;\
        border-radius: 0 2px 2px 0;\
        color: #fff;\
        font-size: 12px;\
        border: 1px solid rgba(252, 5, 5, 0.2);\
        box-shadow: 0 0 8px rgba(252, 5, 5, 0.6);\
        background-color: rgba(252, 5, 5, 0.45);'
    wrap.innerText = "api takeover"
    document.documentElement.appendChild(wrap)
</script>`)
