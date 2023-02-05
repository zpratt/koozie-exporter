import React from 'react';
import {applyMiddleware, createStore} from "redux";
import topokube from "./reducers";
import thunk from "redux-thunk";
import NamespaceList from "./namespaces/namespace-list";
import PodListContainer from "./pods/pod-list";
import {Provider} from "react-redux";

const store = createStore(topokube, applyMiddleware(thunk));

const AppContainer = () => {
    return (
        <Provider store={store}>
            <div fluid={true}>
                <NamespaceList/>
                <PodListContainer namespace={'topokube'}/>
            </div>
        </Provider>
    );
};

export default AppContainer;
