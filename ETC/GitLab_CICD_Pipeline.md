# GitLab - CI/CD Pipeline
> GitLab 사용하면서 CI/CD 파이프라인에 대해 정리 📝


## CI/CD
+ 개념
   + CI(Continuous Integration) : 변경사항이 일어날 때, 정기적으로 테스트, 빌드 과정을 거쳐 공유 레포지토리에 통합되는 과정 
   + CD(Continuous Delivery/Deployment) : CI 과정을 거친 애플리케이션을 자동적으로 레포지토리에 업로드하거나 서비스 환경에 바로 배포하는 것을 의미
+ GiLab CI/CD Pipeline
   + GitLab에서는 Commit/Merge/Release 등 이벤트가 발생하는 것을 감지하고 CI/CD Pipeline이 동작할 수 있는 환경을 제공함

## 기본설정
+ gitlab-runner 설치
+ gitlab-runner 등록
+ gitlab-runner 등록 확인 및 사용자 권한 설정
+ 프로젝트에 .gitlab-ci.yml 작성
+ pipeline 실행
+ ... ing



# 참고
+ [잘 정리된 다른사람의 블로그](https://velog.io/@leesomyoung/CICD%EB%A5%BC-%EC%9C%84%ED%95%9C-gitlab-pipeline-%EA%B5%AC%EC%B6%95%ED%95%98%EA%B8%B0)
+ [잘 정리된 다른사람의 블로그](https://sg-choi.tistory.com/552)
