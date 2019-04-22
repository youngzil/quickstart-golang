package main

import (
	"bytes"
	crand "crypto/rand"
	"fmt"
	"github.com/hashicorp/consul/api"
	"log"
)

func main() {
	TestAPI_ClientPutGetDelete()
}

func TestAPI_ClientPutGetDelete() {
	c := makeClient()

	kv := c.KV()

	// Get a get without a key
	key := testKey()
	pair, _, err := kv.Get(key, nil)
	if err != nil {
		log.Fatal("consul client error : ", err)
	}
	if pair != nil {
		log.Fatal("unexpected value: %#v", pair)
	}

	value := []byte("test")

	// Put a key that begins with a '/', this should fail
	//invalidKey := "/test" //错误，因为key不能以/开头
	invalidKey := "test"
	p := &api.KVPair{Key: invalidKey, Value: value}
	if _, err := kv.Put(p, nil); err != nil {
		log.Fatal("Invalid key not detected: %s", invalidKey, err)
	}

	// Put the key
	p = &api.KVPair{Key: key, Flags: 42, Value: value}
	if _, err := kv.Put(p, nil); err != nil {
		log.Fatal("err: %v", err)
	}

	// Get should work
	pair, meta, err := kv.Get(key, nil)

	if err != nil {
		log.Fatalf("err: %v", err)
	}
	if pair == nil {
		log.Fatalf("expected value: %#v", pair)
	}

	log.Println("pair, meta=", pair.Key, pair.Value, meta)

	if !bytes.Equal(pair.Value, value) {
		log.Fatalf("unexpected value: %#v", pair)
	}
	if pair.Flags != 42 {
		log.Fatalf("unexpected value: %#v", pair)
	}
	if meta.LastIndex == 0 {
		log.Fatalf("unexpected value: %#v", meta)
	}

	log.Println("test get ", invalidKey)
	pair, meta, err = kv.Get(invalidKey, nil)
	log.Println("pair=", pair.Key, string(pair.Value))

	// Delete
	if _, err := kv.Delete(key, nil); err != nil {
		log.Fatalf("err: %v", err)
	}
	if _, err := kv.Delete(invalidKey, nil); err != nil {
		log.Fatalf("err: %v", err)
	}

	// Get should fail
	pair, _, err = kv.Get(key, nil)
	if err != nil {
		log.Fatalf("err: %v", err)
	}
	if pair != nil {
		log.Fatalf("unexpected value: %#v", pair)
	}
}

func makeClient() *api.Client {
	fmt.Println("test begin .")
	config := api.DefaultConfig()
	//config.Address = "localhost:8500"
	fmt.Println("defautl config : ", config)
	client, err := api.NewClient(config)
	if err != nil {
		log.Fatal("consul client error : ", err)
	}
	return client
}

func testKey() string {
	buf := make([]byte, 16)
	if _, err := crand.Read(buf); err != nil {
		panic(fmt.Errorf("Failed to read random bytes: %v", err))
	}

	return fmt.Sprintf("%08x-%04x-%04x-%04x-%12x",
		buf[0:4],
		buf[4:6],
		buf[6:8],
		buf[8:10],
		buf[10:16])
}
