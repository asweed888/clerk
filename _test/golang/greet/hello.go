/* <location: greet.hello />

hello world

*/
package greet

type hello struct {}


func (s *hello) Print() {
    println("this is clerk's default return value")
}
