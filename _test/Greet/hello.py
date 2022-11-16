""" <location: Greet.hello />

hello world

"""


def clerk(location):
    match location:
        case "get": return _get
# end clerk


def _get():
    print("this is clerk's default return value")


def _set():
    print("this is clerk's default return value")
