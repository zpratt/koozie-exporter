const {create} = require('./kubernetes-client-factory');

module.exports = {
    findAllPodsByNamespace: async (namespace) => {
        try {
            const client = create();
            const result = await client.listNamespacedPod(namespace);

            return result.body.items.map((pod) => {
                const containers = pod.spec.containers.map((container) => {
                    return {
                        name: container.name,
                        image: container.image
                    };
                });

                return {
                    name: pod.name,
                    containers
                };
            });
        }
        catch(err) {
            return err;
        }
    }
};