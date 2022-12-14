/* <location: ps.printEnglishGreet />



*/
package ps

type PrintEnglishGreetMod struct {}
var PrintEnglishGreet = &PrintEnglishGreetMod{}


func (s *PrintEnglishGreetMod) Exec() {
    println("this is clerk's default return value")
}
