package portal

type ListOptions struct {
	Page    int `url:"page"`
	PerPage int `url:"per_page"`
}
