import 'core-js';

import ReactDOM from 'react-dom';
import React from 'react';

import { AppWrapper } from './AppWrapper';

export class GrafanaApp {
  async init() {
    ReactDOM.render(
      React.createElement(AppWrapper, {
        app: this,
      }),
      document.getElementById('reactRoot')
    );
  }
}

export default new GrafanaApp();