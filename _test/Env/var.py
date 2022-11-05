""" <location: Env.var />

環境変数

"""


def clerk(location):
    match location:
        case "get": return _get
# end clerk
