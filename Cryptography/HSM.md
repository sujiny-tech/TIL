# HSM(Hardware Security Module)
> 정의, 관련 오픈 소스 분석 진행📝


## 개념
+ 암호키를 안전하게 저장하고, 물리적/논리적으로 보호하는 역할을 수행하는 하드웨어 장치
+ 관련인증 : FIPS 140-2
   + Leve 1
     + 기본적인 보안 요구사항에 대한 평가
     + 물리적 보안성에 대해서는 요구되지 않고, 주로 소프트웨어적 암호화 모듈에 대한 평가임
   + **Level 2**
     + 암호화 모듈 뿐 아니라 하드웨어적 부문에서 보안 적합성 평가 실시
     + 물리적인 장비 해체 방지 등에 대한 대책이 강구되어야 함
   + **Level 3**
     + 제품 분해를 막기 위한 단편적 방편 뿐 아니라 적극적인 대책까지 제시되는 지 평가
     + 무단 접근과 같이 부당한 방법의 변경 또는 오용 감지 시 해당 자치는 주요 보안 변수를 삭제
  + Level 4
     + 물리적인 보호가 어려운 환경을 고려해 평가
     + 클라우드 등 새로운 키 관리 서비스나 제품에 적용하기 위한 레벨
       
## 특징
+ 암호키를 HSM 내부에 저장하여, 안전하게 관리하는 역할을 수행
+ 내부적으로 키를 관리하고, 암호 연산 자체를 내부에서 수행(외부 유출 X)
 
## 관련 오픈 소스
### SoftHSM v2
+ 개념 : PKCS#11 기능을 소프트웨어로 구현한 오픈소스 프로그램
  + PKCS#11
    + Cryptographic Token Interface Standard 로 암호 서비스를 제공하는 표준 API 정의
    + C를 사용하는 ANSI C 함수의 정의
    + 정의된 함수들을 Cryptoki 라는 동적 라이브러리에 만드는 것에 대한 표준 기술임
    
+ 특징
  + OpenDNSSEC 프로젝트의 일부로 오픈소스임
    + PKCS#11 인터페이스를 통해 암호화 키를 처리하고 저장함
    + HSM 시뮬레이터로, 이를 통해 암호화 키 및 보안 관련 작업을 시뮬레이션 할 수 있음
  + 암호화 라이브러리 Botan 또는 OpenSSL에 의존성 존재
    + Botan 2.0.0 : C++로 작성된 오픈소스 라이브러리로, 암호화 및 프로토콜 등을 지원하며 보안 응용 프로그램을 개발하는 데에 사용됨
    + OpenSSL 1.0.0 : C로 작성된 오픈소스 라이브러리로, 암호화/디지털서명 및 SSL/TLS 프로토콜 구현 등 다양한 보안 기능을 제공함
    > Botan이 상대적으로 더 간단한 인터페이스를 갖고, OpenSSL은 사용하기에 상대적으로 복잡하고 메모리 관리 및 버그문제가 존재한다고 함
  + 애플리케이션용 슬롯을 별도로 생성하여 관리함
    + 슬롯(slot) : 독립된 영역으로 분리되어있는 논리적 개념으로 보면 됨
    + SoftHMS은 HSM을 여러 슬롯으로 나눔
    + 구조는 다음과 같음
       > <img src="https://prod-files-secure.s3.us-west-2.amazonaws.com/458aaba5-6bb4-49e3-ac84-7f827aa41d9e/590c5012-a816-44f0-b5f6-3026d9381c61/Untitled.png">
       
+ 지원 범위
  + 기본적으로 많이 활용되는 알고리즘을 지원함
    + RSA, DES/3DES, ECC, DHKE, Hash, HMAC, Key Generation ...
   
+ 설치 방법
  + 윈도우 설치 : SoftHSM v2 for Windows (msi 파일 설치) 후 환경변수 설정
  + 리눅스 설치 : 하단 두번째 github repostory에서 다운로드

+ 사용 방법(윈도우 기준)
  +  명령 프롬프트(cmd)를 통해 실행 (softhsm2-utils.exe 파일 활용)
  +  슬롯 조회 : `softhsm2-utils.exe --show-slots`
    + 출력 : 이용가능한 슬롯들이 출력됨 (초기에는 slot 0만 존재함)
  +  슬롯 초기화 : `softhsm2-utils.exe --init-token --slot <slot_number> --label <label_name>`
    + 명령어 입력후, SO와 User에 대한 PIN 설정 진행
      > SO는 User PIN 설정등을 하는 관리자 계정, User는 일반 HSM을 사용하는 사용자 계정을 의미함
    + 출력 : 입력한 라벨/PIN으로 설정된 슬롯을 포함한 슬롯들이 출력됨
  + 슬롯 삭제 : `softhsm2-utils.exe --delete-token --slot <slot_number>`
    + 결과 : 슬롯의 토큰이 삭제되며, 토큰 삭제 시 슬롯에서 저장된 모든 키와 인증서 삭제됨  
    
  

## 참고
+ [PKCS#11 관련 다른 사람의 블로그 💥](https://blog.naver.com/aepkoreanet/220754502731)
+ [HSM 관련 다른 사람의 블로그🌟](https://hyg4196.tistory.com/128)
+ [OpenDNSSEC SoftHSM v2 - Github](https://github.com/opendnssec/SoftHSMv2)
+ [OpenDNSSEC SoftHSM page](https://www.opendnssec.org/softhsm/)
  
