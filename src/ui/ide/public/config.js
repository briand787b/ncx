// This file is to be overwritten in production environments.
// It acts as an alternative to env vars
function config () {
    return { 
        buildServerAddress: "http://buildsys/api",
    };
}