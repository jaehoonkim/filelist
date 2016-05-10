filelist
========

## 목적
경로상의 "파일" 또는 "디렉토리"를 리스트화해서 result.txt 파일로 만든다.

## 설치
```
go get github.com/sabzil/filelist
```

## 사용법
```
filelist -path="./" -target=file
path: 리스트화 하고 싶은 파일 또는 디렉토리가 있는 경로(지정하지 않으면 현재의 경로)
target: 리스트화 하고 싶은 타입에 따라서 file 또는 directory 를 지정한다.
```
