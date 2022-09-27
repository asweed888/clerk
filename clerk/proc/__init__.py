import hoge.clerk.proc.run


def Clerk(location):
    match location:
        case "run": return hoge.clerk.proc.run.Clerk
