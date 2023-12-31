package purell

import (
	"fmt"
	"net/url"
	"testing"
	"unicode"
)

type testCase struct {
	nm     string
	src    string
	flgs   NormalizationFlags
	res    string
	parsed bool
}

var (
	cases = [...]*testCase{
		{
			"LowerScheme",
			"HTTP://www.SRC.ca",
			FlagLowercaseScheme,
			"http://www.SRC.ca",
			false,
		},
		{
			"LowerScheme2",
			"http://www.SRC.ca",
			FlagLowercaseScheme,
			"http://www.SRC.ca",
			false,
		},
		{
			"LowerHost",
			"HTTP://www.SRC.ca/",
			FlagLowercaseHost,
			"http://www.src.ca/", // Since Go1.1, scheme is automatically lowercased
			false,
		},
		{
			"UpperEscapes",
			`http://www.whatever.com/Some%aa%20Special%8Ecases/`,
			FlagUppercaseEscapes,
			"http://www.whatever.com/Some%AA%20Special%8Ecases/",
			false,
		},
		{
			"UnnecessaryEscapes",
			`http://www.toto.com/%41%42%2E%44/%32%33%52%2D/%5f%7E`,
			FlagDecodeUnnecessaryEscapes,
			"http://www.toto.com/AB.D/23R-/_~",
			false,
		},
		{
			"RemoveDefaultPort",
			"HTTP://www.SRC.ca:80/",
			FlagRemoveDefaultPort,
			"http://www.SRC.ca/", // Since Go1.1, scheme is automatically lowercased
			false,
		},
		{
			"RemoveDefaultPort2",
			"HTTP://www.SRC.ca:80",
			FlagRemoveDefaultPort,
			"http://www.SRC.ca", // Since Go1.1, scheme is automatically lowercased
			false,
		},
		{
			"RemoveDefaultPort3",
			"HTTP://www.SRC.ca:8080",
			FlagRemoveDefaultPort,
			"http://www.SRC.ca:8080", // Since Go1.1, scheme is automatically lowercased
			false,
		},
		{
			"Safe",
			"HTTP://www.SRC.ca:80/to%1ato%8b%ee/OKnow%41%42%43%7e",
			FlagsSafe,
			"http://www.src.ca/to%1Ato%8B%EE/OKnowABC~",
			false,
		},
		{
			"BothLower",
			"HTTP://www.SRC.ca:80/to%1ato%8b%ee/OKnow%41%42%43%7e",
			FlagLowercaseHost | FlagLowercaseScheme,
			"http://www.src.ca:80/to%1Ato%8B%EE/OKnowABC~",
			false,
		},
		{
			"RemoveTrailingSlash",
			"HTTP://www.SRC.ca:80/",
			FlagRemoveTrailingSlash,
			"http://www.SRC.ca:80", // Since Go1.1, scheme is automatically lowercased
			false,
		},
		{
			"RemoveTrailingSlash2",
			"HTTP://www.SRC.ca:80/toto/titi/",
			FlagRemoveTrailingSlash,
			"http://www.SRC.ca:80/toto/titi", // Since Go1.1, scheme is automatically lowercased
			false,
		},
		{
			"RemoveTrailingSlash3",
			"HTTP://www.SRC.ca:80/toto/titi/fin/?a=1",
			FlagRemoveTrailingSlash,
			"http://www.SRC.ca:80/toto/titi/fin?a=1", // Since Go1.1, scheme is automatically lowercased
			false,
		},
		{
			"AddTrailingSlash",
			"HTTP://www.SRC.ca:80",
			FlagAddTrailingSlash,
			"http://www.SRC.ca:80/", // Since Go1.1, scheme is automatically lowercased
			false,
		},
		{
			"AddTrailingSlash2",
			"HTTP://www.SRC.ca:80/toto/titi.html",
			FlagAddTrailingSlash,
			"http://www.SRC.ca:80/toto/titi.html/", // Since Go1.1, scheme is automatically lowercased
			false,
		},
		{
			"AddTrailingSlash3",
			"HTTP://www.SRC.ca:80/toto/titi/fin?a=1",
			FlagAddTrailingSlash,
			"http://www.SRC.ca:80/toto/titi/fin/?a=1", // Since Go1.1, scheme is automatically lowercased
			false,
		},
		{
			"RemoveDotSegments",
			"HTTP://root/a/b/./../../c/",
			FlagRemoveDotSegments,
			"http://root/c/", // Since Go1.1, scheme is automatically lowercased
			false,
		},
		{
			"RemoveDotSegments2",
			"HTTP://root/../a/b/./../c/../d",
			FlagRemoveDotSegments,
			"http://root/a/d", // Since Go1.1, scheme is automatically lowercased
			false,
		},
		{
			"UsuallySafe",
			"HTTP://www.SRC.ca:80/to%1ato%8b%ee/./c/d/../OKnow%41%42%43%7e/?a=b#test",
			FlagsUsuallySafeGreedy,
			"http://www.src.ca/to%1Ato%8B%EE/c/OKnowABC~?a=b#test",
			false,
		},
		{
			"RemoveDirectoryIndex",
			"HTTP://root/a/b/c/default.aspx",
			FlagRemoveDirectoryIndex,
			"http://root/a/b/c/", // Since Go1.1, scheme is automatically lowercased
			false,
		},
		{
			"RemoveDirectoryIndex2",
			"HTTP://root/a/b/c/default#a=b",
			FlagRemoveDirectoryIndex,
			"http://root/a/b/c/default#a=b", // Since Go1.1, scheme is automatically lowercased
			false,
		},
		{
			"RemoveFragment",
			"HTTP://root/a/b/c/default#toto=tata",
			FlagRemoveFragment,
			"http://root/a/b/c/default", // Since Go1.1, scheme is automatically lowercased
			false,
		},
		{
			"ForceHTTP",
			"https://root/a/b/c/default#toto=tata",
			FlagForceHTTP,
			"http://root/a/b/c/default#toto=tata",
			false,
		},
		{
			"RemoveDuplicateSlashes",
			"https://root/a//b///c////default#toto=tata",
			FlagRemoveDuplicateSlashes,
			"https://root/a/b/c/default#toto=tata",
			false,
		},
		{
			"RemoveDuplicateSlashes2",
			"https://root//a//b///c////default#toto=tata",
			FlagRemoveDuplicateSlashes,
			"https://root/a/b/c/default#toto=tata",
			false,
		},
		{
			"RemoveWWW",
			"https://www.root/a/b/c/",
			FlagRemoveWWW,
			"https://root/a/b/c/",
			false,
		},
		{
			"RemoveWWW2",
			"https://WwW.Root/a/b/c/",
			FlagRemoveWWW,
			"https://Root/a/b/c/",
			false,
		},
		{
			"AddWWW",
			"https://Root/a/b/c/",
			FlagAddWWW,
			"https://www.Root/a/b/c/",
			false,
		},
		{
			"SortQuery",
			"http://root/toto/?b=4&a=1&c=3&b=2&a=5",
			FlagSortQuery,
			"http://root/toto/?a=1&a=5&b=2&b=4&c=3",
			false,
		},
		{
			"RemoveEmptyQuerySeparator",
			"http://root/toto/?",
			FlagRemoveEmptyQuerySeparator,
			"http://root/toto/",
			false,
		},
		{
			"Unsafe",
			"HTTPS://www.RooT.com/toto/t%45%1f///a/./b/../c/?z=3&w=2&a=4&w=1#invalid",
			FlagsUnsafeGreedy,
			"http://root.com/toto/tE%1F/a/c?a=4&w=1&w=2&z=3",
			false,
		},
		{
			"Safe2",
			"HTTPS://www.RooT.com/toto/t%45%1f///a/./b/../c/?z=3&w=2&a=4&w=1#invalid",
			FlagsSafe,
			"https://www.root.com/toto/tE%1F///a/./b/../c/?z=3&w=2&a=4&w=1#invalid",
			false,
		},
		{
			"UsuallySafe2",
			"HTTPS://www.RooT.com/toto/t%45%1f///a/./b/../c/?z=3&w=2&a=4&w=1#invalid",
			FlagsUsuallySafeGreedy,
			"https://www.root.com/toto/tE%1F///a/c?z=3&w=2&a=4&w=1#invalid",
			false,
		},
		{
			"AddTrailingSlashBug",
			"http://src.ca/",
			FlagsAllNonGreedy,
			"http://www.src.ca/",
			false,
		},
		{
			"SourceModified",
			"HTTPS://www.RooT.com/toto/t%45%1f///a/./b/../c/?z=3&w=2&a=4&w=1#invalid",
			FlagsUnsafeGreedy,
			"http://root.com/toto/tE%1F/a/c?a=4&w=1&w=2&z=3",
			true,
		},
		{
			"IPv6-1",
			"http://[2001:db8:1f70::999:de8:7648:6e8]/test",
			FlagsSafe | FlagRemoveDotSegments,
			"http://[2001:db8:1f70::999:de8:7648:6e8]/test",
			false,
		},
		{
			"IPv6-2",
			"http://[::ffff:192.168.1.1]/test",
			FlagsSafe | FlagRemoveDotSegments,
			"http://[::ffff:192.168.1.1]/test",
			false,
		},
		{
			"IPv6-3",
			"http://[::ffff:192.168.1.1]:80/test",
			FlagsSafe | FlagRemoveDotSegments,
			"http://[::ffff:192.168.1.1]/test",
			false,
		},
		{
			"IPv6-4",
			"htTps://[::fFff:192.168.1.1]:443/test",
			FlagsSafe | FlagRemoveDotSegments,
			"https://[::ffff:192.168.1.1]/test",
			false,
		},
		{
			"FTP",
			"ftp://user:pass@ftp.foo.net/foo/bar",
			FlagsSafe | FlagRemoveDotSegments,
			"ftp://user:pass@ftp.foo.net/foo/bar",
			false,
		},
		{
			"Standard-1",
			"http://www.foo.com:80/foo",
			FlagsSafe | FlagRemoveDotSegments,
			"http://www.foo.com/foo",
			false,
		},
		{
			"Standard-2",
			"http://www.foo.com:8000/foo",
			FlagsSafe | FlagRemoveDotSegments,
			"http://www.foo.com:8000/foo",
			false,
		},
		{
			"Standard-3",
			"http://www.foo.com/%7ebar",
			FlagsSafe | FlagRemoveDotSegments,
			"http://www.foo.com/~bar",
			false,
		},
		{
			"Standard-4",
			"http://www.foo.com/%7Ebar",
			FlagsSafe | FlagRemoveDotSegments,
			"http://www.foo.com/~bar",
			false,
		},
		{
			"Standard-5",
			"http://USER:pass@www.Example.COM/foo/bar",
			FlagsSafe | FlagRemoveDotSegments,
			"http://USER:pass@www.example.com/foo/bar",
			false,
		},
		{
			"Standard-6",
			"http://test.example/?a=%26&b=1",
			FlagsSafe | FlagRemoveDotSegments,
			"http://test.example/?a=%26&b=1",
			false,
		},
		{
			"Standard-7",
			"http://test.example/%25/?p=%20val%20%25",
			FlagsSafe | FlagRemoveDotSegments,
			"http://test.example/%25/?p=%20val%20%25",
			false,
		},
		{
			"Standard-8",
			"http://test.example/path/with a%20space+/",
			FlagsSafe | FlagRemoveDotSegments,
			"http://test.example/path/with%20a%20space+/",
			false,
		},
		{
			"Standard-9",
			"http://test.example/?",
			FlagsSafe | FlagRemoveDotSegments,
			"http://test.example/",
			false,
		},
		{
			"Standard-10",
			"http://a.COM/path/?b&a",
			FlagsSafe | FlagRemoveDotSegments,
			"http://a.com/path/?b&a",
			false,
		},
		{
			"StandardCasesAddTrailingSlash",
			"http://test.example?",
			FlagsSafe | FlagAddTrailingSlash,
			"http://test.example/",
			false,
		},
		{
			"OctalIP-1",
			"http://0123.011.0.4/",
			FlagsSafe | FlagDecodeOctalHost,
			"http://0123.011.0.4/",
			false,
		},
		{
			"OctalIP-2",
			"http://0102.0146.07.0223/",
			FlagsSafe | FlagDecodeOctalHost,
			"http://66.102.7.147/",
			false,
		},
		{
			"OctalIP-3",
			"http://0102.0146.07.0223.:23/",
			FlagsSafe | FlagDecodeOctalHost,
			"http://66.102.7.147.:23/",
			false,
		},
		{
			"OctalIP-4",
			"http://USER:pass@0102.0146.07.0223../",
			FlagsSafe | FlagDecodeOctalHost,
			"http://USER:pass@66.102.7.147../",
			false,
		},
		{
			"DWORDIP-1",
			"http://123.1113982867/",
			FlagsSafe | FlagDecodeDWORDHost,
			"http://123.1113982867/",
			false,
		},
		{
			"DWORDIP-2",
			"http://1113982867/",
			FlagsSafe | FlagDecodeDWORDHost,
			"http://66.102.7.147/",
			false,
		},
		{
			"DWORDIP-3",
			"http://1113982867.:23/",
			FlagsSafe | FlagDecodeDWORDHost,
			"http://66.102.7.147.:23/",
			false,
		},
		{
			"DWORDIP-4",
			"http://USER:pass@1113982867../",
			FlagsSafe | FlagDecodeDWORDHost,
			"http://USER:pass@66.102.7.147../",
			false,
		},
		{
			"HexIP-1",
			"http://0x123.1113982867/",
			FlagsSafe | FlagDecodeHexHost,
			"http://0x123.1113982867/",
			false,
		},
		{
			"HexIP-2",
			"http://0x42660793/",
			FlagsSafe | FlagDecodeHexHost,
			"http://66.102.7.147/",
			false,
		},
		{
			"HexIP-3",
			"http://0x42660793.:23/",
			FlagsSafe | FlagDecodeHexHost,
			"http://66.102.7.147.:23/",
			false,
		},
		{
			"HexIP-4",
			"http://USER:pass@0x42660793../",
			FlagsSafe | FlagDecodeHexHost,
			"http://USER:pass@66.102.7.147../",
			false,
		},
		{
			"UnnecessaryHostDots-1",
			"http://.www.foo.com../foo/bar.html",
			FlagsSafe | FlagRemoveUnnecessaryHostDots,
			"http://www.foo.com/foo/bar.html",
			false,
		},
		{
			"UnnecessaryHostDots-2",
			"http://www.foo.com./foo/bar.html",
			FlagsSafe | FlagRemoveUnnecessaryHostDots,
			"http://www.foo.com/foo/bar.html",
			false,
		},
		{
			"UnnecessaryHostDots-3",
			"http://www.foo.com.:81/foo",
			FlagsSafe | FlagRemoveUnnecessaryHostDots,
			"http://www.foo.com:81/foo",
			false,
		},
		{
			"UnnecessaryHostDots-4",
			"http://www.example.com./",
			FlagsSafe | FlagRemoveUnnecessaryHostDots,
			"http://www.example.com/",
			false,
		},
		{
			"UnnecessaryHostDots-5",
			"http://www..example...com/",
			FlagsSafe | FlagRemoveUnnecessaryHostDots,
			"http://www.example.com/",
			false,
		},
		{
			"EmptyPort-1",
			"http://www.thedraymin.co.uk:/main/?p=308",
			FlagsSafe | FlagRemoveEmptyPortSeparator,
			"http://www.thedraymin.co.uk/main/?p=308",
			false,
		},
		{
			"EmptyPort-2",
			"http://www.src.ca:",
			FlagsSafe | FlagRemoveEmptyPortSeparator,
			"http://www.src.ca",
			false,
		},
		{
			"Slashes-1",
			"http://test.example/foo/bar/.",
			FlagsSafe | FlagRemoveDotSegments | FlagRemoveDuplicateSlashes,
			"http://test.example/foo/bar/",
			false,
		},
		{
			"Slashes-2",
			"http://test.example/foo/bar/./",
			FlagsSafe | FlagRemoveDotSegments | FlagRemoveDuplicateSlashes,
			"http://test.example/foo/bar/",
			false,
		},
		{
			"Slashes-3",
			"http://test.example/foo/bar/..",
			FlagsSafe | FlagRemoveDotSegments | FlagRemoveDuplicateSlashes,
			"http://test.example/foo/",
			false,
		},
		{
			"Slashes-4",
			"http://test.example/foo/bar/../",
			FlagsSafe | FlagRemoveDotSegments | FlagRemoveDuplicateSlashes,
			"http://test.example/foo/",
			false,
		},
		{
			"Slashes-5",
			"http://test.example/foo/bar/../baz",
			FlagsSafe | FlagRemoveDotSegments | FlagRemoveDuplicateSlashes,
			"http://test.example/foo/baz",
			false,
		},
		{
			"Slashes-6",
			"http://test.example/foo/bar/../..",
			FlagsSafe | FlagRemoveDotSegments | FlagRemoveDuplicateSlashes,
			"http://test.example/",
			false,
		},
		{
			"Slashes-7",
			"http://test.example/foo/bar/../../",
			FlagsSafe | FlagRemoveDotSegments | FlagRemoveDuplicateSlashes,
			"http://test.example/",
			false,
		},
		{
			"Slashes-8",
			"http://test.example/foo/bar/../../baz",
			FlagsSafe | FlagRemoveDotSegments | FlagRemoveDuplicateSlashes,
			"http://test.example/baz",
			false,
		},
		{
			"Slashes-9",
			"http://test.example/foo/bar/../../../baz",
			FlagsSafe | FlagRemoveDotSegments | FlagRemoveDuplicateSlashes,
			"http://test.example/baz",
			false,
		},
		{
			"Slashes-10",
			"http://test.example/foo/bar/../../../../baz",
			FlagsSafe | FlagRemoveDotSegments | FlagRemoveDuplicateSlashes,
			"http://test.example/baz",
			false,
		},
		{
			"Slashes-11",
			"http://test.example/./foo",
			FlagsSafe | FlagRemoveDotSegments | FlagRemoveDuplicateSlashes,
			"http://test.example/foo",
			false,
		},
		{
			"Slashes-12",
			"http://test.example/../foo",
			FlagsSafe | FlagRemoveDotSegments | FlagRemoveDuplicateSlashes,
			"http://test.example/foo",
			false,
		},
		{
			"Slashes-13",
			"http://test.example/foo.",
			FlagsSafe | FlagRemoveDotSegments | FlagRemoveDuplicateSlashes,
			"http://test.example/foo.",
			false,
		},
		{
			"Slashes-14",
			"http://test.example/.foo",
			FlagsSafe | FlagRemoveDotSegments | FlagRemoveDuplicateSlashes,
			"http://test.example/.foo",
			false,
		},
		{
			"Slashes-15",
			"http://test.example/foo..",
			FlagsSafe | FlagRemoveDotSegments | FlagRemoveDuplicateSlashes,
			"http://test.example/foo..",
			false,
		},
		{
			"Slashes-16",
			"http://test.example/..foo",
			FlagsSafe | FlagRemoveDotSegments | FlagRemoveDuplicateSlashes,
			"http://test.example/..foo",
			false,
		},
		{
			"Slashes-17",
			"http://test.example/./../foo",
			FlagsSafe | FlagRemoveDotSegments | FlagRemoveDuplicateSlashes,
			"http://test.example/foo",
			false,
		},
		{
			"Slashes-18",
			"http://test.example/./foo/.",
			FlagsSafe | FlagRemoveDotSegments | FlagRemoveDuplicateSlashes,
			"http://test.example/foo/",
			false,
		},
		{
			"Slashes-19",
			"http://test.example/foo/./bar",
			FlagsSafe | FlagRemoveDotSegments | FlagRemoveDuplicateSlashes,
			"http://test.example/foo/bar",
			false,
		},
		{
			"Slashes-20",
			"http://test.example/foo/../bar",
			FlagsSafe | FlagRemoveDotSegments | FlagRemoveDuplicateSlashes,
			"http://test.example/bar",
			false,
		},
		{
			"Slashes-21",
			"http://test.example/foo//",
			FlagsSafe | FlagRemoveDotSegments | FlagRemoveDuplicateSlashes,
			"http://test.example/foo/",
			false,
		},
		{
			"Slashes-22",
			"http://test.example/foo///bar//",
			FlagsSafe | FlagRemoveDotSegments | FlagRemoveDuplicateSlashes,
			"http://test.example/foo/bar/",
			false,
		},
		{
			"Relative",
			"foo/bar",
			FlagsAllGreedy,
			"foo/bar",
			false,
		},
		{
			"Relative-1",
			"./../foo",
			FlagsSafe | FlagRemoveDotSegments | FlagRemoveDuplicateSlashes,
			"foo",
			false,
		},
		{
			"Relative-2",
			"./foo/bar/../baz/../bang/..",
			FlagsSafe | FlagRemoveDotSegments | FlagRemoveDuplicateSlashes,
			"foo/",
			false,
		},
		{
			"Relative-3",
			"foo///bar//",
			FlagsSafe | FlagRemoveDotSegments | FlagRemoveDuplicateSlashes,
			"foo/bar/",
			false,
		},
		{
			"Relative-4",
			"www.youtube.com",
			FlagsUsuallySafeGreedy,
			"www.youtube.com",
			false,
		},
		{
			"Issue-#24",
			"///foo///bar///",
			FlagRemoveDuplicateSlashes | FlagRemoveTrailingSlash,
			"/foo/bar",
			false,
		},
		/*&testCase{
			"UrlNorm-5",
			"http://ja.wikipedia.org/wiki/%E3%82%AD%E3%83%A3%E3%82%BF%E3%83%94%E3%83%A9%E3%83%BC%E3%82%B8%E3%83%A3%E3%83%91%E3%83%B3",
			FlagsSafe | FlagRemoveDotSegments,
			"http://ja.wikipedia.org/wiki/\xe3\x82\xad\xe3\x83\xa3\xe3\x82\xbf\xe3\x83\x94\xe3\x83\xa9\xe3\x83\xbc\xe3\x82\xb8\xe3\x83\xa3\xe3\x83\x91\xe3\x83\xb3",
			false,
		},
		&testCase{
			"UrlNorm-1",
			"http://test.example/?a=%e3%82%82%26",
			FlagsAllGreedy,
			"http://test.example/?a=\xe3\x82\x82%26",
			false,
		},*/
	}
)

func TestRunner(t *testing.T) {
	for _, tc := range cases {
		runCase(tc, t)
	}
}

func runCase(tc *testCase, t *testing.T) {
	t.Logf("running %s...", tc.nm)
	if tc.parsed {
		u, e := url.Parse(tc.src)
		if e != nil {
			t.Errorf("%s - FAIL : %s", tc.nm, e)
			return
		} else {
			NormalizeURL(u, tc.flgs)
			if s := u.String(); s != tc.res {
				t.Errorf("%s - FAIL expected '%s', got '%s'", tc.nm, tc.res, s)
			}
		}
	} else {
		if s, e := NormalizeURLString(tc.src, tc.flgs); e != nil {
			t.Errorf("%s - FAIL : %s", tc.nm, e)
		} else if s != tc.res {
			t.Errorf("%s - FAIL expected '%s', got '%s'", tc.nm, tc.res, s)
		}
	}
}

func TestDecodeUnnecessaryEscapesAll(t *testing.T) {
	var url = "http://host/"

	for i := 0; i < 256; i++ {
		url += fmt.Sprintf("%%%02x", i)
	}
	s, err := NormalizeURLString(url, FlagDecodeUnnecessaryEscapes)
	if err != nil {
		t.Fatalf("parse error: %s", err)
	}

	const want = "http://host/%00%01%02%03%04%05%06%07%08%09%0A%0B%0C%0D%0E%0F%10%11%12%13%14%15%16%17%18%19%1A%1B%1C%1D%1E%1F%20!%22%23$%25&'()*+,-./0123456789:;%3C=%3E%3F@ABCDEFGHIJKLMNOPQRSTUVWXYZ[%5C]%5E_%60abcdefghijklmnopqrstuvwxyz%7B%7C%7D~%7F%80%81%82%83%84%85%86%87%88%89%8A%8B%8C%8D%8E%8F%90%91%92%93%94%95%96%97%98%99%9A%9B%9C%9D%9E%9F%A0%A1%A2%A3%A4%A5%A6%A7%A8%A9%AA%AB%AC%AD%AE%AF%B0%B1%B2%B3%B4%B5%B6%B7%B8%B9%BA%BB%BC%BD%BE%BF%C0%C1%C2%C3%C4%C5%C6%C7%C8%C9%CA%CB%CC%CD%CE%CF%D0%D1%D2%D3%D4%D5%D6%D7%D8%D9%DA%DB%DC%DD%DE%DF%E0%E1%E2%E3%E4%E5%E6%E7%E8%E9%EA%EB%EC%ED%EE%EF%F0%F1%F2%F3%F4%F5%F6%F7%F8%F9%FA%FB%FC%FD%FE%FF"
	if s != want {
		t.Errorf("DecodeUnnecessaryEscapesAll:\nwant\n%s\ngot\n%s", want, s)
	}
}

func TestEncodeNecessaryEscapesAll(t *testing.T) {
	const base = "http://host/"
	var path []byte

	for i := 0; i < 256; i++ {
		// Since go1.12, url.Parse fails if the raw URL contains ASCII control characters,
		// meaning anything < 0x20 and 0x7f (DEL), so do not add those bytes to the constructed url.
		// See https://github.com/PuerkitoBio/purell/issues/28
		if i != 0x25 && !unicode.IsControl(rune(i)) {
			path = append(path, byte(i))
		}
	}
	s, err := NormalizeURLString(base+string(path), FlagEncodeNecessaryEscapes)
	if err != nil {
		t.Fatalf("parse error: %s", err)
	}

	const want = "http://host/%20!%22#$&'()*+,-./0123456789:;%3C=%3E?@ABCDEFGHIJKLMNOPQRSTUVWXYZ[%5C]%5E_%60abcdefghijklmnopqrstuvwxyz%7B%7C%7D~%A0%A1%A2%A3%A4%A5%A6%A7%A8%A9%AA%AB%AC%AD%AE%AF%B0%B1%B2%B3%B4%B5%B6%B7%B8%B9%BA%BB%BC%BD%BE%BF%C0%C1%C2%C3%C4%C5%C6%C7%C8%C9%CA%CB%CC%CD%CE%CF%D0%D1%D2%D3%D4%D5%D6%D7%D8%D9%DA%DB%DC%DD%DE%DF%E0%E1%E2%E3%E4%E5%E6%E7%E8%E9%EA%EB%EC%ED%EE%EF%F0%F1%F2%F3%F4%F5%F6%F7%F8%F9%FA%FB%FC%FD%FE%FF"
	if s != want {
		t.Errorf("EncodeNecessaryEscapesAll:\nwant\n%s\ngot\n%s", want, s)
	}
}
