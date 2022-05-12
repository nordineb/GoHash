# Go Hash 

Go program that produces SHA-512 hashes of all the values included in the file `values.txt`, The result can be piped into a file for further processing
 

## Alternative with `openssl`
```
❯ echo -n 12345678900 | openssl dgst -binary -sha512 | base64

WB0Tb5keIE7xXYlAZglchX5XMYo2dvbYzSSugG83vHNX7Ls67Lrs6ukheOS2uDFqimqBKCkuusDB3XDiv/YCgg==
````

## Running 
```
❯ go run main.go values.txt
26117037664 - 22oJRMjUM8kG89jmlj9Z0BRZydpK62i/Rx1bRvyD/ZEB/2ZST7Qzcb/JY7VfLlJ38RLGKORKQeursCfrt3aakQ== 
06087145172 - KV4NjffuRVseHumkA0TGK1aC1clKyaVjEETWQ3jaTPHmdLRTrSEIPwp1MAsMr80q3ekiTDDGv2/q86eUnTLrBw== 
26117004448 - eR8h3FjHZk28Gn+VypSNCwso4l1gt7k2onmhxNvfKWP8g0IAbtoKdVqWTM6WptwDKdWvuPvbJzHqhoa7zmgtzg== 
15027135606 - ae9PKOoWyintsJTT206m75S68hEnay6CAem3Vt/EjZMQAaQoTWK0I/6ZUex1ji5iCe07L18ILH9+MDa2HAtQQg== 
12057221524 - M1cykM9x440jjjoi4+GyUVZxANkyBvt+eiZUKKrXimNF4osoR13B7+hndrj98V4G2LhXVItzKQlsOgCtFFanAA== 
04087113790 - aPiJsB+cvsgErSTf9VPLHnxqVXjP77l9tma5IjAnepx8IgTXEwd2bc5BhJZppYEPnr2oj1pBf94Dgs3pajun/A== 
04067040126 - HIQeoykswcDDJuczB9+ziBx0Cav4+LdQz+ysRYgVRbm4VYhRyLPUNQFG+r+SARxtwh4tSW0z4848uRDVg6TM3g== 
05127013898 - PBU3qqJJ9viy6/xjb8VUDNDq1SHWzgFV2yXgy2PHaGLRySlHNSiasDr3FBwCUEQ0TFaENiEd4VRJR/iEuaUvbA== 
```