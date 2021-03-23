// Code generated by "esc -o data_gen.go -pkg data ../../sql"; DO NOT EDIT.

package data

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"sync"
	"time"
)

type _escLocalFS struct{}

var _escLocal _escLocalFS

type _escStaticFS struct{}

var _escStatic _escStaticFS

type _escDirectory struct {
	fs   http.FileSystem
	name string
}

type _escFile struct {
	compressed string
	size       int64
	modtime    int64
	local      string
	isDir      bool

	once sync.Once
	data []byte
	name string
}

func (_escLocalFS) Open(name string) (http.File, error) {
	f, present := _escData[path.Clean(name)]
	if !present {
		return nil, os.ErrNotExist
	}
	return os.Open(f.local)
}

func (_escStaticFS) prepare(name string) (*_escFile, error) {
	f, present := _escData[path.Clean(name)]
	if !present {
		return nil, os.ErrNotExist
	}
	var err error
	f.once.Do(func() {
		f.name = path.Base(name)
		if f.size == 0 {
			return
		}
		var gr *gzip.Reader
		b64 := base64.NewDecoder(base64.StdEncoding, bytes.NewBufferString(f.compressed))
		gr, err = gzip.NewReader(b64)
		if err != nil {
			return
		}
		f.data, err = ioutil.ReadAll(gr)
	})
	if err != nil {
		return nil, err
	}
	return f, nil
}

func (fs _escStaticFS) Open(name string) (http.File, error) {
	f, err := fs.prepare(name)
	if err != nil {
		return nil, err
	}
	return f.File()
}

func (dir _escDirectory) Open(name string) (http.File, error) {
	return dir.fs.Open(dir.name + name)
}

func (f *_escFile) File() (http.File, error) {
	type httpFile struct {
		*bytes.Reader
		*_escFile
	}
	return &httpFile{
		Reader:   bytes.NewReader(f.data),
		_escFile: f,
	}, nil
}

func (f *_escFile) Close() error {
	return nil
}

func (f *_escFile) Readdir(count int) ([]os.FileInfo, error) {
	if !f.isDir {
		return nil, fmt.Errorf(" escFile.Readdir: '%s' is not directory", f.name)
	}

	fis, ok := _escDirs[f.local]
	if !ok {
		return nil, fmt.Errorf(" escFile.Readdir: '%s' is directory, but we have no info about content of this dir, local=%s", f.name, f.local)
	}
	limit := count
	if count <= 0 || limit > len(fis) {
		limit = len(fis)
	}

	if len(fis) == 0 && count > 0 {
		return nil, io.EOF
	}

	return fis[0:limit], nil
}

func (f *_escFile) Stat() (os.FileInfo, error) {
	return f, nil
}

func (f *_escFile) Name() string {
	return f.name
}

func (f *_escFile) Size() int64 {
	return f.size
}

func (f *_escFile) Mode() os.FileMode {
	return 0
}

func (f *_escFile) ModTime() time.Time {
	return time.Unix(f.modtime, 0)
}

func (f *_escFile) IsDir() bool {
	return f.isDir
}

func (f *_escFile) Sys() interface{} {
	return f
}

// FS returns a http.Filesystem for the embedded assets. If useLocal is true,
// the filesystem's contents are instead used.
func FS(useLocal bool) http.FileSystem {
	if useLocal {
		return _escLocal
	}
	return _escStatic
}

// Dir returns a http.Filesystem for the embedded assets on a given prefix dir.
// If useLocal is true, the filesystem's contents are instead used.
func Dir(useLocal bool, name string) http.FileSystem {
	if useLocal {
		return _escDirectory{fs: _escLocal, name: name}
	}
	return _escDirectory{fs: _escStatic, name: name}
}

// FSByte returns the named file from the embedded assets. If useLocal is
// true, the filesystem's contents are instead used.
func FSByte(useLocal bool, name string) ([]byte, error) {
	if useLocal {
		f, err := _escLocal.Open(name)
		if err != nil {
			return nil, err
		}
		b, err := ioutil.ReadAll(f)
		_ = f.Close()
		return b, err
	}
	f, err := _escStatic.prepare(name)
	if err != nil {
		return nil, err
	}
	return f.data, nil
}

// FSMustByte is the same as FSByte, but panics if name is not present.
func FSMustByte(useLocal bool, name string) []byte {
	b, err := FSByte(useLocal, name)
	if err != nil {
		panic(err)
	}
	return b
}

// FSString is the string version of FSByte.
func FSString(useLocal bool, name string) (string, error) {
	b, err := FSByte(useLocal, name)
	return string(b), err
}

// FSMustString is the string version of FSMustByte.
func FSMustString(useLocal bool, name string) string {
	return string(FSMustByte(useLocal, name))
}

var _escData = map[string]*_escFile{

	"/sql/migration-scripts/1-add-reply-field.rb": {
		name:    "1-add-reply-field.rb",
		local:   "../../sql/migration-scripts/1-add-reply-field.rb",
		size:    194,
		modtime: 1606512970,
		compressed: `
H4sIAAAAAAAC/zzKwUrEMBCA4fs8xYhIDkLzAOIha4sUsgq1nkOaDiWwTtZJIhbru0t78PTzw3d7o2sW
PUXWxF8odVoBhD5rFEJ1XdT/ODdHce5e6Y+4iC8xsaslXnIjkwIIiRkfcQ+FcmxTxHP2Yac4J9xC4g0Q
cVcNfVPAux8wduwGHM3JdnimnP1C+UCmbfHp1b6fX3Cg62V1Y3J9i6f++a0bemMf4BeI578AAAD//xsW
xVDCAAAA
`,
	},

	"/sql/migration-scripts/10-permissions-to-jsonb.rb": {
		name:    "10-permissions-to-jsonb.rb",
		local:   "../../sql/migration-scripts/10-permissions-to-jsonb.rb",
		size:    260,
		modtime: 1608436544,
		compressed: `
H4sIAAAAAAAC/3zKsUrEQBCH8X6e4i8iWwjJAxwW0cRqTY4jqUSWZDMcK96szmZF8Xx3SQrBxurjg9/l
RZmTllOQkuUdmqdPIuW3HJRhXo/md5ybgzp3bcpTOOq4hCguL+ElFToZIh9FcIM17Jdti0VHSaNfKeaI
s49yJgCrKviDPa6+iCrbNwf01a1tsGc9hZRClLTB+tDtcdfZ4aFFG2dOu/99Vdd/OJ5TlAlt16MdrEXd
3FeD7WEen8yO6JtYZvoJAAD//7DZqnIEAQAA
`,
	},

	"/sql/migration-scripts/11-channel-kinds.rb": {
		name:    "11-channel-kinds.rb",
		local:   "../../sql/migration-scripts/11-channel-kinds.rb",
		size:    282,
		modtime: 1608614055,
		compressed: `
H4sIAAAAAAAC/3TNz0rDQBDH8fs8xU9E9iAkDxByiM2CYrQlbtVbSDZDu1Anun+KYn13SUU99TR8mc8w
52d5Cj4fnOQse/g0fBB5fkvOM9TrRv1F143Od92lyl/cxvfRTdKl6HYh84MispMISsyDbTxmFn0vobcz
xTjhYCc5EIBZZfzOFhefVDVGtzDVVaOx2PYivAtHVNU1bp2MMPrZFETrVV2Zf4MHbX72JdQ220/OssLT
tW41bsLjnCgRfeKCTj+p2+UKi2Wzvrv/vSroi1jG7wAAAP//6QalxhoBAAA=
`,
	},

	"/sql/migration-scripts/12-nullable.rb": {
		name:    "12-nullable.rb",
		local:   "../../sql/migration-scripts/12-nullable.rb",
		size:    195,
		modtime: 1608614055,
		compressed: `
H4sIAAAAAAAC/zzKwUrDQBDG8fs8xSciOQjZBxAP1fYgrEkJ6XlJNkMZaGd1dlcU67tLFHr6+H/8bm9c
zeZmUcf6AavzF5HxexVjNG/H5hohLGIh3DfuLEebiiQNtcgptzY3RDGp4hHrcCx/2RabNE9xpVgSLjHp
hQCsquVPjrj7po0fdwPGzZPfYc92lpwlacb//9z7w2uHIZ04vGyxHfo9un5Ed/D+gX6IdfkNAAD//530
7RHDAAAA
`,
	},

	"/sql/migration-scripts/13-metadata.rb": {
		name:    "13-metadata.rb",
		local:   "../../sql/migration-scripts/13-metadata.rb",
		size:    298,
		modtime: 1609176019,
		compressed: `
H4sIAAAAAAAC/4TMwUrEMBSF4X2e4ohIF0L7AOKiTgcXto4M4zrcJpcaqDd6k4iD9d2lLgRBmdXhh49z
ftaUpM0YpGF5g5bxaIzyawnKqF6m6ies9UGtvaya5zAp5RDFlhzmVOtYGeOiCK6xDrv8nXVWkkRupfAR
i4uyGACrqvmdHS4+TNsftnsc2pt+i9sSZp/Qdh0GzuQpE8ZjZroyv9zmiUR4Tuj2uwdsdv3jcI+7IP4/
d/Jx4JRo4r/lp2HxXwEAAP//MaWlgCoBAAA=
`,
	},

	"/sql/migration-scripts/14-is-bot.rb": {
		name:    "14-is-bot.rb",
		local:   "../../sql/migration-scripts/14-is-bot.rb",
		size:    176,
		modtime: 1609285553,
		compressed: `
H4sIAAAAAAAC/zzKwYrCMBCA4fs8xSzLksNC8wDiIaU9CMWK9B7adCyBOtFJIor13aUePP388P3+6BxF
D5418Q0lDw8AoWv2Qqguk/qOtaMXa/+VPvtJ+uQD25z8HAsZFIALzLjFNeTSZ4skPcferRTHgIsLvAAi
rqqgOzn8e4JpuvqInSmbGg8STn6miKaqcBdtGRKWbdvUZr+BFxCP7wAAAP//NZ6Ee7AAAAA=
`,
	},

	"/sql/migration-scripts/2-reply-field-default.rb": {
		name:    "2-reply-field-default.rb",
		local:   "../../sql/migration-scripts/2-reply-field-default.rb",
		size:    208,
		modtime: 1606512970,
		compressed: `
H4sIAAAAAAAC/zyKQUsDMRBG7/MrRkRyEDbexcNqIwipQk3PIZsdlkCd1EkiFut/l6zQ7/LxeO/6Srci
ekqsib9Q2nQCEPpsSQjVcVEX8H5O4v2t0h9pkVBTZt9qOpRBJgUQMzM+YD+KdcWhSuASYk9xzniOmc+A
iL0a6Jsi3vzAaJ3ZoRsfrcEtlRIWKmv0L57e7H77ijs6Hk7eZf+yWWXfu3G4Mc/j3jq8u4dfIJ7/AgAA
///tdb9r0AAAAA==
`,
	},

	"/sql/migration-scripts/3-overrides.rb": {
		name:    "3-overrides.rb",
		local:   "../../sql/migration-scripts/3-overrides.rb",
		size:    271,
		modtime: 1606512970,
		compressed: `
H4sIAAAAAAAC/2SMTUrEQBBG932KT0R6ISQHkFlE0qAQf5jJ4LJJOkVoidValR4Ux7tLspDBWRUf7726
vCizStlHLokPkNx/GSP0kaMQ7Pto/4b3QxTvr235Fkfp5pjY5zlOWkhvjQmJGRssh8K8zmKWjrULi4oh
4RgSHw2AxSrokwKuvk3VtG6LtrptHB5ItRtJV6mqazwdSCQOpHjVxP2N+Tnr98911f5Ld649STewnKfJ
rujlzm3dCbzf4XHfNMtn4uE3AAD//+1+Kv4PAQAA
`,
	},

	"/sql/migration-scripts/4-overrides-are-bytea.rb": {
		name:    "4-overrides-are-bytea.rb",
		local:   "../../sql/migration-scripts/4-overrides-are-bytea.rb",
		size:    206,
		modtime: 1606512970,
		compressed: `
H4sIAAAAAAAC/zzKwUrEMBSF4f19iiMiXQjtA4iLUYoImY5ou3AV0uRSAnqjN8ng4Pju0gquDj/nu7zo
atZujtKxHKF1PhEpf9aojOZjaf7D2hDV2uume4+LuhKT2FriW251boh8EsEt1mFftmyLOsnOrxQh4eyT
nAnAqlr+Yo+rb9qZsX/GuLszPfacs1s4b+jvuD+YaT/gcGTVGDhjfH3qMZ8KO0wvj8MDhsmYG/ohlvAb
AAD//7q1MaXOAAAA
`,
	},

	"/sql/migration-scripts/5-attachments.rb": {
		name:    "5-attachments.rb",
		local:   "../../sql/migration-scripts/5-attachments.rb",
		size:    208,
		modtime: 1606512970,
		compressed: `
H4sIAAAAAAAC/0zLQUrFMBDG8f2cYkQkC6E5wMNFpG9XUUp3IiFNhxqwE51MpGK9u7SIvNXwZ37f9ZWt
ReyY2BJ/otTxC0DooyYhNO+z+Q/vpyTe3xq7pFmCpsy+anorjYwGIGZmvMP9UNQjG5XAJcSd4pRxi5k3
QMRdNbRSxJtvaPvHJxzcfXdGpxri60Ks5QSuG8793+OBSgkzlWPs2vZSotKqzy8n+AHi6TcAAP//e1iT
6NAAAAA=
`,
	},

	"/sql/migration-scripts/6-roles.rb": {
		name:    "6-roles.rb",
		local:   "../../sql/migration-scripts/6-roles.rb",
		size:    177,
		modtime: 1606512970,
		compressed: `
H4sIAAAAAAAC/zzKv2rDMBCA8f2e4kopGgrWA5QOLjZdPBlvpQj9OYzAOSUnKcTEefdgD5k+Pvi9v+ma
RbvImviKUt0KIHSpUQjVeVavMSZEMeZT6VOcxZaY2NQSl9yIUwA+MeM37iFfjm2KWM7W7xRDws0n3gAR
d9XQjTx+3KEdpn7Eqf0ZevytcQn5IG3X4ZgWyujWQvbv/wseQByeAQAA//8Sa71EsQAAAA==
`,
	},

	"/sql/migration-scripts/7-permissions.rb": {
		name:    "7-permissions.rb",
		local:   "../../sql/migration-scripts/7-permissions.rb",
		size:    181,
		modtime: 1606512970,
		compressed: `
H4sIAAAAAAAC/zzKvUrFQBBA4X6eYkRkCyH7AGIRSbBJIZJ+2Z8hDCSzOrsrBuO7X5LiVocD3+ODbUVt
YLEkP6gt7ABK342V0Hwt5j7OJVbnno3deFFfOYtrldfSaTAAMYvgK56hWK/tqnopPp4UU8YjZjkAEU/V
0S9FfPqDfprHT5z7t2nE98ZrKhfphwE/SDcuhbMUDHsl/wL/QJJuAQAA//+LM9ZvtQAAAA==
`,
	},

	"/sql/migration-scripts/8-revamp-roles-permissions.rb": {
		name:    "8-revamp-roles-permissions.rb",
		local:   "../../sql/migration-scripts/8-revamp-roles-permissions.rb",
		size:    1468,
		modtime: 1606530386,
		compressed: `
H4sIAAAAAAAC/9STz47aMBDG736KqaqKRUXwAKseTBio1eBQx0iLqioKxKKWgt3aSf+o23evEhJIkHab
w156iuz5/LPn+yavX81K72Z7bWbKfAdX7n8R4tS3UjsFo6/H0WWRJJl2SfJ2NDvpo0sLbU1SFjr3U7cf
EXKwxsA7qD7qUNTLaeFS49NDJYXMwuPBmkcCAJVqqn6qA7z5TWgoUYCk8xBhVeo887VkIaINBFG4XXMQ
Nlf+fohyo9xJe6+t8feEBAKpxOYEWwKPJOADi2V8RsJdDahZCVvAnK1iFIyGtZJvw3BSCyrxc3WenhRI
fJA3+4HNrQPGJa5Q3NTeW+0LmEdRiJTf1DbaHNN9rp4oLyOBbMXhA+7grn38GAQuUSAPMG7c6RYjDgsM
USIENA7oApubBFtTsTujmjbHZDzAu2StTnvlXszDM+45xQu23UO1bXdJzXhcSkM4lw56pK2vTao+g4No
3z9pXZtc7emns9tgd+a5zRTQuMmkXlVjeb6C5rn90Y7UvzLuQIcmHHxJjVF5I2FcDvx5bKb8bROfPv9n
4V+b76Gabd8XPAHccvZxi934r4cuo1Dn/4cok5G/AQAA//+c9A5XvAUAAA==
`,
	},

	"/sql/migration-scripts/9-role-positions.rb": {
		name:    "9-role-positions.rb",
		local:   "../../sql/migration-scripts/9-role-positions.rb",
		size:    185,
		modtime: 1608436544,
		compressed: `
H4sIAAAAAAAC/zzMTWrDMBDF8f2cYkopWhSsA5Qu3No75wOjQHbClgcjcEbJSAoJce4e7EVWjz/8eJ8f
OkfRvWdNfEXJ/R1A6JK9EKrzqN5h7eDF2m+lT36ULvnANic/xUJ6BeACM/7iMuTSmkWSjmPnFopDwNkF
ngERF1XQjRx+PaBsTN2iKf+aGtswUVxFWVX4v2sOmy3uQ/TrhamP5gfgCcQDvAIAAP//WPww0LkAAAA=
`,
	},

	"/sql/migration-scripts/migrate.rb": {
		name:    "migrate.rb",
		local:   "../../sql/migration-scripts/migrate.rb",
		size:    391,
		modtime: 1606512970,
		compressed: `
H4sIAAAAAAAC/2yQvW7CMBSFdz/FIQyQUoxYq7ZLxdChqZT+rFZCLsESien1TVUEeffKJGoQqidb50ff
MdNXY5lgTGHZmNlkUdmSM7GuNo3YndecT5T6JvZ4QEliwtW6ehorZTc4C/d4+kjTVfJuPlfp2/NrogBg
Og3ibBlrfSXHWpzJNGXrLQqHU195OsfC2TfiEb10JHWJDbsK42Pvmy9bzB+Hd6u1jv6y/uCFKkTc5AeM
j/2wdjHYbzTn0S3oZ037MPQOwg0NBRcjr8H/JaQC4i5wRh0M1YWine+KO3/iZBv2iEP3zTSKVPD9BgAA
///PI/IbhwEAAA==
`,
	},

	"/sql/migration-scripts/migration_utils.rb": {
		name:    "migration_utils.rb",
		local:   "../../sql/migration-scripts/migration_utils.rb",
		size:    753,
		modtime: 1609747556,
		compressed: `
H4sIAAAAAAAC/4SSz47aMBCH736KqXsgkVBQpZ6Q6KEt/XcABKWXCKHEmQRXZZyOHVi07LuvnAQIrNjl
gmP/5pvPk7x/N6gsD1JNA6QdcJUehGD8X2lG6JVF7/Lw1xrqCfFlOZ+PJ7/Xf8bzxc/pBEbw4aMQGeZQ
oFsrQ7kuBAAAo6uY4NdiOonKhC0G3/Q/jBiTLJBNLvJMGYYCKbswdshWG6ohKRa6Wb2JjNq6EzSW7YZc
tT5WVQijT4APCkt3atEB39ytPvZqZz170Qv8f1hHaglTIt1K9EHuZQiZgWN+PDfLoz1rh0FzDUbnDusC
CTlxGDzCsEV4U7+EpzB8aaIMESpXHzTDhFHnDQRhLDVZl5BCi7xDlqtYfv3czkJ5PMy+Ry0GhllKybae
jtwkvDV0kP2z8dVvWFlkH2w6xXJpPf1eukys3RvOuhWzdu+VKsPuqsKwu5/eGHuV/mGskyvR/QyVoXp4
zwEAAP//TJX4BfECAAA=
`,
	},

	"/sql/migration-scripts/run-migrations.sh": {
		name:    "run-migrations.sh",
		local:   "../../sql/migration-scripts/run-migrations.sh",
		size:    316,
		modtime: 1606512970,
		compressed: `
H4sIAAAAAAAC/6SOQUvDQBCF7/MrXscQWqGOepUU1Ar2pKR4WhbZdRezuN3K1rSI+N8liUbv3uYx7/tm
jibS7rLYkMSnPazZNUTru4f6+qbi4uPqcn37OER1qj+ZDk2IHgrzBlwMC4a+gNsSsFzVFRdTPDnM79FN
LuRkNv5Pd8ZYiPN7SW2MOF+UZyhLvB4cZkzAeHqavXExpJdfti8ohe+MSQU5htadYOSWq1p+AHLb5Onf
X1Fu7Tt6NcsmPGfz5k+ypa8AAAD//2304lU8AQAA
`,
	},

	"/sql/queries/emotes.sql": {
		name:    "emotes.sql",
		local:   "../../sql/queries/emotes.sql",
		size:    1058,
		modtime: 1606512970,
		compressed: `
H4sIAAAAAAAC/5yTT0+DQBDFz91PMQcObUKbtN5oPJCy1TV1MUD1SDYwMY0uWKBRv72B8mcDS6ueupnd
zpvf8N58DomQaMEmQ1EglWmBTyJ6Awu/MCKM+9QLgPHAheouLC9zmJY/IXNMqA5cSDRhn2MWMmdGnu3d
nvowNZYmGCsTjJvZmpBWysF3HEo5dEcDClvPfVSlyMs99SjUenALxpJMbO40amVlpXa/w6JtnYMlRfJN
fLqjm0Btu2gAyESt1k171ZaRDKabMM6pBw8u42BHx9MhwzhUN+Xq64sOSFOtobV/VLiXY9zVoQfPpHhF
lU0HdFbSLV2VsuO4ehekV61SN1QM084B3Rh/scw2S+VV24xznM3TDDFwT73ySznQfuaObyQFPZzj6fBx
KQE6kf9EoXzsfiaYgZUm2Hihsdkvwrb+CQAA//+5o4QAIgQAAA==
`,
	},

	"/sql/queries/federation.sql": {
		name:    "federation.sql",
		local:   "../../sql/queries/federation.sql",
		size:    203,
		modtime: 1601703329,
		compressed: `
H4sIAAAAAAAC/2SOMcuDMBiE9/dX3JBBIQ5+36Z0kDbWgFUw2o4i+hY6mECU0p9fsCCFTrfcPfdEEeww
c4JsmipnR0bCLx5JV0Y1LXTV1sh5Yj+sD2f7rbIg2FKiW9j3+iRRuJl7w/7JPqRrVnbKIBCxhPiTEP9h
SrQ/nXnd5treHRJnmYwq1bHdaYRvHuVNffl1oFuhGoWP8wEiTt8BAAD//4bcy2jLAAAA
`,
	},

	"/sql/queries/guilds.sql": {
		name:    "guilds.sql",
		local:   "../../sql/queries/guilds.sql",
		size:    3044,
		modtime: 1613782906,
		compressed: `
H4sIAAAAAAAC/7xWQW+bMBi98yu+A4ekSiu13XbotEMCJvVETAVmXU8VbawWKcAEpGv//WQwjg0mSVtp
OX4473t+ft+zT08hTzJ2BU7Jkpott+lmDVdFzixMIhRSwIQG0JQrmFggfk3hHrszWQn+5qzUKu0akmRs
V7tJH+ttye7j0G9qU+vX3I9RBBP7fAb2xQzsyxnYX6YQIhqHBJMlnHy3LElzyeoG103qpOUZIR85FE4s
LwxWgql1e41CJFnCD7DPVRSXbZjcLHtlj5aLfEQRvANjvl7HFStpocIMRLtfseyBlRVM+GKuj0Ts730K
AQEnIJ6PHQpuACSg15gs1aYhy4oXxqG8ssj2bqDrbNpHo/2cuCBI8eKF2qd1g/Oc5DnbDP0gPux3hFhk
rPVcUVRpnRa5si6p2VNRvu0qK1Yn66RO9tpmBvbXGdjfRu0T/1nv9sVJCPXiG3dOkdyXFSGqcW1VGyp5
IZXc7ZbXL0ebdvvY01gu+UzTJavlKV1lSf7WGxTZ0WSPIHRRCIs7eTKmEWwmXpvBLgPEKMbYd6PjRlE6
beDlvTzNOmhOjgTb4WGLMeeK76JqVHOTAiLNNA3UhHtHmkR9yBGmCvxRVDtcs+0U4IOmMwnQhZtmMJEp
1nFRpME2dLyi5BA6aPvprPurCZybAROCQvgZYAJP7Z0VEH3ZmdK9D9rSUzJxnN1tWj+3d5Bhtv4vL6fY
5rVAwblygwtSThATOjmZfuBAeEsT5MEjVrnKKR23E//AFcUuek2rutKaod84opFy1Yj6uSyoc9bVxuJi
qjYm26zLl7b9HuE+G0S7QO5CVc8NUTQ202GPUHRVvPQy1XTRSCLKzPd2YGzVv2i8tKxq7bHQKSjBjtDQ
Q9S5Bg+HEYUwuIWA+Hdqo0WSt7kweGktkryCiXyBdO77yAMLVwtOcW1w4CEDchYH7df9+o+voTfj/EHb
sHopNq2Oftf9CwAA///wKuYX5AsAAA==
`,
	},

	"/sql/queries/invites.sql": {
		name:    "invites.sql",
		local:   "../../sql/queries/invites.sql",
		size:    521,
		modtime: 1592590385,
		compressed: `
H4sIAAAAAAAC/5SRQU+DQBCF7/sr3oFDW+kBvVF7aLpj3QShWRY9NljnQAKLYbHqvzcgkNbERC97eJn9
5r03yyVsXnEIza4uT7x7K8oXJRHWlkVKEW0Neu2gJO508gBlT0XLTgDA0z1pGpRuYA0vWAkxQZU9Nlyx
bb9HEPIHH0W2lxtDF6CUDDLHbt09uELwJ7zkkls+Zzf1uxOSIjL0f7fbhvN2aGBgdi2oOCVtoGKTjDjM
euCE8rGvnSueSz50AfypMjHH4ybKKB1+eIEP79qHdyPmQpPJdKziHRbnNpJXtuOesMrt53iHxW+RpgN1
iXp9E0vM+kZxe+kNif4hrLEM5quvAAAA//9iNxiqCQIAAA==
`,
	},

	"/sql/queries/messages.sql": {
		name:    "messages.sql",
		local:   "../../sql/queries/messages.sql",
		size:    1717,
		modtime: 1609176019,
		compressed: `
H4sIAAAAAAAC/5xVTY/aPBC++1fMwQdYeaUXePvFdiWy2KWRIFRJKO0JuWRaIkGyir0t+++rEMcxKmFX
uSDPh2ceHs8zub2FTB5wDF6SLFAp+QthnGdI/CASYQx+EC/BBBT0CADA7CndJxufs5M13cksw721VwoL
a5ibTXKeacx0ZYjDD0xUdfa2Os0zY0wLlBqTjWcSQ3zcP290bsssf2NRpAnWl7WW290BM63qtlomUksC
0CdfvflKRNCjAwZ0yICOGND/GdA3DOhbBvQdg2C57vUZ0PcM6AcGdPBf+TPoQyjiVRj4wQxu7gixbHHc
o0ZLGB5xW+R/FOFiLmIBn8LlwrJG1p9FKBwq4B7ogAB4AXfYK71D460ZLn0jt+8MtanjPeldXlRvFYm5
mMY18+QV3S/XPKt2070Ol/q8VvOg3YsqGB9k9nwdnkPc5HT2+SWiJ8ZqohYgfITJA/7MCyTLkIsQHr67
US6iKZn7Cz+GyUIeXairx0TamTCDXrGw+sK9WDRYIxHXSji9ejm0IklNi/tqGi9z4wykoyV7uRVPpbV2
OCbeHU0j5pfBGLG3o6kTusNx9snLeOw6qZR8EVKTc1Lvpf7D9j/cLKgrLdysDk3qpXelg015XXkTWqd6
53NxTJVWZ7IW3/wojsxXwfgGJ+Ncm6XnH33SKrNtDdaRc3ijcqPf/Q0AAP//3bxVwrUGAAA=
`,
	},

	"/sql/queries/roles.sql": {
		name:    "roles.sql",
		local:   "../../sql/queries/roles.sql",
		size:    3336,
		modtime: 1608614055,
		compressed: `
H4sIAAAAAAAC/7SXS4+bPBSG9/yKs2CRjDKLhHxfpVRdpOAkSMREXDrtasQ0VgcJcAXMaPrvKxvHmOAh
meayss3hvA+vjy+5v4ciyckC7JIkNQloRmBBC2K4OERBBC6OfGCjFYwMEL/1S5rtH11nIkdYRGcAJzlp
ezbNaNl2NzSt6ra7S4tfyVOmxO9oldYpLfjA2Pi29GIUKgDmtI01Z0rbUtpzpf2f0v5faX9qFCBAURxg
F6/h7rNhSFMckpGDKeSN/DQc5KEIwSrwt40rxsMGBUg6Al/AnPKcS+wcXGGDMzXtmtT85RUt+YuwyJPi
jxEiD9kR3Bkn0vuBgwL4+kPapObe0lcVON45ywiJZCGK5DtNqia9yinhVU1LQy8T8XIR7HLiep+gaEy1
Gh2HDvbEFSm77og8isDjluRPpByeiiZGJ7Xc75lKRBXXesV/0ICRrH1Z823utlLN6YQVJpjWGHwMto9X
nmtH4PiA/Wjj4nXna0lOXwmjWJU0Hyq3s75VO50dAzrzuSNlnlZVSgv0llZ19ZDWz/Sltp+ToiBZZ3rR
dzeMwpHoTRuwHQq2bhi6Pg5Bg8XVRTI25oaAY8/rkY7PZxK9drO6MZ4Ye4ewahCvTsJMOXbJGma4pTE9
HJ0r8e99UhOFq7sNKQ/4ZoTpnlQs83ywoI8odIVuDVJ0fDkLyLoISFhzDpNcZtfFEgQnj6ITWB83bnYp
oca7kNT9klL3aJWG5VN26VZR2bE57ZiH9nfsCZjz8dGB15FXFhZP1JxFKsM1i7mr/t4KvwGIZiK0LLqD
4jKcc4t3COc2Rn2saNlTdhHW3cVYBr5w2HMp18eYicvqUIkIKX7LHtJqAq4jxu/wQ2JNwHXEDv8QhvRk
zD9J/g0AAP//WqVuwAgNAAA=
`,
	},

	"/sql/queries/sessions.sql": {
		name:    "sessions.sql",
		local:   "../../sql/queries/sessions.sql",
		size:    478,
		modtime: 1613399927,
		compressed: `
H4sIAAAAAAAC/3yRQU+DQBCF7/sr3oEDRJoUPGhQD0TGSIKtYRc9GkLHSCJsw24iP98UgWI1Pe7byfe+
zKxWaMuGI0g2ptat0oXhLk0Q6ZaFpIzuFQ7RW5qIh3z7NA0a8fpIOU1P3MEJboSYefFuN31F3HMl0o2k
XCHdqO3MgDui/SnyQf2+7kpb69YTL3FWkITrBD6c0Idz6S07JNvj9FhTPCexoqOlJLVADpqj+Vh9iMIT
c1U3rPRv/1MwAPyBu4Y/ubLg3nZlZV3e6+oD751u0Oov1/NwgeAqvF6vvQFwfocDmudl/XgklJEi/HeL
hcrtwPoOAAD//7GHk9DeAQAA
`,
	},

	"/sql/queries/uploads.sql": {
		name:    "uploads.sql",
		local:   "../../sql/queries/uploads.sql",
		size:    429,
		modtime: 1608436544,
		compressed: `
H4sIAAAAAAAC/3zQwWoCMRAG4LPzFP8hhxXiwbYnSw9bd6wL6wq7aXuU0AxU0Fgwh9qnLzFWiBVPOUzm
m/lnNIK3W5mgdG623shCgnU2WEzkWz6obnvuDOrWLBHLexTxWdWVxnTng/iwMocv0WjtVjT69Y8M6a1s
XrlHocYa6k5D3Wuoh+Ej0Xnci4R83M4L9dzw1OQwDY4yDSJNs265SIvQ+5w7xmkbPEGNr/h19XyY2/1n
5p96EharZ+349ZIqnUvEv4ukXhTx1X/sZfwsdiUbCXLl0BU3bBi38/0GAAD//wmjMyatAQAA
`,
	},

	"/sql/queries/users.sql": {
		name:    "users.sql",
		local:   "../../sql/queries/users.sql",
		size:    3293,
		modtime: 1615571641,
		compressed: `
H4sIAAAAAAAC/9xWO2/bOhTe9SvO4MG+cAIkuXfxRQY/aIeFLBmS3LSTwEZsK0CyApFOk39fUHyYlGRD
SNulS4Ackt/jvOSrKziQks5gQ/me0XrxhkqSFzCrDtSLkY+WCYg4uxZ/U7yaegB+9USKVIab6yK4q6uv
eUHlRYHpBOcvhJPaCcWc8CNr4+0IYz+qOvPWUbi1TzwAHAQogg8hDqQmCAMY24+VRrh3NU/ct1rBheeO
mQbh8QFFqOsc7mF087/nmTzOs0wcw4y+0icPBzGKEsBBEirJY4P4ce7vUQzj0c2kBaDYezC0Lo0yBZNt
UCkGlVfALF1U3OaZwuh2CqO7KYz+ncLovzZv4+6Mess5nNhl9UEXrY+sTbKuapp/O0ga0WU2izrs8DxU
JU1jWr/QempJaeXRUIrKLsNg7eNlYpdY3IdV6O13q3mCvBglLhjcuwqu3dMIJfsowMHGfWUbVGN0eX5+
YVRMSFZXTknPfDg9PrS72/fczlbWGMxKcnj7+8yZ7scrp3xOraUmp0ksfI3sAcyDld21InxrEzaTg15z
xpnDtgz3QTL+Z9Ldf5Kmf+vsnzPCqc64Gl/Z5SYhTbubK43MjvLbLqis2AVIdWEA4IZyjWY5liHp1wB3
c2oDxZTLrrkgS10YICumfFHxC1CyHYdAYdYgWe7sVh7oTg3alnKSEU4cOB08zUZqQmf7cP783Cd2nmVJ
tTnmRebnjPfs/OYsbQ5Pq1gGO0t5V7Gc59VB7+PW52bSMmho9TvHpQ5KlycZ5y1qVY3J/vG7syVsqxfa
vHHLblGJwhtx/aV/L3lES02/rqvSyv4K+ShB8Mds25l397gpqwc2wpAKhNEKRbD4bNLVXqyESVZ9jg+S
/50V7/B5Pt7iBNyNKBY5a/anQ4Q+4TiJxx4AgArdNP/0/N4E6E38xJ14vHv8nnNa5IzTrIcLzpFFhNPU
z8ucpwYhxTuLGO/OcApFv5NV4A00LJZ4UWyP/EgK1tdDW1p+sT61dkXVkSqq1byOWB0/aXZfn3T2Ep5U
/wwAAP//peCI4t0MAAA=
`,
	},

	"/sql/schemas/models.sql": {
		name:    "models.sql",
		local:   "../../sql/schemas/models.sql",
		size:    6007,
		modtime: 1616510121,
		compressed: `
H4sIAAAAAAAC/9RYUW+rNhR+76/wYyOl0va8J5I6vdYIyYBot5omywlnrXcBZ9jtbffrJ8AGnOJAmjbT
fcM55tjfdz5/nHgeYi/GKPZmPkZkgYJVjPBXEsUR2kgoJLq+QghVz5Tcohm5i3BIPL+aGGx8f1rF1yFZ
euE9+hXfo2s9eXI1+eXq6sgCvtixlJ6yDM4YT1GMv8ZoE5DfNvhwG0zK76JI0Ow+xt5BcLEKMbkL7D2i
EC9wiIM5bgA3oVWAbrGPY4zmXjT3bvEQoIUogD/kJ0H6IjKgERTPUNTA7HDLkZ2kF76F0Hrzg3FCAgVT
XOQ0EPkODNZq0IfiTBLOr1yPSqvNDmp0XYi/eApjq1nGc5bBEY16z0yxGmb9Q6SYepIoWnq+T4JD7ETS
mVBotlr52AsuLum7J54m1OdSjWSgfuH9tV4LyUtlfY4OhuBWby5BsYQpNhKxt9+X4Z79Nokupel6K4Oi
jkBKLnIjaj2sd9nNd9ohxi97XrtCOeWtki8kVoOqR4hubKvv+RhZB83RPhAt36mnAugm9I/qYPuqgA2h
mLH8CIbTamJxbrJZpBvO2qBDaJ9fvlC0RjuIvJx8LO4q1VykokAkiPEdDt9YU+lz/Va75vkD26bgCo+1
rfOKYJ12TcFkFK90Cdm2bUvO5rdO939oz8DuZtLSaUJj8jQIzrZds/+pYW3a0jNYnfkjy3NITWH00G1c
dVPh+thOrSROvzqi1jlT8CCKV4fQGy+rWuwPrPRgJwZFxq3v1qCGbS5JEDukrR1DJCDR31Lk2x9MyS1O
K1WrrO4ER0Ldq3a03L7U6Hoyrl20rebsjtFsTSeaNm9MLig/kj9z1Xye6lHT9rnO50aCfPOtQbd44W38
GP3UHEbJtylQa7aZdPPzSJYuwsISpGQPDQ166DSr6XvO6bsanXkBTEFCPYVissRR7C3Xh/1pwg9nmJYg
V5CrPjPE2RYSbQq629+pyoI6P62eoSh4aR5Vf6dNBvbpK41Faz6HdfeUYrvHDHIlkYIXhf74s69d/FwL
Os86BjsPpoD6POOK/v7IFaRcKkrW6Jqsbba7onlPzlIeI+3m8/tYnAkFdM1238w5KZ+PXp+VcdfH+qS7
N73S5PJYafU4FjHJtG/0Hroyp4sOC5OB28Vk0d9MOBmZt/vniReQ0NPLeZIEz4DwoUW+uakvG1F937Vl
u2+QJ0evATsXY+Wzo5zuf2KV7dL4dd8bj/i/4PijZkleLz3Ym3xh8rHZbznovyV2A7GINqt2idaEmFDP
Xstly43+FwAA//8fVy0adxcAAA==
`,
	},

	"/sql": {
		name:  "sql",
		local: `../../sql`,
		isDir: true,
	},

	"/sql/migration-scripts": {
		name:  "migration-scripts",
		local: `../../sql/migration-scripts`,
		isDir: true,
	},

	"/sql/queries": {
		name:  "queries",
		local: `../../sql/queries`,
		isDir: true,
	},

	"/sql/schemas": {
		name:  "schemas",
		local: `../../sql/schemas`,
		isDir: true,
	},
}

var _escDirs = map[string][]os.FileInfo{

	"../../sql": {
		_escData["/sql/migration-scripts"],
		_escData["/sql/queries"],
		_escData["/sql/schemas"],
	},

	"../../sql/migration-scripts": {
		_escData["/sql/migration-scripts/1-add-reply-field.rb"],
		_escData["/sql/migration-scripts/10-permissions-to-jsonb.rb"],
		_escData["/sql/migration-scripts/11-channel-kinds.rb"],
		_escData["/sql/migration-scripts/12-nullable.rb"],
		_escData["/sql/migration-scripts/13-metadata.rb"],
		_escData["/sql/migration-scripts/14-is-bot.rb"],
		_escData["/sql/migration-scripts/2-reply-field-default.rb"],
		_escData["/sql/migration-scripts/3-overrides.rb"],
		_escData["/sql/migration-scripts/4-overrides-are-bytea.rb"],
		_escData["/sql/migration-scripts/5-attachments.rb"],
		_escData["/sql/migration-scripts/6-roles.rb"],
		_escData["/sql/migration-scripts/7-permissions.rb"],
		_escData["/sql/migration-scripts/8-revamp-roles-permissions.rb"],
		_escData["/sql/migration-scripts/9-role-positions.rb"],
		_escData["/sql/migration-scripts/migrate.rb"],
		_escData["/sql/migration-scripts/migration_utils.rb"],
		_escData["/sql/migration-scripts/run-migrations.sh"],
	},

	"../../sql/queries": {
		_escData["/sql/queries/emotes.sql"],
		_escData["/sql/queries/federation.sql"],
		_escData["/sql/queries/guilds.sql"],
		_escData["/sql/queries/invites.sql"],
		_escData["/sql/queries/messages.sql"],
		_escData["/sql/queries/roles.sql"],
		_escData["/sql/queries/sessions.sql"],
		_escData["/sql/queries/uploads.sql"],
		_escData["/sql/queries/users.sql"],
	},

	"../../sql/schemas": {
		_escData["/sql/schemas/models.sql"],
	},
}
