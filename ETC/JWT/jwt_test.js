/*--------------------sign-------------------- */
var jwt      = require('jsonwebtoken');
var valid_tokenKey = "Test tokenKey"; //토큰키 서버에서 보관 중요
var payLoad  = {'uid':14554, 'name':'sujiny2222'};
var token = jwt.sign(payLoad,valid_tokenKey,{
    algorithm : 'HS256', //"HS256", "HS384", "HS512", "RS256", "RS384", "RS512" default SHA256
    expiresIn : 1440 // default : (s)
});
console.log("token : ", token);

/*-------------------decode------------------- */
var invalid_tokenKey = "TEST_KEY11";

//비동기 처리
jwt.verify(token, valid_tokenKey, function(err, decoded){
    console.log("sync:", decoded);
});

//비동기 처리
jwt.verify(token, invalid_tokenKey, function(err, decoded){
    console.log("sync:", decoded);
});

/*
//동기 처리
try {
    var decoded = jwt.verify(token, valid_tokenKey);
    console.log("async:", decoded);
} catch(err) {
    console.log(err);
}
*/

