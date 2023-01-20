# BoltDB
> key/value 기반의 DB로 바이트 배열로 값을 넣고, 빠르고 효율적이라 함. 해당 패키지의 함수들을 간단히 정리하자 ✏

+ `func bolt.Open(path string, mode os.FileMode, options *bolt.Options) (*bolt.DB, error)` : 해당 경로에 있는 Boltdb를 open 하는 함수 
   + input  : 경로, 파일모드, 옵션들   
      > 파일모드는 유닉스/리눅스 파일권한 참고
    
   + output : db, 에러발생시 에러코드(없으면 nil)   


+ `func (*bolt.DB).Update(fn func(*bolt.Tx) error) error` : transaction을 read-write하기 위해 사용하는 함수   

+ `func (*bolt.DB).View(fn func(*bolt.Tx) error) error` : transaction을 read 하기 위해(read-only) 사용하는 함수
+ db.Batch()
+ db.Begin()   

+ `func (*bolt.Bucket).Get(key []byte) []byte` : 버킷의 key 값을 검색하는 함수
   + input  : 검색하고 싶은 버킷의 key   

   + output : 입력받은 키에 대해 해당하는 value를 바이트 슬라이스로 반환 (없으면 nil)
+ 
+ 
+ ing...

## 참고
+ [BoltDB github ✨](https://github.com/boltdb/bolt)   

