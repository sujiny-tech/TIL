# build
> 소스를 보다가 익숙치 않은 build 설정 부분에 대해 조사하고, 이해한대로 정리📝


## 제약조건 설정

- build tag, test tag 설정이 가능함
- `// +build` : Go1.17버전 전까지 사용된 태그 뒤에는 빈줄이 와야 함
  ```
  // +build main1

  package main
  ```

- `//go:build` : Go1.17버전부터 `// +build` 대신 bool 표현식을 지원 (&&, ||, !)
  - OR : //go:build ex1 || ex2
    > 이전 버전 : // +build ex1 ex2 (공백으로 구분)
  - AND : //go:build ex1 && ex2
    > 이전 버전 : // +build ex1,ex2 (쉼표로 구분)
  - NOT : //go:build !ex1
    > 이전 버전 : // +build !ex1 (변경되지 않음)
    
- 이전버전에서는 여러 줄의 주석으로 설정해줬어야 하지만, Go1.17부터는 bool 표현식을 통해 한줄로 표현이 가능해짐
  - 이전버전
    ```
    // +build ex1 ex2
    // +build 386
    ```
  - Go1.17 버전~
    ```
    //go:build (ex1 || ex2) && 386
    ```



### 참고
- [go package build](https://pkg.go.dev/go/build)
- [Difference between '//go:build' and '//+build'](https://stackoverflow.com/questions/68360688/whats-the-difference-between-gobuild-and-build-directives?rq=3)
- [hdr-Build_constraints](https://pkg.go.dev/cmd/go#hdr-Build_constraints)
- [Golang conditional compliation](https://stackoverflow.com/questions/10646531/golang-conditional-compilation/67937234#67937234)
