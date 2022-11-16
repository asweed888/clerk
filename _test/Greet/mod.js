import _hello from "./hello.js"


export default { clerk }
function clerk(location) {
    switch (location) {
        case "hello": return _hello.clerk
    }
}