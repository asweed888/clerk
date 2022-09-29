from . import (
    LimitPageNumber,
)


def clerk(location):
    match location:
        case "LimitPageNumber": return LimitPageNumber.clerk
