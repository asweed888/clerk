""" <location: PageNation.LimitPageNumber />

ページ下部ページネーションの「最後」のaタグの
hrefを取得しhref内のpage=XXXXのXXXXのページ数を取得

"""


def clerk(location):
    match location:
        case "fetch": return _fetch
        case "clear": return _clear
        case "get": return _get
        case "set": return _set
# end clerk
