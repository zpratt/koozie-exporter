module.exports = {
    name: 'health',
    register: (server) => {
        server.route({
            method: 'GET',
            path: '/health',
            handler: async () => {
                return {
                    status: 'OK'
                };
            }
        });
    }
};
