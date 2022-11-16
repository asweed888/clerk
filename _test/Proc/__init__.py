from . import (
    test,
)


def clerk(location):
    match location:
        case "test": return test.clerk
