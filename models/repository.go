package models

type Repository struct {
	Name              string `json:"name"`
	NameWithNamespace string `json:"name_with_namespace"`
	SshUrlToRepo      string `json:"ssh_url_to_repo"`
	HttpUrlToRepo     string `json:"http_url_to_repo"`
	WebUrl            string `json:"web_url"`
	PathWithNamespace string `json:"path_with_namespace"`
}