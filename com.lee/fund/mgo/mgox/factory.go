package mgox

import (
	"com.lee/fund/mgo"
)

func Open(name string) (*Database, error) {
	di, err := getInfo(name)
	if err != nil {
		return nil, err
	}

	session, err := open(di)
	if err != nil {
		return nil, err
	}

	return &Database{db: session.Copy().DB("")}, nil
}

func MustOpen(name string) *Database {
	di, err := getInfo(name)
	if err != nil {
		panic(err)
	}

	session, err := open(di)
	if err != nil {
		panic(err)
	}

	return &Database{db: session.Copy().DB("")}
}

func open(di *dbInfo) (*mgo.Session, error) {
	var err error

	if di.session != nil {
		return di.session, nil
	}

	// since config only need be initialized once, we can reuse the locker for initializing database
	_Locker.Lock()
	defer _Locker.Unlock()

	// double check
	if di.session != nil {
		return di.session, nil
	}

	// [mongodb://][user:pass@]host1[:port1][,host2[:port2],...][/database][?options]
	di.session, err = mgo.Dial(di.cfg.Settings.String("ConnString", ""))
	if err != nil {
		return nil, err
	}

	switch di.cfg.Settings.Int("ConsistencyMode", int(mgo.Monotonic)) {
	case int(mgo.Eventual):
		di.session.SetMode(mgo.Eventual, true)
	case int(mgo.Monotonic):
		di.session.SetMode(mgo.Monotonic, true)
	case int(mgo.Strong):
		di.session.SetMode(mgo.Strong, true)
	}

	return di.session, nil
}
