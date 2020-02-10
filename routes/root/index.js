const {findAllNamespaces} = require('./namespace-service');
const {findAllPodsByNamespace} = require('./pod-service');

module.exports = {
    name: 'root',
    register: (server) => {
        server.route({
            method: 'GET',
            path: '/api',
            handler: async (req) => {
                return findAllNamespaces();

            }
        });

        server.route({
            method: 'GET',
            path: '/api/pods/{namespace}',
            handler: async (req) => {
                const {namespace} = req.params;
                return findAllPodsByNamespace(namespace);
            }
        });
    }
};
