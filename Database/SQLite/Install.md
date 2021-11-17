# SQLite 설치 및 SQLite Browser 설치
> lightweight in-memory DB

## SQLite 설치 (Windows 10)
+ [SQLite Windows 다운로드](https://www.sqlite.org/download.html) 링크에서 'sqlite-tools-win32-x86-XXX.zip' 파일 다운로드   

+ 해당 파일 압축풀면 3가지 .exe 파일 존재
+ **sqlite3.exe 파일을 실행해서 데이터베이스 구동**
  + cmd 창을 열고, 설치한 디렉터리(즉, 해당 파일이 있는 곳)로 이동   

  + [sqlite 튜토리얼에서 제공해주는 테스트 db](https://www.sqlitetutorial.net/sqlite-sample-database/)가 작동하는지 확인
     + 작동 전, 해당 링크를 통해 'chinook.db' 파일 다운로드   

     + `sqlite3.exe chinook.db` 입력
        > <img src="https://user-images.githubusercontent.com/72974863/142098607-7c200ad1-9cd0-4447-97de-72ba3e210136.png">   
     
     + `.table` 입력을 통해 db내에 테이블 확인
        > <img src="https://user-images.githubusercontent.com/72974863/142099103-fbae851f-87f0-44c9-8590-f68fe9a577c2.png">   

     + `.quit` 입력을 통해 종료   
        > <img src="https://user-images.githubusercontent.com/72974863/142100714-2941507c-ea53-47af-bb73-2b43eef4a0f4.png">   



## DB Browser for SQLite 설치  

+ **...[TODO] install SQLite Browser...**   




## 참고 💫
+ [SQLite inmemorydb](https://www.sqlite.org/inmemorydb.html)
+ [SQLite 인메모리 데이터베이스](https://runebook.dev/ko/docs/sqlite/inmemorydb)
+ [SQLite tutorial](https://www.sqlitetutorial.net/)
+ [DB Browser for SQLite 다운로드](https://sqlitebrowser.org/)
