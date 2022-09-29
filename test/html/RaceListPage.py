""" <location: html.RaceListPage />

レース一覧ページのhtml

"""


def clerk(location):
    match location:
        case "open": return _open
        case "get": return _get
        case "close": return _close
# end clerk
