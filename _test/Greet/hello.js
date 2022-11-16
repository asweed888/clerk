/* <location: Greet.hello />

hello world

*/

export default { clerk }
function clerk(location) {
    switch (location) {
        case "get": return _get
        case "print": return _print
    }
}
// end clerk


function _get(){
    return "Hello world"
}


function _print(){
    console.log("Hello World")
}
