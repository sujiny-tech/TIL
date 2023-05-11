# AWS Intro
> 원티드 프리온보딩 백엔드 챌린지 수강하면서 알게 된 AWS 내용을 정리✍️

## AWS(Amazon Web Service) ?
> 아마존닷컴에서 개발한 클라우드 컴퓨팅 플랫폼


+ 일반적으로 서버 구축 방식은 크게 두 분류로 나뉨
   + On-premise : 기업/학교 등에서 자체 보유한 전산실 서버에 직접 설치하여 운영하는 방식.
   + **Cloud Computing** : 인터넷을 통해 데이터를 저장하거나 서버/네트워킹/소프트웨어와 같은 도구, 애플리케이션 등 다양한 서비스를 제공하는 방식.   

+ **Cloud Computing 이점**
   + 저렴한 비용 (직접 서버를 설치하고, 또다른 이슈로 서버를 철거해야되는 경우가 발생될 때 관리적으로도 시간적으로 유연하지 않음)
   + 속도 및 민첩성 (몇 분만에 전 세계에 배포가 가능하고, 데이터센터 운영/유지관리에 비용투자 불필요하므로 민첩성이 높음)

+ 2020년 기준으로 **AWS 사용률이 52% 정도**라고 함. 그다음으로는 Azure > Google Cloud > ... 순으로 사용률이 높음.
   + 우리나라 같은 경우에는 이미 AWS를 사용하여 구축된 서비스를 운영하는 회사가 많으므로 자연스럽게 AWS 비중이 높을 수 밖에 없음.   
   + 그외 Google Cloud의 경우, AWS에 없는 머신러닝 플랫폼이 차별화되어있어서 요즘 많이 사용된다고 함.
      + 각자 클라우드마다 서비스 차별화가 있으므로, 골라서 사용한다고 함.   

## AWS 지원 서비스
> AWS에서 지원하는 서비스는 엄청나게 많으므로 강사님께서 그 중에 알면 좋은 서비스들로 4가지 카테고리로 분류해주셨음.

+ **인프라** 관련
   + AWS API Gateway, AWS S3, AWS ELB, AWS CloudFront, AWS Secret Manager, 스냅샷
+ **컴퓨팅 파워**(서버)
   + AWS EC2, AWS Elastic Beanstalk, AWS ECS, AWS Fargate, AWS Lambda(Serverless)
+ **Message Queue**
   + AWS SQS, AWS MSK, AWS Kinesis
+ **Database**
   + AWS RDS, AWS DynamoDB, AWS ElasticCache



## 참고
+ 이전에 KDT AI 데브코스에서 배웠을 때 정리한 글도 있으므로 같이 참고하면 좋을 듯함.
   + [KDT AI 데브코스 - Cloud Computing 작성 글](https://github.com/sujiny-tech/k-digital-training-AI-dev/blob/main/Web-Handling-with-Python/Cloud%20Computing.md)   


