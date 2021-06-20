const path = require('path');
const webpack = require('webpack');
const MiniCssExtractPlugin = require('mini-css-extract-plugin');

module.exports = {
  entry: './src/js/main.js',
  output: {
    path: path.resolve(__dirname, '..', 'backend', 'static'),
    filename: "js/bundle.js"
  },
  plugins: [new webpack.ProgressPlugin(), new MiniCssExtractPlugin({filename: "css/bundle.css"})],
  module: {
    rules: [{
      test: /\.(js|jsx)$/,
      include: [path.resolve(__dirname, 'src')],
      loader: 'babel-loader',
      options: {
          presets: ['@babel/preset-env'],
          plugins: ['@babel/plugin-proposal-class-properties']
      }
    },
    {
      test: /\.css$/i,
      use: [MiniCssExtractPlugin.loader, "css-loader", "sass-loader"],
    }, {
      test: /\.scss$/i,
      use: [MiniCssExtractPlugin.loader, "css-loader", "sass-loader"],
    }]
  },
  mode: "development",
  devtool: 'source-map',
  optimization: {
   usedExports: true,
  },
  node: {
    // prevent webpack from injecting eval / new Function through global polyfill
    global: false,
  },
};
