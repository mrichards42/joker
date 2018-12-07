// This file is generated by generate-std.joke script. Do not edit manually!

package time

import (
	"time"
	. "github.com/candid82/joker/core"
)

var timeNamespace = GLOBAL_ENV.EnsureNamespace(MakeSymbol("joker.time"))

var add_ Proc = func(_args []Object) Object {
	_c := len(_args)
	switch  {
	case _c == 2:
		
		t := ExtractTime(_args, 0)
    d := ExtractInt(_args, 1)
		_res := t.Add(time.Duration(d))
		return MakeTime(_res)

	default:
		PanicArity(_c)
	}
	return NIL
}

var format_ Proc = func(_args []Object) Object {
	_c := len(_args)
	switch  {
	case _c == 2:
		
		t := ExtractTime(_args, 0)
    layout := ExtractString(_args, 1)
		_res := t.Format(layout)
		return MakeString(_res)

	default:
		PanicArity(_c)
	}
	return NIL
}

var from_unix_ Proc = func(_args []Object) Object {
	_c := len(_args)
	switch  {
	case _c == 2:
		
		sec := ExtractInt(_args, 0)
    nsec := ExtractInt(_args, 1)
		_res := time.Unix(int64(sec), int64(nsec))
		return MakeTime(_res)

	default:
		PanicArity(_c)
	}
	return NIL
}

var hours_ Proc = func(_args []Object) Object {
	_c := len(_args)
	switch  {
	case _c == 1:
		
		d := ExtractInt(_args, 0)
		_res := time.Duration(d).Hours()
		return MakeDouble(_res)

	default:
		PanicArity(_c)
	}
	return NIL
}

var minutes_ Proc = func(_args []Object) Object {
	_c := len(_args)
	switch  {
	case _c == 1:
		
		d := ExtractInt(_args, 0)
		_res := time.Duration(d).Minutes()
		return MakeDouble(_res)

	default:
		PanicArity(_c)
	}
	return NIL
}

var now_ Proc = func(_args []Object) Object {
	_c := len(_args)
	switch  {
	case _c == 0:
		
		
		_res := time.Now()
		return MakeTime(_res)

	default:
		PanicArity(_c)
	}
	return NIL
}

var parse_duration_ Proc = func(_args []Object) Object {
	_c := len(_args)
	switch  {
	case _c == 1:
		
		s := ExtractString(_args, 0)
		t, err := time.ParseDuration(s); PanicOnErr(err); _res := int(t)
		return MakeInt(_res)

	default:
		PanicArity(_c)
	}
	return NIL
}

var round_ Proc = func(_args []Object) Object {
	_c := len(_args)
	switch  {
	case _c == 2:
		
		d := ExtractInt(_args, 0)
    m := ExtractInt(_args, 1)
		_res := int(time.Duration(d).Round(time.Duration(m)))
		return MakeInt(_res)

	default:
		PanicArity(_c)
	}
	return NIL
}

var seconds_ Proc = func(_args []Object) Object {
	_c := len(_args)
	switch  {
	case _c == 1:
		
		d := ExtractInt(_args, 0)
		_res := time.Duration(d).Seconds()
		return MakeDouble(_res)

	default:
		PanicArity(_c)
	}
	return NIL
}

var since_ Proc = func(_args []Object) Object {
	_c := len(_args)
	switch  {
	case _c == 1:
		
		t := ExtractTime(_args, 0)
		_res := int(time.Since(t))
		return MakeInt(_res)

	default:
		PanicArity(_c)
	}
	return NIL
}

var sleep_ Proc = func(_args []Object) Object {
	_c := len(_args)
	switch  {
	case _c == 1:
		
		d := ExtractInt(_args, 0)
		_res := sleep(d)
		return _res

	default:
		PanicArity(_c)
	}
	return NIL
}

var string_ Proc = func(_args []Object) Object {
	_c := len(_args)
	switch  {
	case _c == 1:
		
		d := ExtractInt(_args, 0)
		_res := time.Duration(d).String()
		return MakeString(_res)

	default:
		PanicArity(_c)
	}
	return NIL
}

var sub_ Proc = func(_args []Object) Object {
	_c := len(_args)
	switch  {
	case _c == 2:
		
		t := ExtractTime(_args, 0)
    u := ExtractTime(_args, 1)
		_res := int(t.Sub(u))
		return MakeInt(_res)

	default:
		PanicArity(_c)
	}
	return NIL
}

var truncate_ Proc = func(_args []Object) Object {
	_c := len(_args)
	switch  {
	case _c == 2:
		
		d := ExtractInt(_args, 0)
    m := ExtractInt(_args, 1)
		_res := int(time.Duration(d).Truncate(time.Duration(m)))
		return MakeInt(_res)

	default:
		PanicArity(_c)
	}
	return NIL
}

var unix_ Proc = func(_args []Object) Object {
	_c := len(_args)
	switch  {
	case _c == 1:
		
		t := ExtractTime(_args, 0)
		_res := int(t.Unix())
		return MakeInt(_res)

	default:
		PanicArity(_c)
	}
	return NIL
}

var until_ Proc = func(_args []Object) Object {
	_c := len(_args)
	switch  {
	case _c == 1:
		
		t := ExtractTime(_args, 0)
		_res := int(time.Until(t))
		return MakeInt(_res)

	default:
		PanicArity(_c)
	}
	return NIL
}


func init() {

timeNamespace.ResetMeta(MakeMeta(nil, "Provides functionality for measuring and displaying time.", "1.0"))

timeNamespace.InternVar("add", add_,
	MakeMeta(
		NewListFrom(NewVectorFrom(MakeSymbol("t"), MakeSymbol("d"))),
		`Returns the time t+d.`, "1.0"))

timeNamespace.InternVar("format", format_,
	MakeMeta(
		NewListFrom(NewVectorFrom(MakeSymbol("t"), MakeSymbol("layout"))),
		`Returns a textual representation of the time value formatted according to layout,
  which defines the format by showing how the reference time, defined to be
  Mon Jan 2 15:04:05 -0700 MST 2006
  would be displayed if it were the value; it serves as an example of the desired output.
  The same display rules will then be applied to the time value..`, "1.0"))

timeNamespace.InternVar("from-unix", from_unix_,
	MakeMeta(
		NewListFrom(NewVectorFrom(MakeSymbol("sec"), MakeSymbol("nsec"))),
		`Returns the local Time corresponding to the given Unix time, sec seconds and
  nsec nanoseconds since January 1, 1970 UTC. It is valid to pass nsec outside the range [0, 999999999].`, "1.0"))

timeNamespace.InternVar("hours", hours_,
	MakeMeta(
		NewListFrom(NewVectorFrom(MakeSymbol("d"))),
		`Returns the duration (passed as a number of nanoseconds) as a floating point number of hours.`, "1.0"))

timeNamespace.InternVar("minutes", minutes_,
	MakeMeta(
		NewListFrom(NewVectorFrom(MakeSymbol("d"))),
		`Returns the duration (passed as a number of nanoseconds) as a floating point number of minutes.`, "1.0"))

timeNamespace.InternVar("now", now_,
	MakeMeta(
		NewListFrom(NewVectorFrom()),
		`Returns the current local time.`, "1.0"))

timeNamespace.InternVar("parse-duration", parse_duration_,
	MakeMeta(
		NewListFrom(NewVectorFrom(MakeSymbol("s"))),
		`Parses a duration string. A duration string is a possibly signed sequence of decimal numbers,
  each with optional fraction and a unit suffix, such as 300ms, -1.5h or 2h45m. Valid time units are
  ns, us (or µs), ms, s, m, h.`, "1.0"))

timeNamespace.InternVar("round", round_,
	MakeMeta(
		NewListFrom(NewVectorFrom(MakeSymbol("d"), MakeSymbol("m"))),
		`Returns the result of rounding d to the nearest multiple of m. d and m represent time durations in nanoseconds.
  The rounding behavior for halfway values is to round away from zero. If m <= 0, returns d unchanged.`, "1.0"))

timeNamespace.InternVar("seconds", seconds_,
	MakeMeta(
		NewListFrom(NewVectorFrom(MakeSymbol("d"))),
		`Returns the duration (passed as a number of nanoseconds) as a floating point number of seconds.`, "1.0"))

timeNamespace.InternVar("since", since_,
	MakeMeta(
		NewListFrom(NewVectorFrom(MakeSymbol("t"))),
		`Returns the time in nanoseconds elapsed since t.`, "1.0"))

timeNamespace.InternVar("sleep", sleep_,
	MakeMeta(
		NewListFrom(NewVectorFrom(MakeSymbol("d"))),
		`Pauses the execution thread for at least the duration d (expressed in nanoseconds).
  A negative or zero duration causes sleep to return immediately.`, "1.0"))

timeNamespace.InternVar("string", string_,
	MakeMeta(
		NewListFrom(NewVectorFrom(MakeSymbol("d"))),
		`Returns a string representing the duration in the form 72h3m0.5s.`, "1.0"))

timeNamespace.InternVar("sub", sub_,
	MakeMeta(
		NewListFrom(NewVectorFrom(MakeSymbol("t"), MakeSymbol("u"))),
		`Returns the duration t-u in nanoseconds.`, "1.0"))

timeNamespace.InternVar("truncate", truncate_,
	MakeMeta(
		NewListFrom(NewVectorFrom(MakeSymbol("d"), MakeSymbol("m"))),
		`Returns the result of rounding d toward zero to a multiple of m. If m <= 0, returns d unchanged.`, "1.0"))

timeNamespace.InternVar("unix", unix_,
	MakeMeta(
		NewListFrom(NewVectorFrom(MakeSymbol("t"))),
		`Returns t as a Unix time, the number of seconds elapsed since January 1, 1970 UTC.`, "1.0"))

timeNamespace.InternVar("until", until_,
	MakeMeta(
		NewListFrom(NewVectorFrom(MakeSymbol("t"))),
		`Returns the duration in nanoseconds until t.`, "1.0"))

}
