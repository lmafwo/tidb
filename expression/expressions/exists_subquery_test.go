// Copyright 2015 PingCAP, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// See the License for the specific language governing permissions and
// limitations under the License.

package expressions

import (
	. "github.com/pingcap/check"
	"github.com/pingcap/tidb/util/types"
)

var _ = Suite(&testExistsSubQuerySuite{})

type testExistsSubQuerySuite struct {
}

func (s *testExistsSubQuerySuite) TestExistsSubQuery(c *C) {
	tbl := []struct {
		in     []interface{}
		result int64 // 0 for false, 1 for true.
	}{
		{[]interface{}{1}, 1},
		{[]interface{}{nil}, 1},
		{[]interface{}{}, 0},
		{nil, 0},
	}

	for _, t := range tbl {
		in := make([][]interface{}, 0, len(t.in))
		for _, v := range t.in {
			in = append(in, []interface{}{convert(v)})
		}

		sq := newMockSubQuery(in, []string{"c"})
		expr := NewExistsSubQuery(sq)

		c.Assert(expr.IsStatic(), IsFalse)

		str := expr.String()
		c.Assert(len(str), Greater, 0)

		v, err := expr.Eval(nil, nil)
		c.Assert(err, IsNil)

		val, err := types.ToBool(v)
		c.Assert(err, IsNil)
		c.Assert(val, Equals, t.result)
	}
}
