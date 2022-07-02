const {defineConfig} = require('@vue/cli-service')
module.exports = defineConfig({
    transpileDependencies: true,
    chainWebpack: config => {
        config.plugin('define').tap(definitions => {
            Object.assign(definitions[0]['process.env'], {
                NODE_HOST: '"http://localhost:8888"',
            });
            return definitions;
        });
    },
    devServer: {
        port: 8082,  // 端口号的配置
        host:'localhost',
        open: true   // 自动打开浏览器
    }
})
