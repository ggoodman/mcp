// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package internal

type Integration struct {
	ID           int64  `json:"id"`
	Name         string `json:"name"`
	Description  string `json:"description"`
	Vendor       string `json:"vendor"`
	SourceUrl    string `json:"source_url"`
	Homepage     string `json:"homepage"`
	License      string `json:"license"`
	Instructions []byte `json:"instructions"`
}
