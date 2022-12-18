# EssayGenie
![DALL·E 2022-12-06 23 45 40 - An oil painting art of genie that writes quality essay](https://user-images.githubusercontent.com/41055141/206066998-3eb39d54-4ae6-4735-ab99-af78360bbf48.png)
* A simple tool to make an automatized essay in English regarding the given Korean keyword in a minute 🇰🇷🇺🇸
## Frontend
* KakaoTalk Plus Friend
## Backend
* Go-chi, [Colly](https://github.com/gocolly/colly), Cobra, Vyper
* OpenGPT API, Papago API, and Quillbot API
## Usage
```
> 한국 사회의 저출생 고령화 문제에 대해 200자 이내로 작성해주세요.
> Korea is facing a major demographic challenge in the form of a low birth rate and aging population. This problem has been developing for several decades, and is now reaching a critical point. The consequences of this demographic shift are significant and far-reaching, and will require careful planning and effective policy intervention to address.
```
## How to start
```
$ go run main.go serve
$ docker build --no-cache -t eg:0.9 .
$ sudo docker run -itd -p 8080:8080 eg:0.9
```
## Heads Up!
* [Not to use localhost but to use 0.0.0.0 as gulf address](https://stackoverflow.com/questions/72978008/accessing-docker-ports-from-browser-localhost8000-doesnt-work)
