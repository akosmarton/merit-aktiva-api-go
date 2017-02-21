package aktiva

type Tax struct {
	Id   string
	Code string
	Name string
}

func (a *Aktiva) GetTaxes() ([]Tax, error) {
	var taxes []Tax

	if err := a.httpGet(epGetTaxes, &taxes); err != nil {
		return nil, err
	}
	return taxes, nil
}
