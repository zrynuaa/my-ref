// 本地需安装NodeJS环境，该文档是在8.11.3版本下测试的，高版本应该也可以使用
// 依赖包 npm install js-sha3
// node testhash.js

const crypto = require('crypto');
const fs = require('fs');
const sha3js = require('js-sha3');

var array = ["1B.txt","1KB.txt","1MB.txt","30MB.txt"];
//var array = ["1MB.txt"];
const N = 10000
array.forEach(function(v,i,a){
        console.log(v);
        var data = fs.readFileSync(v);

        //md5 sha1 sha256 sha512 ripemd160
        testEachHash("md5", data)
        testEachHash("sha1", data)
        testEachHash("sha224", data)
        testEachHash("sha256", data)
        testEachHash("SHA384", data)
        testEachHash("sha512", data)
        testEachHash("ripemd160", data)

        testsha3224(data)
        testsha3256(data)
        testsha3384(data)
        testsha3512(data)
});

function testsha3224(data){
    t1 = new Date().getTime();
    for( var i = 0; i<N; i++) {
        var hash = sha3js.sha3_224.create();
        hash.update(data);
        //console.log(hash.hex());
        hash.digest('hex');
    }
    t2 = new Date().getTime()
    interval = (t2 - t1)/1000;
    console.log("sha3-224: ", interval, "s")
}

function testsha3256(data){
    t1 = new Date().getTime();
    for( var i = 0; i<N; i++) {
        var hash = sha3js.sha3_256.create();
        hash.update(data);
        //console.log(hash.hex());
        hash.digest('hex');
    }
    t2 = new Date().getTime()
    interval = (t2 - t1)/1000;
    console.log("sha3-256: ", interval, "s")
}

function testsha3384(data){
    t1 = new Date().getTime();
    for( var i = 0; i<N; i++) {
        var hash = sha3js.sha3_384.create();
        hash.update(data);
        //console.log(hash.hex());
        hash.digest('hex');
    }
    t2 = new Date().getTime()
    interval = (t2 - t1)/1000;
    console.log("sha3-384: ", interval, "s")
}

function testsha3512(data){
    t1 = new Date().getTime();
    for( var i = 0; i<N; i++) {
        var hash = sha3js.sha3_512.create();
        hash.update(data);
        //console.log(hash.hex());
        hash.digest('hex');
    }
    t2 = new Date().getTime()
    interval = (t2 - t1)/1000;
    console.log("sha3-512: ", interval, "s")
}

function testEachHash(hashname, data) {
    t1 = new Date().getTime();
    for( var i = 0; i<N; i++) {
        const hash = crypto.createHash(hashname);
        hash.update(data)
        //console.log(hash.digest('hex'));
        hash.digest('hex');
    }
    t2 = new Date().getTime()
    interval = (t2 - t1)/1000;
    console.log(hashname, ": ", interval, "s")
}


