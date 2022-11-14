const _hello = require("./hello")


module.exports.clerk = (location) => {
    switch (location) {
        case "hello": return _hello.clerk
    }
}