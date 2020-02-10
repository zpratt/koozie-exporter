import React, {useEffect} from 'react';
import {connect, useDispatch} from 'react-redux'
import {findAllNamespaces} from './namespace-repository';

const NamespaceList = (props) => {
    const dispatch = useDispatch();

    const namespaces = props.namespaces.map((namespace, index) => {
        return <li key={index}>{namespace}</li>
    });

    useEffect(() => {
        dispatch(findAllNamespaces());
    }, []);
    return <ul>{namespaces}</ul>;
};

const mapStateToProps = (state) => {
    return {
        namespaces: state.app.namespaces ? state.app.namespaces : []
    };
};

const mapDispatchToProps = () => {
    return {};
};

const NamespaceListContainer = connect(mapStateToProps, mapDispatchToProps)(NamespaceList);

export default NamespaceListContainer;