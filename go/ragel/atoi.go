
//line atoi.go.rl:1
package ragel

import (
  "fmt"
)


//line atoi.go:11
const atoi_start int = 1
const atoi_first_final int = 3
const atoi_error int = 0

const atoi_en_main int = 1


//line atoi.go.rl:10


func Atoi(data string) (int64, error) {
  cs, p, pe := 0, 0, len(data)
  var neg bool
  var val int64
  
//line atoi.go:27
	{
	cs = atoi_start
	}

//line atoi.go:32
	{
	if p == pe {
		goto _test_eof
	}
	switch cs {
	case 1:
		goto st_case_1
	case 0:
		goto st_case_0
	case 2:
		goto st_case_2
	case 3:
		goto st_case_3
	case 4:
		goto st_case_4
	}
	goto st_out
	st_case_1:
		switch data[p] {
		case 43:
			goto st2
		case 45:
			goto tr2
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto tr3
		}
		goto st0
st_case_0:
	st0:
		cs = 0
		goto _out
tr2:
//line atoi.go.rl:17
 neg = true 
	goto st2
	st2:
		if p++; p == pe {
			goto _test_eof2
		}
	st_case_2:
//line atoi.go:74
		if 48 <= data[p] && data[p] <= 57 {
			goto tr3
		}
		goto st0
tr3:
//line atoi.go.rl:18
 val = val * 10 + (int64(data[p]) - '0') 
	goto st3
	st3:
		if p++; p == pe {
			goto _test_eof3
		}
	st_case_3:
//line atoi.go:88
		if data[p] == 10 {
			goto st4
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto tr3
		}
		goto st0
	st4:
		if p++; p == pe {
			goto _test_eof4
		}
	st_case_4:
		goto st0
	st_out:
	_test_eof2: cs = 2; goto _test_eof
	_test_eof3: cs = 3; goto _test_eof
	_test_eof4: cs = 4; goto _test_eof

	_test_eof: {}
	_out: {}
	}

//line atoi.go.rl:24


  if neg {
    val = -val
  }

  if cs < atoi_first_final {
    return 0, fmt.Errorf("atoi parse error: %s", data)
  }

  return val, nil
}