/* <location: ps.printJapaneseGreet />



*/
package ps

type PrintJapaneseGreetMod struct {}
var PrintJapaneseGreet = &PrintJapaneseGreetMod{}


func (s *PrintJapaneseGreetMod) Exec() {
    println("this is clerk's default return value")
}
