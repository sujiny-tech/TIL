# 데이터베이스/테이블 생성

### 테이블 생성
> DB Browser for SQLite를 통해 데이터베이스 생성/삭제 가능 


+ DB Browser for SQLite를 실행한 다음, "새 데이터베이스" 클릭
+ 맨 위의 빈칸에 해당 테이블 이름을 작성하고, **필드와 제약조건을 걸 수 있음**
   + 필드 추가 시, 체크박스가 있는 열의 의미
      + **NN**(Not Null): null 값 허용 하지 않는다는 의미
      + **PK**(Primary Key): 기본 키로 설정한다는 의미
      + **AI**(Autoincrement): 자동으로 증가한다는 의미
      + **U**(Unique): 유일한(중복되지 않은) 데이터라는 의미

+ 추가적으로 sqlite는 **boolean 타입이 존재하지 않음** ❗❗
   > 따라서 나는 boolean 타입 대신 integer 타입인 0과 1으로 지정해서 수행하였음

### 데이터베이스 생성
방법 1)
+ 만들어낸 .sql 파일을 DB Browser fo SQLite을 통해 데이터베이스 생성 가능   

+ 왼쪽 "파일" 탭에서 "가져오기" 중에서 SQL 파일로부터 데이터베이스 가져오기 클릭
   > <img src="https://user-images.githubusercontent.com/72974863/143370488-75cd090b-4127-4cea-b43a-4ebd767760ae.png" width="90%" hegith="90%">   

방법 2)
+ CMD 오픈해서, `sqlite3 <생성하려고하는 데이터베이스.db>`
   + sqlite3.exe 위치에 가서 명령어 입력 혹은 환경변수 설정
+ 데이터 베이스 생성 후, 해당 경로에서 데이터베이스 조회(목록 조회)
   > <img width="432" alt="image" src="https://github.com/sujiny-tech/TIL/assets/72974863/abe0eaab-89c9-47ca-a51e-93cb94a55081">

