package clubhouse

import "net/url"

func buildPath(path string, params ...string) string {
	base, _ := url.Parse(path)

	query := url.Values{}
	for i := 0; i < len(params); i += 2 {
		query.Set(params[i], params[i+1])
	}

	base.RawQuery = query.Encode()

	return base.String()
}
