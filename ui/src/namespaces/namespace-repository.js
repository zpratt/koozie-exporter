export const findAllNamespaces = () => {
    return (dispatch) => {
        const namespaceUrl = '/api';

        fetch(namespaceUrl)
            .then((result) => {
                return result.json()
            })
            .then((response) => {
                dispatch({
                    type: 'namespaces',
                    data: response.namespaceNames
                });
            });
    };
};
