package database

type IConnectionDB interface {
	Close()
	Conn() interface{}
}

type ConnectDB struct {
	source interface{}
}

func (c *ConnectDB) Source() interface{} {
	return c.source
}
