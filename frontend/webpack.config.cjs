const path = require('path');
const webpack = require('webpack');
const TerserPlugin = require('terser-webpack-plugin');

module.exports = {
  entry: './src/main.js',
  output: {
    path: path.resolve(__dirname, '..', 'backend', 'static', 'js'),
    filename: 'bundle.js'
  },
  plugins: [new webpack.ProgressPlugin()],
  module: {
    rules: [{
      test: /\.(js|jsx)$/,
      include: [path.resolve(__dirname, 'src')],
      loader: 'babel-loader',
      options: {
          presets: ['@babel/preset-env'],
          plugins: ['@babel/plugin-proposal-class-properties']
      }
    }]
  },
  mode: "production",
  optimization: {
   usedExports: true,
  },
  node: {
    // prevent webpack from injecting eval / new Function through global polyfill
    global: false,
  },
};
