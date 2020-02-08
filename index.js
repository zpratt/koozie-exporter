const Glue = require('@hapi/glue');

const manifest = {
    server: {
        port: process.env.PORT,
        host: `${process.env.HOST}`
    },
    register: {
        plugins: [
            {plugin: './routes/root'},
            {plugin: './routes/health'}
        ]
    }
};

const options = {
    relativeTo: __dirname
};

const init = async () => {

    const server = await Glue.compose(manifest, options);

    await server.start();
    console.log('Server running on %s', server.info.uri);
};

process.on('unhandledRejection', (err) => {
    console.log(err);
    process.exit(1);
});

init();