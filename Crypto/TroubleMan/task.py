# TroubleMan
# @Ou_sama

from secret import flag,params,func
from Crypto.Util.number import getRandomNBitInteger
flag = [ord(x) for x in list(flag)]

rand = getRandomNBitInteger(100)
test = (296*func(rand)-288)/(2*func(rand)+248)
assert(abs(test-func(rand+1))<10**-4)

cipher = []

for i in flag :
    cipher.append(func(i))

with open("out.txt","w") as f:
    f.write(f'func(0) = -3890/45\n')
    f.write("Cipher = [")
    for i in cipher :
        f.write(str(i)+",")
    f.write("]\n")
f.close()