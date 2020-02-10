import React from 'react';
import ReactDOM from 'react-dom'
import {createStore} from 'redux';
import NamespaceListContainer from "./src/namespaces/namespace-list";
import {Provider} from 'react-redux'
import {applyMiddleware} from 'redux';
import thunk from 'redux-thunk'
import topokube from './src/reducers';

const mainEl = document.getElementsByTagName('main')[0];
const store = createStore(topokube, applyMiddleware(thunk));

ReactDOM.render(<Provider store={store}><NamespaceListContainer/></Provider>, mainEl);