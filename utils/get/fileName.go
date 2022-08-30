package get


func ModuleMainFileName(lang string) string {
    switch lang {
        case "python": return "__init__.py" 
        default: return ""
    }
}
