## Bayarcoek
Encrypt project anda dengan aman untuk memberi kemanan dalam transaksi client.
Terinspirasi secara penuh oleh [NodeJS Bayarcoek](https://github.com/tfkhdyt/bayarcoek)

## Disclaimer
Saya tidak bertanggung jawab atas kehilangan data anda jika terjadi ketidak sengajaan
yang disebabkan oleh penggunaan script ini.

## Installation
#### Global Run
Setup [GOPATH](https://dasarpemrogramangolang.novalagung.com/A-gopath-dan-workspace.html) terlebih dahulu
> go install github.com/itzngga/Bayarcoek@latest

#### Once Run
> go run github.com/itzngga/Bayarcoek@latest

## Usage
1. Masuk ke folder target
> cd path/to/project

2. Run Command
- Global Run
```bash 
# encrypt
bayarcoek encrypt

# decrypt
bayarcoek decrypt
```
- Once Run
```bash 
# encrypt
go run github.com/itzngga/Bayarcoek@latest encrypt

# decrypt
go run github.com/itzngga/Bayarcoek@latest decrypt
```

## Example
### Info
- file *main.go* bisa diganti oleh directory dan bisa lebih dari satu file/directory
- default path private key adalah *keys/private.key*
- default path public key adalah *keys/public.key*
- default ekstensi adalah *bayarcoek*
- public key harus ditambahkan jika ingin generate dengan url

### Encrypt dengan private key path
> bayarcoek encrypt --privPath keys/private.key main.go

### Encrypt dengan private key string
> bayarcoek encrypt --privateKey ###BEGINRSAPRIVATEKEYSAMPLE main.go

### Encrypt dengan custom ekstensi
> bayarcoek encrypt --privPath keys/private.key -e bayarcok main.go

### Encrypt dengan generate url public key
> bayarcoek encrypt --privPath keys/private.key --pubPath keys/public.key -u main.go

### Decrypt dengan public key path
> bayarcoek decrypt --pubPath keys/public.key main.bayarcoek

### Decrypt dengan public key string
> bayarcoek decrypt --publicKey ###BEGINRSAPUBLICKEYSAMPLE main.bayarcoek

### Decrypt dengan custom ekstensi
> bayarcoek decrypt --pubPath keys/public.key -e bayarcok main.bayarcok

### Decrypt dengan url yang digenerate dari encrypt
> bayarcoek decrypt -u https://anonfiles.com/sampleURL main.bayarcoek

### Generate RSA Private dan Public Key
> bayarcoek key --priv keys/private.key --pub keys/public.key

### Generate RSA Private dan Public Key Otomatis
> bayarcoek key

# License
[GNU](https://github.com/ItzNgga/Bayarcoek/blob/master/LICENSE)

# Bugs
Please submit an issue if necessary

# Contribute
Pull Request are pleased to