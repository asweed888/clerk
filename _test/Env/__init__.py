from . import (
    var,
)


def clerk(location):
    match location:
        case "var": return var.clerk
