""" <location: Proc.test />

テスト用のプロシージャ

"""


def clerk(location):
    match location:
        case "run": return _run
# end clerk


def _run():
    print("this is clerk's default return value")
