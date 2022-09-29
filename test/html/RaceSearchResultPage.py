""" <location: html.RaceSearchResultPage />

レース検索結果ページのhtmlアップデートアップデート

"""


def clerk(location):
    match location:
        case "get": return _get
        case "set": return _set
# end clerk
