// DO NOT EDIT!!! GENERETED BY ragel AS:
//  ragel -Z -G2 -e -o ragel_parse_datetime.go parse_datetime.go.rl
// it might be helpful to generate the FSM graph in debug:
//   ragel -Vp parse_datetime.go.rl -o parse_datetime.dot
//   dot parse_datetime.dot -Tpng -o parse_datetime.png

package anytime

import (
    "fmt"
    "strconv"
    "errors"
    "strings"
)

%%{
machine datetime_parser;

include common "datetime.rl";

main := fulldatetime;

write data;
}%%

%% write data noerror noprefix nofinal;

func ragelParse(data string) (st parsedTime, err error) {
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

