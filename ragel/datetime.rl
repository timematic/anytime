
%%{
machine common;

action mark_pb { pb = p }

action mark_zone { st.Zoned = true }

action parse_month_digit {
    st.Month, _ = strconv.Atoi(data[pb:p])
}

action parse_mmdd_4_digit {
    st.Month, _ = strconv.Atoi(data[pb:pb+2])
    st.Day, _ = strconv.Atoi(data[pb+2:pb+4])
}
action parse_year_4_digit {
    st.Year, _ = strconv.Atoi(data[pb:pb+4])
}

action parse_year_2_digit {
    st.Year = parse_year_2_digits(data[pb:pb+2])
}

action parse_day_of_year {
    st.DayOfYear, _ = strconv.Atoi(data[pb:pb+3])
}

action parse_yyyymmdd {
    st.Year, _ = strconv.Atoi(data[pb:pb+4])
    st.Month, _ = strconv.Atoi(data[pb+4:pb+6])
    st.Day, _ = strconv.Atoi(data[pb+6:pb+8])
}

action parse_hhmmss_digits {
    switch p - pb {
        case 4: 
            st.Hour, _ = strconv.Atoi(data[pb:pb+2])
            st.Minute, _ = strconv.Atoi(data[pb+2:pb+4])
        case 6:
            st.Hour, _ = strconv.Atoi(data[pb:pb+2])
            st.Minute, _ = strconv.Atoi(data[pb+2:pb+4])
            st.Second, _ = strconv.Atoi(data[pb+4:pb+6])
        default:
            err = errors.New("invalid hhmmss digits")
            return
    }
}

action set_bc {
    st.Ad_bc = ADBC_BC;
}

action set_ampm {
    if st.Hour > 12 {
        err = errors.New("hour out of range")
        return st, err
    }
    if apm, err := parse_ampm(data[pb:]); err != nil {
        return st, err
    } else {
        switch apm {
            case AMPM_AM:
                if (st.Hour == 12) {
                    st.Hour -= 12; // 12:00:00 am == 00:00:00
                }
            case AMPM_PM: {
                if (st.Hour < 12) {
                    st.Hour += 12
                }
                // else {} // 12:00:00 pm = 12:00:00, do nothing
            }
        }
    }
}

action set_month_1 { st.Month = 1 }
action set_month_2 { st.Month = 2 }
action set_month_3 { st.Month = 3 }
action set_month_4 { st.Month = 4 }
action set_month_5 { st.Month = 5 }
action set_month_6 { st.Month = 6 }
action set_month_7 { st.Month = 7 }
action set_month_8 { st.Month = 8 }
action set_month_9 { st.Month = 9 }
action set_month_10 { st.Month = 10 }
action set_month_11 { st.Month = 11 }
action set_month_12 { st.Month = 12 }

action parse_day_digit {
    switch p - pb {
        case 1,2: st.Day, _ = strconv.Atoi(data[pb:p])
        default:
            err = fmt.Errorf("invalid day digits %s", data[pb:p])
            return
    }
}

action parse_hour_2_digit {
    st.Hour, _ = strconv.Atoi(data[pb:pb+2])
}
action parse_hour_1_digit {
    st.Hour, _ = strconv.Atoi(data[pb:pb+1])
}
action parse_minute_2_digit {
    st.Minute, _ = strconv.Atoi(data[pb:pb+2])
}
action parse_minute_1_digit {
    st.Minute, _ = strconv.Atoi(data[pb:pb+1])
}
action parse_second_2_digit {
    st.Second, _ = strconv.Atoi(data[pb:pb+2])
}
action parse_second_1_digit {
    st.Second, _ = strconv.Atoi(data[pb:pb+1])
}
action parse_second_fraction {
    switch p - pb {
        case 1: st.Millisecond = parse_digits(data[pb:pb+1]) * 100
        case 2: st.Millisecond = parse_digits(data[pb:pb+2])  * 10
        case 3: st.Millisecond = parse_digits(data[pb:pb+3]) 
        case 4: 
            st.Millisecond = parse_digits(data[pb:pb+3]) 
            st.Microsecond = parse_digits(data[pb+3:pb+4])  * 100
        case 5: 
            st.Millisecond = parse_digits(data[pb:pb+3]) 
            st.Microsecond = parse_digits(data[pb+3:pb+5]) * 10
        case 6:
            st.Millisecond = parse_digits(data[pb:pb+3])
            st.Microsecond = parse_digits(data[pb+3:pb+6])
        case 7:
            st.Millisecond = parse_digits(data[pb:pb+3])
            st.Microsecond = parse_digits(data[pb+3:pb+6])
            st.Nanosecond =  parse_digits(data[pb+6:pb+7]) * 100
        case 8:
            st.Millisecond = parse_digits(data[pb:pb+3])
            st.Microsecond = parse_digits(data[pb+3:pb+6])
            st.Nanosecond =  parse_digits(data[pb+6:pb+8]) * 10
        default: 
            st.Millisecond = parse_digits(data[pb:pb+3])
            st.Microsecond = parse_digits(data[pb+3:pb+6])
            st.Nanosecond =  parse_digits(data[pb+6:p])
    }
}

action mark_negative_offset { st.NegtiveZoneOffset = true }

action parse_offset_hour {
    switch p - pb {
        case 1,2: st.ZoneOffsetHour, _ = strconv.Atoi(data[pb:p])
        default:
            err = errors.New("invalid offset hour")
            return
    }
}

action parse_offset_minute {
    switch p - pb {
        case 1,2: st.ZoneOffsetMinute, _ = strconv.Atoi(data[pb:p])
        default:
            err = errors.New("invalid offset minute")
            return
    }
}

action parse_offset_digits {
    // 1 as 1 hour
    // 12 as 12 hours
    // 123 as 1 hour 23 minutes
    // 1234 as 12 hours and 34 minutes
    // 如果超过4位则移除前缀0直到保留后4位；移除前缀0后如果还超过4位则溢出报错
    // - 00000012 as 12 minutes
    // - 0000001234 as 12 hours and 34 minutes
    for p - pb > 4 &&  data[pb] =='0' {
        pb += 1 
    }
    switch p-pb {
        case 1,2:{st.ZoneOffsetHour, _ = strconv.Atoi(data[pb:p])}
        case 3,4:{
            num := parse_digits(data[pb:p])
            st.ZoneOffsetHour = num/100
            st.ZoneOffsetMinute = num%100
            if st.ZoneOffsetMinute >=60 || st.ZoneOffsetHour>=15 {
                err = errors.New("invalid offset digits")
                return
            } 
        }
        default: 
            err = errors.New("invalid offset digits")
            return
    }
}

action parse_timezone_abbr {
    st.ZoneName = data[pb:p]
    st.Zoned = true
}

# whitespace
sp = ' ';

# 1..9
nonzerodigit = '1'..'9';

# 00..59
sexagesimal = '0'..'5' . '0'..'9';

# 0..9 | 10..31
day = (nonzerodigit | '0' . nonzerodigit | '1'..'2' . '0'..'9' | '3' . '0'..'1') >mark_pb %parse_day_digit;
day_2digit =  ('0' . nonzerodigit | '1'..'2' . '0'..'9' | '3' . '0'..'1') >mark_pb %parse_day_digit;


# Jan .. Dec
month_name = (
('Jan' | 'January' | 'jan' | 'january') %set_month_1 |
('Feb' | 'February' | 'feb' | 'february') %set_month_2 |
('Mar' | 'March' | 'mar' | 'march') %set_month_3 |
('Apr' | 'April' | 'apr' | 'april') %set_month_4 |
('May' | 'may') %set_month_5 |
('Jun' | 'June' | 'jun' | 'june') %set_month_6 |
('Jul' | 'July' | 'jul' | 'july') %set_month_7 |
('Aug' | 'August' | 'aug' | 'august') %set_month_8 |
('Sep' | 'Sept' | 'September' | 'sep' | 'sept' | 'september') %set_month_9 |
('Oct' | 'October' | 'oct' | 'october') %set_month_10 |
('Nov' | 'November' | 'nov' | 'november') %set_month_11 |
('Dec' | 'December' | 'dec' | 'december') %set_month_12) '.'?;
month_digits = (nonzerodigit | '0' . nonzerodigit | '1' . '0'..'2' ) >mark_pb %parse_month_digit;
month = (month_name | month_digits);
month_2_digit = ('0' . nonzerodigit | '1' . '0'..'2' ) >mark_pb %parse_month_digit;

week_day_name = ('Monday'|'Mon' | 'Tuesday'|'Tue' | 'Wednesday'|'Wed' | 'Thursday'|'Thu' | 'Friday'|'Fri' | 'Saturday'|'Sat' | 'Sunday'|'Sun');

year_4digit = digit{4} >mark_pb %parse_year_4_digit;
year_2digit = digit{2} >mark_pb %parse_year_2_digit;

day_of_year = digit{3} >mark_pb %parse_day_of_year;
 
mmdd = digit{4} >mark_pb %parse_mmdd_4_digit;

ad_bc = 'AD' | ('BC' %set_bc);

datesp = ('-' | '/' | '.' | sp);
ymd = year_4digit datesp month datesp day;
dmy = day datesp month_name datesp year_4digit;
mdy = month_name datesp day datesp year_4digit;
yyyyddd = year_4digit ('-' | '/' | '.')? day_of_year;
yyyymmdd = year_4digit mmdd;
date_rfc1123 = week_day_name ',' sp day sp month_name sp year_4digit;
date_rfc850 = week_day_name ',' sp day '-' month_name '-' year_2digit;
date_rfc822 = day sp month_name sp year_2digit;
date = ( ymd | mdy | dmy | yyyyddd | yyyymmdd | date_rfc1123 | date_rfc850 | date_rfc822);

# 01..23
hour_2_digit = ('0'..'1' . '0'..'9' | '2' . '0'..'3') >mark_pb %parse_hour_2_digit;
hour_1_digit = digit >mark_pb %parse_hour_1_digit;
hour = (hour_1_digit | hour_2_digit);

minute_2_digit = sexagesimal >mark_pb %parse_minute_2_digit;
minute_1_digit = digit >mark_pb %parse_minute_1_digit;
minute = (minute_1_digit | minute_2_digit);

second_2_digit = sexagesimal >mark_pb %parse_second_2_digit;
second_1_digit = digit >mark_pb %parse_second_1_digit;
second_fraction = ('.' | ',') digit{1,9} >mark_pb %parse_second_fraction;
second = (second_1_digit | second_2_digit) . second_fraction?;

timeoffset_hour = (digit | '0'..'1' . '0'..'9' | '2' . '0'..'3') >mark_pb %parse_offset_hour;
timeoffset_minute = (digit | sexagesimal) >mark_pb %parse_offset_minute;
timeoffset_hhmm = timeoffset_hour ':' (timeoffset_minute)?;
timeoffset_digits = digit{1,} > mark_pb %parse_offset_digits;
timenumoffset = ('+' | '-' %mark_negative_offset) (timeoffset_hhmm | timeoffset_digits);

timezone_abbreviation = (alpha | '/' | '_'){3,} >mark_pb %parse_timezone_abbr;
timezone = ('Z' | timenumoffset (sp timezone_abbreviation)? | timezone_abbreviation) %mark_zone;

am_pm = ('am' | 'pm' | 'AM' | 'PM') >mark_pb %set_ampm;

hhmmss = hour ( ':' minute ( ':' second)? )? (sp? am_pm)?;
hhmmss_digits = ( (digit{4} | digit{6}) >mark_pb %parse_hhmmss_digits) . second_fraction? ;
time =  (hhmmss_digits | hhmmss) . (sp)? . timezone?;
time_without_zone = hhmmss_digits | hhmmss;

fulldate = ( date . ('T' | sp)? . timezone? . (sp ad_bc)?);

# "Mon Jan 02 15:04:05 -0700 2006"
ruby_datetime = week_day_name sp month_name sp day_2digit sp time sp year_4digit; # "Mon Jan 02 15:04:05 -0700 2006"
pg_datetime = (week_day_name sp)? month_name sp day_2digit sp time_without_zone sp year_4digit (sp timezone)?; # "Mon Jan 02 15:04:05 2006 PST"
america_datetime = month_name sp (day . ','?) sp (year_4digit . ','?) (sp time)?; # "January 02, 2006, 15:04:05"
unix_datetime = week_day_name sp month_name sp (sp? day) sp time sp year_4digit; # "Mon Jan  2 15:04:05 -0700 2006"

fulldatetime = fulldate | ( date ('T' | sp | '_' | 't') time (sp ad_bc)?) | ruby_datetime | pg_datetime | unix_datetime | america_datetime;

fulltime = ('T'? . time) | ( date ('T' | sp) time );

}%%
 