from kazoo.client import KazooClient

zk = KazooClient(hosts='127.0.0.1:2181,127.0.0.1:2182,127.0.0.1:2183')
zk.start()

from kazoo.client import KazooState

def my_listener(state):
    if state == KazooState.LOST:
        print('# Register somewhere that the session was lost')
    elif state == KazooState.SUSPENDED:
        print('# Handle being disconnected from Zookeeper')
    else:
        print('# Handle being connected/reconnected to Zookeeper')

zk.add_listener(my_listener)

def watcher(event):
    print('-'*24, '>', event)

def dataW(*args):
    print('='*24, '>', args)

children = zk.get_children('/my/', watch=watcher)
data = zk.get('/my/college', watch=dataW)
print(children, data)

zk.set("/my/college", b'FGIET RAE BARELI')
ans = zk.get("/my/college")
print(ans)
zk.stop()
