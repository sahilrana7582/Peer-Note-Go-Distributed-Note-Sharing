package models

type File struct {
	ID         int      `json:"id"`
	FileName   string   `json:"file_name"`
	CourseCode string   `json:"course_code"`
	Professor  string   `json:"professor"`
	PeerID     int      `json:"peer_id"`
	Keywords   []string `json:"keywords"`
	FilePath   string   `json:"file_path"`
}
