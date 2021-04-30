const path = require('path');
const webpack = require('webpack');
const TerserPlugin = require('terser-webpack-plugin');
const MiniCssExtractPlugin = require('mini-css-extract-plugin');
const CssMinimizerPlugin = require('css-minimizer-webpack-plugin');

module.exports = {
  entry: './src/main.js',
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
      use: [MiniCssExtractPlugin.loader, "css-loader"],
    }]
  },
  mode: "production",
  optimization: {
   usedExports: true,
   minimizer: [new TerserPlugin(), new CssMinimizerPlugin()],
  },
  node: {
    // prevent webpack from injecting eval / new Function through global polyfill
    global: false,
  },
};
