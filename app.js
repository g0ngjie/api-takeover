const express = require('express')
const cors = require('cors')
const bodyParser = require('body-parser')

const { createProxyMiddleware } = require('http-proxy-middleware');
const { responseInterceptor } = require('http-proxy-middleware');

const app = express()
app.use(cors())
const port = 3000

// app.use((req) => {
//     const start = new Date()
//     const ms = new Date() - start
//     console.log(`${req.method} ${req.url} - ${ms}ms`)
// })

app.use(createProxyMiddleware({
    target: 'https://gw-mobile-test.beantechyun.cn',
    changeOrigin: true,
    // ws: true,j
    logLevel: "debug",
    selfHandleResponse: true,
    // onProxyRes(proxyRes, req, res) {
    //     // console.log("[debug]proxyRes.statusMessage:", proxyRes.statusMessage)
    //     // console.log("[debug]res.statusMessage:", res.statusMessage)
    //     // console.log("[debug]req.url:", req.url)
    //     // console.log("[debug]req:", req)
    //     proxyRes.on("data", (data) => {
    //         console.log("received data as stream");
    //         const bufferAsString = data.toString("utf-8")
    //         console.log(bufferAsString);
    //     });
    // }
    onProxyRes: responseInterceptor(async (responseBuffer, proxyRes, req, res) => {
        // setTimeout(() => req.next())     // call next() after returning `responseBuffer`
        const json = responseBuffer.toString("utf-8")
        console.log("[debug]json:", json)
        return responseBuffer;
    })
}));

app.listen(port, () => {
    console.log(`Example app listening on port ${port}`)
})