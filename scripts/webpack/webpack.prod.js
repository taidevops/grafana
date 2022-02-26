'use strict';

const { merge } = require('webpack-merge');
const common = require('./webpack.common.js');

module.exports = (env = {}) =>
  merge(common, {
    mode: 'production',
    devtool: 'source-map',

    module: {
      // Note: order is bottom-to-top and/or right-to-left
      rules: [
        {
          test: /\.tsx?$/,
          exclude: /node_modules/,
          use: [
            {
              loader: 'babel-loader',
              options: {
                cacheDirectory: true,
                cacheCompression: false,
              },
            },
          ],
        },
      ],
    },

    // enable persistent cache for faster builds
    cache: {
      type: 'filesystem',
      name: 'grafana-default-production',
      buildDependencies: {
        config: [__filename],
      },
    },
  })
