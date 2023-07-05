package portal

type Providers struct {
	client *Client
}

func (p Providers) ListProvider() {}

func (p Providers) CreateProvider() {}

func (p Providers) GetProviderByID() {}

func (p Providers) SynchronizeData() {}

func (p Providers) SynchronizeDataById() {}
