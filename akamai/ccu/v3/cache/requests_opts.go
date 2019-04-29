package cache

// URLs specifies a list of URLs to invalidate or delete.
type URLs struct {
	// Objects represents list of the server paths.
	// Purges content for both http and https schemes.
	Objects []string `json:"objects"`
}

// URLsWithHostname specifies a list of server paths to invalidate or delete
// for a common hostname.
type URLsWithHostname struct {
	// Hostname identifies the domain from which the content is purged.
	Hostname string `json:"hostname"`

	// Objects represents list of the server paths.
	// Purges content for both http and https schemes.
	Objects []string `json:"objects"`
}

// CPCodes specifies a list of CP codes to invalidate or delete.
type CPCodes struct {
	// Objects represents list of CP codes.
	Objects []int `json:"objects"`
}


// CacheTags specifies a list of cache tags to invalidate or delete.
type CacheTags struct {
	// Objects represents list of cache tags.
	Objects []string `json:"objects"`
}