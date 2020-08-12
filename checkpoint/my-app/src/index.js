import React from 'react';
import ReactDOM from 'react-dom';
import './index.css';
import Applications from './Applications/Applications';
import Home from './Home/Home';
import Nav from './nav';
import * as serviceWorker from './serviceWorker';
import allReducer from './reducers';
import { createStore } from 'redux';
import {Provider} from 'react-redux';
import {BrowserRouter as Router, Switch, Route} from 'react-router-dom';

const store = createStore(
  allReducer,
  window.__REDUX_DEVTOOLS_EXTENSION__ && window.__REDUX_DEVTOOLS_EXTENSION__());

ReactDOM.render(
  <Provider store={store}>
    <React.StrictMode>
      <Router>
        <Nav />
        <Switch>
          <Route path="/applications" component={Applications}/>
          <Route path="/home" component={Home}/>
        </Switch>
      </Router>
    </React.StrictMode>
  </Provider>,
  document.getElementById('root')
);

// If you want your app to work offline and load faster, you can change
// unregister() to register() below. Note this comes with some pitfalls.
// Learn more about service workers: https://bit.ly/CRA-PWA
serviceWorker.unregister();
