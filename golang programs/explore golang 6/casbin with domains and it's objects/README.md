check normal condition
check object and action wild card
check admin 
check object wildcarded but action not

p, owner, server1, *, *
p, owner, server2, *, *
p, admin, *, *, *

p, player, server1, *, read
p, guest, server1, startup, read
p, guest, server1, about, read
p, game-master, server1, game-setting, *

g, alice, owner, server1
g, bob, owner, server2
g, ketan, admin, *

g, aman, player, server1
g, tridip, guest, server1
g, vatsal, game-master, server1