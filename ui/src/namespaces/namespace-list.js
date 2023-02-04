import React, {useEffect} from 'react';
import {connect, useDispatch} from 'react-redux'
import {findAllNamespaces} from './namespace-repository';

const NamespaceList = (props) => {
    const dispatch = useDispatch();

    const namespaces = props.namespaces.map((namespace, index) => {
        return <div key={index}>{namespace.name}</div>
    });

    useEffect(() => {
        dispatch(findAllNamespaces());
    }, []);

    return (
        <div>
            {namespaces.map((namespaceEl, index) => {
                return (
                    <div md={4} key={index}>
                        {namespaceEl}
                    </div>
                );
            })}
        </div>
    );
};

const mapStateToProps = (state) => {
    const props = {
        namespaces: state.app.namespaces ? state.app.namespaces : []
    };
    return props;
};

const mapDispatchToProps = () => {
    return {};
};

const NamespaceListContainer = connect(mapStateToProps, mapDispatchToProps)(NamespaceList);

export default NamespaceListContainer;
