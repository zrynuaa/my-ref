package main

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"golang.org/x/crypto/ripemd160"
	"golang.org/x/crypto/sha3"

	"bytes"
	"encoding/binary"
	"fmt"
	"hash"
	"io/ioutil"
	"math/rand"
	"time"
)

func main() {
	//生成多个随机测试数据，分别是1字节、1MB、30MB、90MB
	//ioutil.WriteFile("1B.txt", []byte("1"), 0644)
	//generateRanddata("1KB.txt", 25)
	//generateRanddata("1MB.txt", 720)
	//generateRanddata("30MB.txt", 645*6)
	//generateRanddata("90MB.txt", 671*10)

	//开始测试
	//files := []string{"1B.txt","1KB.txt","1MB.txt","30MB.txt"}
	files := []string{"1B.txt","1KB.txt"}

	for _,v :=range files{
		fmt.Println("\n",v)
		dat,_ := ioutil.ReadFile(v)//读取相应文件中的信息
		plain := []byte(dat)
		fmt.Println(len(plain))

		//fmt.Print("MD5: ")
		md := md5.New()
		testMD(md, plain)

		//fmt.Print("SHA-1: ")
		md = sha1.New()
		testMD(md, plain)

		//fmt.Print("SHA224: ")
		md = sha256.New224()
		testMD(md, plain)

		//fmt.Print("SHA256: ")
		md = sha256.New()
		testMD(md, plain)

		//fmt.Print("SHA384: ")
		md = sha512.New384()
		testMD(md, plain)

		//fmt.Print("SHA512: ")
		md = sha512.New()
		testMD(md, plain)

		//fmt.Print("SHA512/224: ")
		md = sha512.New512_224()
		testMD(md, plain)

		//fmt.Print("SHA512/256: ")
		md = sha512.New512_256()
		testMD(md, plain)

		//fmt.Print("RIPEMD160: ")
		ripe := ripemd160.New()
		testMD(ripe, plain)

		//fmt.Print("SHA3-224: ")
		md = sha3.New224()
		testMD(md, plain)

		//fmt.Print("SHA3-256: ")
		md = sha3.New256()
		testMD(md, plain)

		//fmt.Print("SHA3-384: ")
		md = sha3.New384()
		testMD(md, plain)

		//fmt.Print("SHA3-512: ")
		md = sha3.New512()
		testMD(md, plain)

		//fmt.Print("SHAKE128: ")
		//testShake128(plain)
		//
		//fmt.Print("SHAKE256: ")
		//testShake256(plain)

	}
}

func testMD(md hash.Hash, msg []byte) time.Duration {
	//fmt.Println("输出大小", md.Size())

	t1 := time.Now()

	for i := 0; i < 10000; i++ {
		md.Reset()
		md.Write(msg)
		md.Sum(nil)
		//fmt.Printf("%x", md.Sum(nil))
		//fmt.Println(" ")
	}
	t2 := time.Now()
	fmt.Println(t2.Sub(t1))
	return t2.Sub(t1)
}



func check(e error){
	if e!=nil{
		panic(e)
	}
}

func generateRanddata(filename string, round int){
	rand.Seed(time.Now().Unix())
	buf := bytes.NewBuffer([]byte{})
	var data []byte

	for index:=0; index<round; index++{
		x := rand.Int()
		binary.Write(buf, binary.BigEndian, int32(x))
		data = append(data, buf.Bytes()...)
	}

	err:= ioutil.WriteFile(filename, data, 0644)
	check(err)
	//fmt.Println(len(data))
}




