// Go MySQL Driver - A MySQL-Driver for Go's database/sql package
//
// Copyright 2017 The Go-MySQL-Driver Authors. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this file,
// You can obtain one at http://mozilla.org/MPL/2.0/.

// +build go1.8

package mysql

import (
	"crypto/tls"
	"database/sql"
	"database/sql/driver"
	"errors"
<<<<<<< HEAD
	"fmt"
=======
>>>>>>> 258d5c409a01370dfe542ceadc3d1669659150fe
)

func cloneTLSConfig(c *tls.Config) *tls.Config {
	return c.Clone()
}

func namedValueToValue(named []driver.NamedValue) ([]driver.Value, error) {
	dargs := make([]driver.Value, len(named))
	for n, param := range named {
		if len(param.Name) > 0 {
			// TODO: support the use of Named Parameters #561
			return nil, errors.New("mysql: driver does not support the use of Named Parameters")
		}
		dargs[n] = param.Value
	}
	return dargs, nil
}

func mapIsolationLevel(level driver.IsolationLevel) (string, error) {
	switch sql.IsolationLevel(level) {
	case sql.LevelRepeatableRead:
		return "REPEATABLE READ", nil
	case sql.LevelReadCommitted:
		return "READ COMMITTED", nil
	case sql.LevelReadUncommitted:
		return "READ UNCOMMITTED", nil
	case sql.LevelSerializable:
		return "SERIALIZABLE", nil
	default:
<<<<<<< HEAD
		return "", fmt.Errorf("mysql: unsupported isolation level: %v", level)
=======
		return "", errors.New("mysql: unsupported isolation level: " + string(level))
>>>>>>> 258d5c409a01370dfe542ceadc3d1669659150fe
	}
}
