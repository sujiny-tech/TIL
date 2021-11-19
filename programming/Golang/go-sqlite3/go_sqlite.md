# go sqlite3 사용

+ 
+ TODO...

## 👉 간단한 예제 코드 : [SQLite_example.go](https://github.com/sujiny-tech/TIL/blob/main/programming/Golang/go-sqlite3/SQLite_example.go) 분석
> 아래 참고 중 'go sqlite example' 파일을 기반으로 변형해봄

## 관련해서 발생한 에러들 🖥
+ `UNIQUE constraint failed : User.Id`
   + 이미 저장되어있는 데이터 값이 중복되는 경우 발생 ❗
   + 어떤 컬럼에 UNIQUE 제약조건을 설정하면, 해당하는 컬럼에 중복된 값이 저장될 수 없음.

## 참고 💫
+ [go sqlite example](https://www.codeproject.com/Articles/5261771/Golang-SQLite-Simple-Example)
+ [go-sqlite3 github](https://github.com/mattn/go-sqlite3)   
+ 
