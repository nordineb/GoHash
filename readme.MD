# Go Hash 

Go program that produces SHA-512 hashes of all the values included in the file `values.txt`, The result can be piped into a file for further processing
 

## Alternative with `openssl`
```
❯ echo -n 12345678900 | openssl dgst -binary -sha512 | base64

WB0Tb5keIE7xXYlAZglchX5XMYo2dvbYzSSugG83vHNX7Ls67Lrs6ukheOS2uDFqimqBKCkuusDB3XDiv/YCgg==
````