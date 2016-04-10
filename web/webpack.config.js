var path = require('path');

module.exports = {
    context: path.join(__dirname, 'static/js/src'),
    entry: {
        article: './article.js'
    },
    output: {
        path: path.join(__dirname, 'static/js/dist'),
        filename: '[name].bundle.js'
    },
    module: {
        loaders: [{
            test: /\.js$/,
            exclude: /(node_modules|bower_components)/,
            loader: 'babel',
            query: {
                presets: ['react']
            }
        }]
    }
};