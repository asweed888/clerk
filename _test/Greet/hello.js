/* <location: Greet.hello />

hello world

*/


module.exports.clerk = (location) => {
    switch (location) {
        case "get": return _get
        case "set": return _set
        case "find": return _find
        case "fetch": return _fetch
    }
}
// end clerk


function _get(){
    console.log("this is clerk's default return value")
}


function _set(){
    console.log("this is clerk's default return value")
}


function _find(){
    return "Hello World"
}


function _fetch(){
    console.log("this is clerk's default return value")
}
