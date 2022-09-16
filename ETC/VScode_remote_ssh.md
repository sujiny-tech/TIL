# VScode Remote-SSH 접속
> VScode에서 원격접속하여 vscode 에디터로 코드 수정 할 수 있음

+ 소스를 업데이트해야 할 때, 곧바로 서버에서 코드를 수정하기보단 일일히 로컬에서 작업했었음
  > 지금까지는 로컬에서 소스 수정>동작 테스트>백업한 다음, 서버에 업데이트를 했는데,  
  > 이 경우가 아닌, 서버 외부접속을 막고 내부에서 테스트해야될 경우가 있어서 원격접속이 필요했음 ❗  
  > vscode REST client로 테스트할 거여서 vscode Remote-SSH 접속을 처음해봤는데, 조금 헤맨 부분이 있어서 기록📝 

### 방법
1. vscode extenstion: Remote Development 설치
2. F1을 눌러 'Remote-SSH: Open SSH Configuration File...' 선택
3. 보통 'C:\Users\<계정명>\.ssh\'으로 지정하며, 해당 경로에 config 파일 생성
4. config 파일에 HOST, HostName, User, Port 등 기록하고 저장
   + HOST : 해당 서버 이름/별명 지정해주면 됨
   + HostName : 원격접속하려는 서버 주소(AWS의 경우, PublicIP)
   + User : 서버 계정
   + Port : 기본적으로 22로 잡는다고 함, 이외의 포트로 접속하려는 경우 작성 필요
   + IdentityFile : AWS의 경우, 서버 키(pem key) 경로
5. F1을 눌러 'Remote-SSH: Connect to Host' 선택
   + config 파일에서 지정한 HOST명으로 뜰텐데, 원격접속하고자하는 HOST 선택
   + 새창으로 VSCODE창이 띄워지면서, 원격 호스트이 OS선택하라고 나옴
   + 해당하는 OS로 선택하면, 연결 끝.
 
