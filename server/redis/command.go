package redis

import "bytes"

// String-related commands

func Set(key string, value []byte) error {
	con := Connection{}
	// var val []byte
	/* 	c, err := client.popCon()
	   	if err == io.EOF {
	   		c, err = client.openConnection()
	   		if err != nil {
	   			println(err.Error())
	   			goto End
	   		}

	   		data, err = client.rawSend(c, b)
	   	} */
	c, _ := con.Connect()
	b := CommandBytes("SET", key, string(value))
	_, err := con.RawSend(c, b)
	defer c.Close()
	return err
}

func Get(key string) ([]byte, error) {
	con := Connection{}
	c, _ := con.Connect()
	b := CommandBytes("GET", key)
	data, err := con.RawSend(c, b)
	defer c.Close()
	return data.([]byte), err
}

func RandomKey() ([]byte, error) {
	con := Connection{}
	c, _ := con.Connect()
	b := CommandBytes("RANDOMKEY")
	data, err := con.RawSend(c, b)
	defer c.Close()
	return data.([]byte), err
}

func Keys(pattern string) ([]string, error) {
	con := Connection{}
	c, _ := con.Connect()
	b := CommandBytes("KEYS", pattern)
	res, err := con.RawSend(c, b)

	if err != nil {
		return nil, err
	}

	var ok bool
	var keydata [][]byte

	if keydata, ok = res.([][]byte); ok {
		// key data is already a double byte array
	} else {
		keydata = bytes.Fields(res.([]byte))
	}
	ret := make([]string, len(keydata))
	for i, k := range keydata {
		ret[i] = string(k)
	}
	return ret, nil
}

/*
func (client *Client) Getset(key string, val []byte) ([]byte, error) {
	res, err := client.sendCommand("GETSET", key, string(val))

	if err != nil {
		return nil, err
	}

	data := res.([]byte)
	return data, nil
}
*/
