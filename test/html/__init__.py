from . import (
    RaceSearchResultPage,
    RaceListPage,
)


def clerk(location):
    match location:
        case "RaceSearchResultPage": return RaceSearchResultPage.clerk
        case "RaceListPage": return RaceListPage.clerk
