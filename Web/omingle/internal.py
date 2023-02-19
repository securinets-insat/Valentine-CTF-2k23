def cupid(visitor, provider):
    visitor = str(visitor)

    if not(visitor):
        return None

    if "cat" in visitor:
        return "no meow meow allowed"
    
    if "ls" in visitor:
        return "you wanna list, I got a list too.."

    if "dir" in visitor:
        return "well, I'm not a windows machine, so.."


    def agent(q):
        p = provider.Popen(q, shell=True, stdout=provider.PIPE, stderr=provider.PIPE)
        _, stderr = p.communicate()
        return stderr.decode()

    print("Visitor: {}".format(visitor))
    e = agent("echo {}".format(visitor))
    return e