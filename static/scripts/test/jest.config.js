/** @type {import('jest').Config} */
export default {
    testEnvironment: 'jsdom',
    moduleFileExtensions: ['js'],
    rootDir: './',
    modulePaths: ['<rootDir>/../'],
    moduleDirectories: ['node_modules', '<rootDir>/../'],
};
