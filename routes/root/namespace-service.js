const {create} = require('./kubernetes-client-factory');

module.exports = {
    findAllNamespaces: async () => {
        try {
            const client = create();
            const result = await client.listNamespace();

            const namespaceNames = result.response.body.items.map((item) => {
                return item.metadata.name
            });

            return {
                namespaceNames
            };
        }
        catch(err) {
            return err;
        }
    }
};