package urlshortener

import (
	"net/http"
)

func utilHandler(urls map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		//	Match path and redirect or call fallback
		path := r.URL.Path
		if dest, ok := urls[path]; ok {
			http.Redirect(rw, r, dest, http.StatusFound)
			return
		} else {
			fallback.ServeHTTP(rw, r)
		}
	}
}

func FileHandler(d []byte, ext string, fallback http.Handler) (http.HandlerFunc, error) {
	switch ext {
	case "json":
		urls, err := parseJSON(d)
		if err != nil {
			panic(err)
		}
		return utilHandler(urls, fallback), err
	case "yml":
		urls, err := parseYAML(d)
		if err != nil {
			panic(err)
		}
		return utilHandler(urls, fallback), err
	default:
		break
	}
	return nil, nil
}

// MapHandler will return an http.HandlerFunc (which also
// implements http.Handler) that will attempt to map any
// paths (keys in the map) to their corresponding URL (values
// that each key in the map points to, in string format).
// If the path is not provided in the map, then the fallback
// http.Handler will be called instead.
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return utilHandler(pathsToUrls, fallback)
}
