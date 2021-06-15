const express = require('express')
//const bodyParser = require('body-parser')
const jwt = require('jsonwebtoken')
const app = express();// 라우팅 및 미들웨어 웹 프레임 워크(javascript)

const SECRET_KEY = "secret key" //

// 회원가입 한 유저 정보
let registryUsers = {

}

app.use(express.urlencoded({extended: false}))
app.use(express.json())
app.set('JWT_SECRET', SECRET_KEY) // JWT 인코딩/디코딩을 위해 필요한 키값

//로그인 체크하는 부분
const loginCheck = (req,res, next) => {
  let { token } = req.headers
  let { JWT_SECRET } = app.settings //설정한 secret값 대입 

  //token이 없으면 로그인 실패
  if (!token) {
    return res.status(401).json({
      msg: '인증실패'
    })
  }
  //있으면 검증
  jwt.verify(token, JWT_SECRET, (err, decoded) => {

    if (!decoded) {
      return res.status(401).json({
        msg: '인증실패'
      })
    }
    
    req.user = {
      id : decoded.id,
      name: decoded.name, 
      age : decoded.age
    }
    next()
  })
}

// 로그인 (get 요청)
app.get('/', (req, res, next) => {  

  let { id, password } = req.query //쿼리문을 통해 id, password 받고

  //해당 id 존재하는지 체크
  if ( !registryUsers[id] ) {
    return res.status(401).json({
      msg: 'login failed'
    })
  }

  //그에따른 비밀번호 일치하는지 체크
  if ( registryUsers[id].password !== password ) {
    return res.status(401).json({
      msg: 'login failed'
    })
  }

  let { JWT_SECRET } = app.settings //설정한 secret값 대입
  let payload = { 
    id : registryUsers[id].id,
    name: registryUsers[id].name,
    age : registryUsers[id].age
  }
  let options = { }// registered 정보
  //let options = {algorithm:"HS384"} 

  // JWT(JWS) 토큰생성 / default - HS256 (HMAC using SHA256)
  jwt.sign(payload, JWT_SECRET, options, (err, token) => {
    return res.status(201).json({ //완료 후, token 전달
      token,
      msg: 'success'
    })
  })
});

// 회원가입 (post)
app.post('/', (req, res, next) => {
  let {id, password, name, age} = req.body
  
  //입력받은 id 존재여부 체크
  if(registryUsers[id]) {
    return res.status(200).json({
      msg: '해당 아이디는 이미 등록되어있습니다.'
    })
  }

  // 유저정보 저장
  registryUsers[id] = {
    id,
    password,
    name,
    age
  }     

  return res.status(201).json({
    msg: 'success'
  })
})

// 로그인이 필요한 정보요청
app.get('/api', loginCheck, (req, res, next) => {
  return res.status(200).json({
    //data: [1,2,3,4],
    //post: [1,2,3,4,5],
    id : req.user.id,
    age : req.user.age,
    name: req.user.name
  })
})

//서버 listening (3000 port 사용)
app.listen(3000, () => {
  console.log('STARTING SERVER 3000 PORT');
});