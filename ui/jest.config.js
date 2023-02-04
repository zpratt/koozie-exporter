const nextJest = require('next/jest');

const configBuilder = nextJest({
    dir: './'
});

module.exports = configBuilder({
    moduleDirectories: ['node_modules', 'src'],
    moduleNameMapper: {
        '@/(.*)$': '<rootDir>/src/$1'
    },
    testEnvironment: 'jest-environment-jsdom'
});
