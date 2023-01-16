# type 변환
> 형변환 관련된 내용 정리✍

### []byte → big.Int
``` 
result := new(big.Int) // big.Int형 변수 선언
result.SetBytes(before_data) // []byte 타입인 before_data 값을 big.Int로 변환
```

... ing
