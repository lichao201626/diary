package redis

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
