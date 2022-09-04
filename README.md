## Url-shortener

A services that will help you shorten your links

## Installation


```sh
git clone https://github.com/VrMolodyakov/url-shortener
make start
```

## Unit tes

```sh
make test 
```

| package | Coverage |
| ------ | ------ |
| urlDb  |  100.0% of statements |
| service | 97.9% of statements |
| handler  | 83.3% of statements |


## Api

```
http://localhost:8080/encode
```
####To encode url

------------


####Request Body :
 - url - your url to shorten

Example :

```
{
    "url": "https://github.com/VrMolodyakov/url-shortener"
}

```
Response :

```
{
   "shortUrl": "http://localhost:8080/bRTcI5raimH",
   "fulltUrl": "https://github.com/VrMolodyakov/url-shortener"
}
```

------------


```
http://localhost:8080/encode
```
####To shorten with a custom link

------------


####Request Body :
 - url - your url to shorten
 - custom - custo url

Example :

```
{
    "url": "https://github.com/VrMolodyakov/url-shortener",
    "custom":"shortener"
}

```
Response :

```
{
   "shortUrl": "http://localhost:8080/shortener",
   "fulltUrl": "https://github.com/VrMolodyakov/url-shortener"
}
```
