package main

/**
 * Hash with MD5
 */

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
)

func Checksum(data []byte, salt string) string {

	var jsonObj GA
	err := json.Unmarshal(data, &jsonObj)

	if err != nil {
		fmt.Println(err)
	}

	hash := md5.New()
	defer hash.Reset()
	str := fmt.Sprintf("%s%d%s%s%d%s%d%f%s%s%d%d%d%d%s%s%s%d",
		salt,
		jsonObj.v,
		jsonObj.tid,
		jsonObj.gtm,
		jsonObj._p,
		jsonObj._z,
		jsonObj._pp,
		jsonObj.cid,
		jsonObj.ul,
		jsonObj.sr,
		jsonObj._s,
		jsonObj.sid,
		jsonObj.sct,
		jsonObj.seg,
		jsonObj.dl,
		jsonObj.dr2,
		jsonObj.en,
		jsonObj._ss,
	)

	hash.Write([]byte(str))

	return hex.EncodeToString(hash.Sum(nil))
}
