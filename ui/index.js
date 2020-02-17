import React from 'react';
import ReactDOM from 'react-dom';
import 'muicss/dist/css/mui.css';
import AppContainer from './src/app-container';

const mainEl = document.getElementsByTagName('main')[0];

ReactDOM.render(<AppContainer/>, mainEl);