import {combineReducers} from 'redux';

const app = (state = [], action) => {
    const actions = {
        'namespaces': {
            namespaces: action.data
        }
    };

    return actions[action.type] ? {
        ...state,
        ...actions[action.type]
    }: {...state};
};

export default combineReducers({app});