import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
import path from "path";
import { viteStaticCopy } from "vite-plugin-static-copy";

const BUILD_DIR = path.join(__dirname, "..", "resources")

// https://vitejs.dev/config/
export default defineConfig({
  base: "./",
  build: {
    rollupOptions: {
      output: {
        entryFileNames: "assets/index.js",
        assetFileNames: 'assets/[name][extname]',
        chunkFileNames: 'assets/[name].js'
      }
    }
  },
  server: {
    proxy: {
      "/channel": "http://localhost:1234"
    }
  },
  plugins: [
    vue(),
    viteStaticCopy({
      targets: [
        {
          src: "dist/index.html",
          dest: path.join(BUILD_DIR, "panel")
        },
        {
          src: "dist/assets",
          dest: path.join(BUILD_DIR)
        }
      ]
    })
  ],
})
