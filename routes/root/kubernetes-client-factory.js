const kubernetes = require('@kubernetes/client-node');

function buildClient() {
    const config = new kubernetes.KubeConfig();
    config.loadFromCluster();

    return config.makeApiClient(kubernetes.CoreV1Api);
}

module.exports = {
    create: buildClient
};