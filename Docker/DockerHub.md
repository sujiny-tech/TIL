# DockerHub 사용법
> 사용을 위해서는 DockerHub 가입필요(회원가입->이메일 확인)    

### 쉘에서 DockerHub 로그인
+ `docker login` 입력 후, Username/Password 입력하여 로그인
+ `docker logout` 입력하면, 해당 로그인된 계정 로그아웃

### Repository 생성
+ DockerHub에서 로그인 후, Repository 생성 가능
   + 공개유무에 따라 public, prviate 선택하여 생성 가능함
+ Repository 생성하면, 해당 repository에 이미지를 업로드할지 알려주는 command가 뜸
   + `docker push <Username/repositoryName>:tagname`

### 이미지 업로드
+ 업로드하기 위한 DockerHub에서 repository 생성 필요함
+ Repository에 image의 tag에 따라 업로드할 수 있음
+ 이전에 생성한 이미지의 이름에 특정 tag를 부여하는 방법은 다음 아래와 같음
   + `docker tag <already_created_image_name:tag> <Username/repositoryName>:tagname`
   + `docker images`를 통해 특정된 tag가 부여되었는지 확인
+ `docker push <Username/repositoryName>:tagname`을 통해 특정 tag의 이미지 업로드

### 이미지 다운로드
+  `docker pull <Username/repository>:tagname`을 통해 특정 tag의 이미지 다운로드
