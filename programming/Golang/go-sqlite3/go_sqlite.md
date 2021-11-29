# go sqlite3 사용


## 👉 간단한 예제 코드 : [SQLite_example.go](https://github.com/sujiny-tech/TIL/blob/main/programming/Golang/go-sqlite3/SQLite_example.go) 분석
> 아래 참고 중 'go sqlite example' 파일을 기반으로 변형해봄

+ 예제코드는 sqlite db 파일을 생성하고, 열고, 테이블을 생성한 다음, 데이터를 삽입하고, 출력하는 루틴임   

+ 예제코드에서 사용된 함수 정리 🖥
   + `os.Create("원하는 파일 이름.db")` : 원하는 파일이름으로 db를 생성함
   + `sql.Open("sqlite3", "./원하는 파일 이름.db")` : 열고 싶은 파일의 경로를 파라미터에 입력하여, 해당 파일 열음
   + `createTable` 함수 : 입력받은 db에 테이블을 생성하는 함수
   + `insertUser` 함수 : db, index, name, password(생성한 테이블에 맞는 데이터값)을 입력받아 삽입하는 함수
   + `displayUser` 함수 : 입력받은 db의 User 테이블에 있는 데이터를 출력해주는 함수

## 관련해서 발생한 에러들 🖥
+ `UNIQUE constraint failed : User.Id`
   + 이미 저장되어있는 데이터 값이 중복되는 경우 발생 ❗
   + 어떤 컬럼에 UNIQUE 제약조건을 설정하면, 해당하는 컬럼에 중복된 값이 저장될 수 없음.

+ `no such table : <table_name>`
   + 해당 db에 <table_name>으로 된 테이블이 존재하지 않아서 발생하는 문제
   + 나같은 경우, sql로 db 생성하고 저장하지 않은 상태에서 db를 사용해서 에러가 발생하였음
   + 또는 해당 테이블 이름에 오타가 없는지 체크 필요


## 참고 💫
+ [go sqlite example](https://www.codeproject.com/Articles/5261771/Golang-SQLite-Simple-Example)
+ [go-sqlite3 github](https://github.com/mattn/go-sqlite3)   

