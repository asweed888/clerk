/* <location: goModFile.moduleName />

go.modファイルの1行目のmodule名

*/
package goModFile

type moduleNameMod struct {}
var ModuleName = &moduleNameMod{}


func (s *moduleNameMod) Get() {
    println("this is clerk's default return value")
}
