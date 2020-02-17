import React, {useEffect} from 'react';
import {connect, useDispatch} from 'react-redux'
import {Panel, Row, Col} from 'muicss/react';
import {findAllNamespaces} from './namespace-repository';

const NamespaceList = (props) => {
    const dispatch = useDispatch();

    const namespaces = props.namespaces.map((namespace, index) => {
        return <Panel key={index}>{namespace}</Panel>
    });

    useEffect(() => {
        dispatch(findAllNamespaces());
    }, []);

    return (
        <Row>
            {namespaces.map((namespaceEl, index) => {
                return (
                    <Col md={4} key={index}>
                        {namespaceEl}
                    </Col>
                );
            })}
        </Row>
    );
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