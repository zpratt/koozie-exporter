export const findPodsByNamespace = (namespace) => {
    return (dispatch) => {
        const namespaceUrl = `/api/pods/${namespace}`;

        fetch(namespaceUrl)
            .then((result) => {
                return result.json()
            })
            .then((response) => {
                dispatch({
                    type: 'pods',
                    data: response
                });
            });
    };
};
