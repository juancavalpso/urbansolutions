package modelo

import (
	"crypto/tls"
	"io/ioutil"
	"net"
	"net/http"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type MyEntity struct {
	Data []byte `json:"data" bson:"data"`
}

type Materiales struct {
	ID     primitive.ObjectID `bson:"_id,omitempty"`
	Title  string             `bson:"title,omitempty"`
	Author string             `bson:"author,omitempty"`
	Tags   []string           `bson:"tags,omitempty"`
}

var mongoConn *mgo.Session

func (){
	var err error
	mongoConn, err = createConnection()
	if err != nil {
		panic(err)
	}
}

func createConnection() (*mgo.Session, error) {
	dialInfo := mgo.DialInfo{
		Addrs: []string{
			"cluster0-shard-00-00.fazen.mongodb.net:27017",
			"cluster0-shard-00-01.fazen.mongodb.net:27017",
			"cluster0-shard-00-02.fazen.mongodb.net:27017"},
		Username: "bitaprod",
		Password: "bitaprod",
	}
	tlsConfig := &tls.Config{}
	dialInfo.DialServer = func(addr *mgo.ServerAddr) (net.Conn, error) {
		conn, err := tls.Dial("tcp", addr.String(), tlsConfig)
		return conn, err
	}
	return mgo.DialWithInfo(&dialInfo)
}

func GetOutil() Materiales{
	session := mongoConn.Copy()
	defer session.Close()

	materiales := Materiales{}
	err := session.DB("testeando").C("materiales").Find(bson.M{}).One(&entity)
	if err != nil {
		panic(err)
	}

	return materiales
}

// func SaveOutil(){
// 	payload, err := ioutil.ReadAll(req.Body)
// 	if err != nil {
// 		panic(err)
// 	}
// 	session := mongoConn.Copy()
// 	defer session.Close()

// 	entity := MyEntity{Data: payload}
// 	err = session.DB("test").C("data").Insert(entity)
// 	if err != nil {
// 		panic(err)
// 	}
// }

