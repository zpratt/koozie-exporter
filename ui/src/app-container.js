import {Container} from 'muicss/react';
import React from 'react';
import {applyMiddleware, createStore} from "redux";
import topokube from "./reducers";
import thunk from "redux-thunk";
import NamespaceListContainer from "./namespaces/namespace-list";
import PodListContainer from "./pods/pod-list";
import {Provider} from "react-redux";

const store = createStore(topokube, applyMiddleware(thunk));

const AppContainer = () => {
    return (
        <Provider store={store}>
            <Container fluid={true}>
                <NamespaceListContainer/>
                <PodListContainer namespace={'topokube'}/>
            </Container>
        </Provider>
    );
};

export default AppContainer;