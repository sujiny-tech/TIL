var jwt = require('jsonwebtoken');
var tokenKey = "TEST_KEY11";
var token = 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1aWQiOjE0NTU0LCJpYXQiOjE2MjM2NDIwODgsImV4cCI6MTYyMzY0MzUyOH0.9IiT9ZltMf1MCuK5SEctNfOiusjqNN6RBAtc1BYz5sQ';

//비동기 처리
jwt.verify(token, tokenKey, function(err, decoded){
    console.log("sync:", decoded);
});

//동기 처리
try {
    var decoded = jwt.verify(token, tokenKey);
    console.log("async:", decoded);
} catch(err) {
    console.log(err);
}