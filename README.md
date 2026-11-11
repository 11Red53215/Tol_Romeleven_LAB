# Tol_Romeleven_LAB

Cryptography homework — Term 1.

- `launcher.go` — menu that runs the labs
- `Week 2/` — lab1.go to lab6.go
- `Week 3/` — hash proof and password-cracking labs
- `utils/crack/` — shared MD5, SHA-1 and SHA-512 cracking helpers
- `Week 3/wordlists/nord_vpn.txt` — password wordlist
- `Week 3/output/` — verbose cracking logs (created when the labs run)

Run:

```
go run launcher.go
```

## Week 3 labs

| File | Lab | Result |
| --- | --- | --- |
| `Week 3/lab0.go` | Lab #0 — hash two inputs with MD5, SHA-1, SHA-256, SHA-512, SHA3-256 and compare | prints Match / No Match |
| `Week 3/lab1.go` | Lab #1 — crack MD5 `6a85dfd77d9cb35770c9dc6728d73d3f` | `Google101` |
| `Week 3/lab2.go` | Lab #2 — crack SHA-1 `aa1c7d931cf140bb35a5a16adeb83a551649c3b9` | `Princess1` |
| `Week 3/lab3.go` | Lab #3 — crack SHA-512 `485f5c...ca5dd38` | `Gutter11` |
| `Week 3/lab4.go` | Lab #4 — mini-CTF (cat leak) | `cryptoCTF{meowmeow}` |

Each cracking lab prints every word/hash pair to the console and also writes
the same verbose output to `Week 3/output/` (`verbose_md5.txt`,
`verbose_sha1.txt`, `verbose_sha512.txt`).
