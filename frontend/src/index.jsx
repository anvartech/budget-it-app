// import bugsnag from 'bugsnag-js';
import bugsnagReact from '@bugsnag/plugin-react';

import React from 'react';
import { hydrate, render } from 'react-dom';
import createPlugin from 'bugsnag-react';

import { Router } from 'react-router-dom';
import { Provider } from 'react-redux';
import { createBrowserHistory } from 'history';

import store from 'utils/store';
import * as serviceWorker from 'serviceWorker';

import App from './App';
import './index.css';

const history = createBrowserHistory();

const bugsnagClient = bugsnag(process.env.BUGSNAG_API);

bugsnagClient.use(bugsnagReact, React)
// const ErrorBoundary = bugsnagClient.use(createPlugin(React));
const ErrorBoundary = bugsnagClient.getPlugin('react');

const rootElement = document.getElementById('root');
const toRender = (
  <ErrorBoundary>
    <Router history={history}>
      <Provider store={store}>
        <App />
      </Provider>
    </Router>
  </ErrorBoundary>
);

if (rootElement.hasChildNodes()) {
  hydrate(toRender, rootElement);
} else {
  render(toRender, rootElement);
}
// ReactDOM.render(
//     ,
//     rootElement
// );

// If you want your app to work offline and load faster, you can change
// unregister() to register() below. Note this comes with some pitfalls.
// Learn more about service workers: http://bit.ly/CRA-PWA
serviceWorker.unregister();
