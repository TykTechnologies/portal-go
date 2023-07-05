package portal

type Providers struct {
	client *Client
}

func (u Providers) ListProvider() {}

func (u Providers) CreateProvider() {}

func (u Providers) GetProviderByID() {}

func (u Providers) SynchronizeData() {}

func (u Providers) SynchronizeDataById() {}
