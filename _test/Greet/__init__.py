from . import (
    hello,
)


def clerk(location):
    match location:
        case "hello": return hello.clerk
