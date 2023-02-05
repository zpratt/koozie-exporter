import React, {useEffect} from 'react';
import {useDispatch, useSelector} from 'react-redux'
import {findAllNamespaces} from './namespace-repository';

const NamespaceList = (props) => {
    const dispatch = useDispatch();
    const namespaces = useSelector(state => state.app.namespaces);

    useEffect(() => {
        dispatch(findAllNamespaces());
    }, []);

    return (
        <div>
            {namespaces.map((namespace, index) => {
                return (
                    <div md={4} key={index}>
                        <div>{namespace.name}</div>
                    </div>
                );
            })}
        </div>
    );
};

export default NamespaceList;
