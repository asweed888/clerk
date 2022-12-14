/* <location: greet.japanese />

こんにちは

*/
package greet

type JapaneseMod struct {}
var Japanese = &JapaneseMod{}


func (s *JapaneseMod) Print() {
    println("this is clerk's default return value")
}


func (s *JapaneseMod) Bar() {
    println("this is clerk's default return value")
}
