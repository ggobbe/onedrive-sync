package api

import "time"

// Error type
type Error struct {
	Code       string     `json:"code"`
	Message    string     `json:"message"`
	InnerError InnerError `json:"innererror"`
}

// InnerError type
type InnerError struct {
	Code string `json:"code"`
}

// Drive type
type Drive struct {
	ID        string `json:"id"`
	DriveType string `json:"driveType"`
	Owner     Owner  `json:"owner"`
	Quota     Quota  `json:"quota"`
}

// Owner type
type Owner struct {
	User User `json:"user"`
}

// User type
type User struct {
	DisplayName string `json:"displayName"`
	ID          string `json:"id"`
}

// Quota type
type Quota struct {
	Deleted   int    `json:"deleted"`
	Remaining int    `json:"remaining"`
	State     string `json:"state"`
	Total     int    `json:"total"`
	Used      int    `json:"used"`
}

// Children type
type Children struct {
	Value []Item `json:"value"`
}

// Item type
type Item struct {
	ID                   string    `json:"id"`
	Name                 string    `json:"name"`
	ETag                 string    `json:"eTag"`
	CTag                 string    `json:"cTag"`
	CreatedBy            By        `json:"createdBy"`
	CreatedDateTime      time.Time `json:"createdDateTime"`
	LastModifiedBy       By        `json:"lastModifiedBy"`
	LastModifiedDateTime time.Time `json:"lastModifiedDateTime"`
	Size                 int       `json:"size"`
	WebURL               string    `json:"webUrl"`
	ParentReference      Parent    `json:"parentReference"`
	Folder               Folder    `json:"folder"`
}

// By type
type By struct {
	Application User `json:"application"`
	User        User `json:"user"`
}

// Parent type
type Parent struct {
	DriveID string `json:"driveId"`
	ID      string `json:"id"`
	Path    string `json:"path"`
}

// Folder type
type Folder struct {
	ChildCount int `json:"childCount"`
}
