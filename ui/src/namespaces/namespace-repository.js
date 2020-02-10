export const findAllNamespaces = () => {
    return (dispatch) => {
        const namespaceUrl = 'http://www.mocky.io/v2/5e40d7162f00005500583077';
        // const namespaceUrl = '/api';

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
