// GENERETED .hpp BY ragel AS:
//   ragel-go -G2 -e -o ragel_parse_datetime.go ragel_parse_datetime.go.rl
// it might be helpful to generate the FSM graph in debug:
//   ragel -Vp ragel_parse_datetime.go.rl -o ragel_parse_datetime.dot
//   dot ragel_parse_datetime.dot -Tpng -o ragel_parse_datetime.png

package ragel

import (
    "fmt"
    "strconv"
    "errors"
)

%%{
machine datetime_parser;

include common "datetime.rl";

main := fulldatetime;

write data;
}%%

%% write data noerror noprefix nofinal;

func Parse(data string) (st ParsedDatetime, err error) {
    cs, p, pe := 0, 0, len(data)
    eof := pe
    pb := p

    %%write init;
    %%write exec;

    if p != eof {  // input date not fully consumed
        err = errors.New("input data not fully consumed");
        return
    }

    if cs < datetime_parser_first_final {
        err = fmt.Errorf("time parse error: %s", data)
    }
    return
}

