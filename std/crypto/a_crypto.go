// This file is generated by generate-std.joke script. Do not edit manually!

package crypto

import (
	. "github.com/candid82/joker/core"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
)

var cryptoNamespace = GLOBAL_ENV.EnsureNamespace(MakeSymbol("joker.crypto"))



var hmac_ Proc = func(_args []Object) Object {
	_c := len(_args)
	switch {
	case _c == 3:
		algorithm := ExtractKeyword(_args, 0)
		message := ExtractString(_args, 1)
		key := ExtractString(_args, 2)
		_res := hmacSum(algorithm, message, key)
		return MakeString(_res)

	default:
		PanicArity(_c)
	}
	return NIL
}

var md5_ Proc = func(_args []Object) Object {
	_c := len(_args)
	switch {
	case _c == 1:
		data := ExtractString(_args, 0)
		 t := md5.Sum([]byte(data))
		_res := string(t[:])
		return MakeString(_res)

	default:
		PanicArity(_c)
	}
	return NIL
}

var sha1_ Proc = func(_args []Object) Object {
	_c := len(_args)
	switch {
	case _c == 1:
		data := ExtractString(_args, 0)
		 t := sha1.Sum([]byte(data))
		_res := string(t[:])
		return MakeString(_res)

	default:
		PanicArity(_c)
	}
	return NIL
}

var sha224_ Proc = func(_args []Object) Object {
	_c := len(_args)
	switch {
	case _c == 1:
		data := ExtractString(_args, 0)
		 t := sha256.Sum224([]byte(data))
		_res := string(t[:])
		return MakeString(_res)

	default:
		PanicArity(_c)
	}
	return NIL
}

var sha256_ Proc = func(_args []Object) Object {
	_c := len(_args)
	switch {
	case _c == 1:
		data := ExtractString(_args, 0)
		 t := sha256.Sum256([]byte(data))
		_res := string(t[:])
		return MakeString(_res)

	default:
		PanicArity(_c)
	}
	return NIL
}

var sha384_ Proc = func(_args []Object) Object {
	_c := len(_args)
	switch {
	case _c == 1:
		data := ExtractString(_args, 0)
		 t := sha512.Sum384([]byte(data))
		_res := string(t[:])
		return MakeString(_res)

	default:
		PanicArity(_c)
	}
	return NIL
}

var sha512_ Proc = func(_args []Object) Object {
	_c := len(_args)
	switch {
	case _c == 1:
		data := ExtractString(_args, 0)
		 t := sha512.Sum512([]byte(data))
		_res := string(t[:])
		return MakeString(_res)

	default:
		PanicArity(_c)
	}
	return NIL
}

var sha512_224_ Proc = func(_args []Object) Object {
	_c := len(_args)
	switch {
	case _c == 1:
		data := ExtractString(_args, 0)
		 t := sha512.Sum512_224([]byte(data))
		_res := string(t[:])
		return MakeString(_res)

	default:
		PanicArity(_c)
	}
	return NIL
}

var sha512_256_ Proc = func(_args []Object) Object {
	_c := len(_args)
	switch {
	case _c == 1:
		data := ExtractString(_args, 0)
		 t := sha512.Sum512_256([]byte(data))
		_res := string(t[:])
		return MakeString(_res)

	default:
		PanicArity(_c)
	}
	return NIL
}

func init() {

	cryptoNamespace.ResetMeta(MakeMeta(nil, "Implements common cryptographic and hash functions.", "1.0"))

	
	cryptoNamespace.InternVar("hmac", hmac_,
		MakeMeta(
			NewListFrom(NewVectorFrom(MakeSymbol("algorithm"), MakeSymbol("message"), MakeSymbol("key"))),
			`Returns HMAC signature for message and key using specified algorithm.
  Algorithm is one of the following: :sha1, :sha224, :sha256, :sha384, :sha512.`, "1.0"))

	cryptoNamespace.InternVar("md5", md5_,
		MakeMeta(
			NewListFrom(NewVectorFrom(MakeSymbol("data"))),
			`Returns the MD5 checksum of the data.`, "1.0"))

	cryptoNamespace.InternVar("sha1", sha1_,
		MakeMeta(
			NewListFrom(NewVectorFrom(MakeSymbol("data"))),
			`Returns the SHA1 checksum of the data.`, "1.0"))

	cryptoNamespace.InternVar("sha224", sha224_,
		MakeMeta(
			NewListFrom(NewVectorFrom(MakeSymbol("data"))),
			`Returns the SHA224 checksum of the data.`, "1.0"))

	cryptoNamespace.InternVar("sha256", sha256_,
		MakeMeta(
			NewListFrom(NewVectorFrom(MakeSymbol("data"))),
			`Returns the SHA256 checksum of the data.`, "1.0"))

	cryptoNamespace.InternVar("sha384", sha384_,
		MakeMeta(
			NewListFrom(NewVectorFrom(MakeSymbol("data"))),
			`Returns the SHA384 checksum of the data.`, "1.0"))

	cryptoNamespace.InternVar("sha512", sha512_,
		MakeMeta(
			NewListFrom(NewVectorFrom(MakeSymbol("data"))),
			`Returns the SHA512 checksum of the data.`, "1.0"))

	cryptoNamespace.InternVar("sha512-224", sha512_224_,
		MakeMeta(
			NewListFrom(NewVectorFrom(MakeSymbol("data"))),
			`Returns the SHA512/224 checksum of the data.`, "1.0"))

	cryptoNamespace.InternVar("sha512-256", sha512_256_,
		MakeMeta(
			NewListFrom(NewVectorFrom(MakeSymbol("data"))),
			`Returns the SHA512/256 checksum of the data.`, "1.0"))

}
