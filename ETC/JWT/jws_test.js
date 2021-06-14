const jws= require('../node-jws');
const signature = jws.sign({
    header: {alg : 'HS256'},
    payload: 'sujinlee test', 
    secret: 'secret value for HMAC',
    encoding: 'utf-8',
}); //sign

console.log("(JWS) signature : ", signature);

//decode sign
const parts = jws.decode(signature);
console.log("(JWS) decoding : ", parts);

const alg_ = 'HS256';
const secret_ = 'secret value for HMAC';

//verify jws
var valid=jws.verify(signature, alg_, secret_);
console.log("verify result : ", valid);

