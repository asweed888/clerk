/* <location: greet.english />

hello

*/
package greet

type EnglishMod struct {}
var English = &EnglishMod{}


func (s *EnglishMod) Print() {
    println("this is clerk's default return value")
}


func (s *EnglishMod) Foo() {
    println("this is clerk's default return value")
}
