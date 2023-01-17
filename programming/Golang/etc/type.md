# type 변환
> 형변환 관련된 내용 정리✍
   
   
### ✔️[]byte → big.Int
``` 
result := new(big.Int) // big.Int형 변수 선언
result.SetBytes(before_data) // []byte 타입인 before_data 값을 big.Int로 변환
```

### ✔️big.Int → []byte
```
convert_bytes := result.Bytes() //big.Int인 result에 대해 Bytes로 변환
```

### ✔️string → Int
```
str2int := strcov.Atoi(Age) //string인 값 int로 변환
```

### ✔️string → Int
```
int2str := strcov.Itoa(Age) //int인 값 strnig으로 변환
```

### ✔️Type Assertion VS Type Conversion
+ Type Assertion
   + 변환해주는 것이라기 보단, 해당하는 타입의 값을 확인하는 것을 의미함
   + 다시말해, 인터페이스가 갖고있는 실제 값을 확인하는 것
  
+ Type Conversion
   + 말그대로, 타입을 변환해주는 casting과 같은 의미임
   + go에서는 다른 타입의 변수에 값을 할당해주려면, 명시적으로 형변환을 해줘야 함
   ```
   var tall float64 = 167.5
   var tall_upper uint = uint(tall)
   ```

    

... ing   

#### 참고
 + [golang math/big package](https://pkg.go.dev/math/big)
 + [golang strcov package](https://pkg.go.dev/strconv)   
 
