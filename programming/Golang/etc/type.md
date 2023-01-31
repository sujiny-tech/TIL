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
### ✔️ Explicit type conversion VS Implicit type conversion
+ Java에서는 명시적 타입 변환(Explicit type conversion), 암시적 타입 변환(Implicit type conversion) 둘 다 지원
+ 반면 Go에서는 명시적 타입 변환만 지원해줌
   + 명시적 타입 변환 : 직접 사용자가 데이터 타입을 변경
   + 암시적 타입 변환(또는 묵시적 타입 변환) : 컴파일러에 의해 자동으로 타입 변환이 이뤄지는 것
      + C, Java에서처럼 int 데이터와 float 데이터를 덧셈한다면, 자동으로 float로 타입이 변환되는 경우
      + 표현범위가 좁은 데이터 타입에서 넓은 데이터 타입으로의 변환만 허용됨. (int는 double로 형변환 가능, double은 int로 불가능)
    

... ing   

#### 참고
 + [golang math/big package](https://pkg.go.dev/math/big)
 + [golang strcov package](https://pkg.go.dev/strconv)   
 
