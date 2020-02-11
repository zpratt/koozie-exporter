import React, {useEffect} from 'react';
import {connect, useDispatch} from 'react-redux'
import {findPodsByNamespace} from './pod-repository';

const PodList = (props) => {
    const dispatch = useDispatch();

    const podsListElements = props.pods.map((pod, index) => {
        return <li key={index}>{pod.containers.map((container) => `${container.name}, `)}</li>
    });

    useEffect(() => {
        dispatch(findPodsByNamespace(props.namespace));
    }, []);
    return <ul>{podsListElements}</ul>;
};

const mapStateToProps = (state) => {
    return {
        pods: state.app.pods || []
    };
};

const mapDispatchToProps = () => {
    return {};
};

const PodListContainer = connect(mapStateToProps, mapDispatchToProps)(PodList);

export default PodListContainer;