const jws= require('../node-jws');
const parts = jws.decode('eyJhbGciOiJIUzI1NiJ9.c3VqaW5sZWUgdGVzdA.Z9BKLH2WxoVdiL8-WA5Ii1-XmzjOf_SPqFqnfcnwPQg');

console.log("(JWS) decoded : ", parts);