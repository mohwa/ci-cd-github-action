'use strict';
const path = require('path');
const webpack = require('webpack');
const HtmlWebPackPlugin = require('html-webpack-plugin');
const HtmlWebpackInjector = require('html-webpack-injector');

const rootPath = path.resolve('.');
const outputPath = path.resolve('./dist');

module.exports = () => {
    return {
        context: rootPath,
        mode: 'development',
        target: 'web',
        entry: {
            'app': path.resolve('src/index.js'),
        },
        output: {
            path: outputPath,
            publicPath: '/',
        },
        devtool: 'source-map',
        resolve: {
            modules: ['src', 'node_modules'],
            extensions: ['.js', '.jsx', '.json', 'css'],
        },
        module: {
            rules: [
                {
                    test: /\.(js|jsx)$/,
                    use: ['babel-loader'],
                    include: [path.resolve('src')],
                    exclude: /node_modules/,
                },
            ],
        },
        plugins: [
            new webpack.ProgressPlugin(),
            new HtmlWebPackPlugin({
                template: 'src/index.html',
                filename: 'index.html',
            }),
            new HtmlWebpackInjector(),
            new webpack.DefinePlugin({}),
        ],
        optimization: {
            runtimeChunk: {
                name: 'runtime',
            },
            splitChunks: {
                minSize: 30000,
                maxSize: 1010000,
                minChunks: 1,
                cacheGroups: {
                    vendor: {
                        test: /[\\/]node_modules[\\/](react|react-dom|core-js)[\\/]/,
                        name: 'vendor',
                        chunks: 'all',
                    },
                    defaultVendors: {
                        reuseExistingChunk: true,
                    },
                },
            },
            minimize: false,
        },
        devServer: {
            hot: true,
            port: 3000,
            historyApiFallback: {
                index: '/index.html',
            },
            // https://github.com/gitpod-io/gitpod/issues/26
            allowedHosts: 'all',
            proxy: {
                '/api/rest': 'http://localhost:3001',
                '/api/graphql/query': 'http://localhost:3001'
            }
        },
    };
};
