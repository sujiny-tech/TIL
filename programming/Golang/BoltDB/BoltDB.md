# BoltDB
> key/value 기반의 DB로 바이트 배열로 값을 넣고, 빠르고 효율적이라 한다. 해당 패키지의 함수들을 간단히 정리하자 ✏

+ `func bolt.Open(path string, mode os.FileMode, options *bolt.Options) (*bolt.DB, error)` : 해당 경로에 있는 Boltdb를 open 하는 함수 
   + input  : 경로, 파일모드, 옵션들   
      > 파일모드는 유닉스/리눅스 파일권한 참고
    
   + output : db, 에러발생시 에러코드(없으면 nil)   


+ db.Update()
+ db.View()
+ db.Batch()
+ db.Begin()   

+ Bucket.get()
+ 
+ 
+ ing...

## 참고
+ [BoltDB github ✨](https://github.com/boltdb/bolt)   

