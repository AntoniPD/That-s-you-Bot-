package models

type Payload struct {
	URL               string   `json:"url,omitempty"`
	TemplateType      string   `json:"template_type,omitempty"`
	Sharable          bool     `json:"sharable,omitempty"`
	ImageAspectRation string   `json:"image_aspect_ratio,omitempty"`
	Text              string   `json:"text,omitempty"`
	Buttons           []Button `json:"buttons,omitempty"`
}
