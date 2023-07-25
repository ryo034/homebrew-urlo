package usecase

import (
	"github.com/fatih/color"
	"urlo/core/domain"
)

type interactor struct {
	repo      domain.Repository
	presenter OutputPort
}

func NewInteractor(repo domain.Repository, presenter OutputPort) InputPort {
	return &interactor{repo, presenter}
}

func (i *interactor) List(jsonOutput bool, jsonStringOutput bool) error {
	res, err := i.repo.FindAll()
	if err != nil {
		return err
	}
	if res.IsEmpty() {
		i.presenter.NotURLFound()
		return nil
	}

	if jsonOutput {
		return i.presenter.ListAsJson(res)
	}

	if jsonStringOutput {
		return i.presenter.ListAsJsonString(res)
	}

	i.presenter.List(res)
	return nil
}

func (i *interactor) Add(v domain.UrlMap, override bool) error {
	rs, err := i.repo.FindAll()
	if err != nil {
		return err
	}
	if override && rs.IsAlreadyExist(v.Title()) {
		return i.Update(rs.Update(v))
	}

	if rs.IsAlreadyExist(v.Title()) {
		i.presenter.AlreadyExist(v)
		return nil
	}
	if err = i.repo.Add(v); err != nil {
		return err
	}
	i.presenter.AddSuccessfully(v)
	return nil
}

func (i *interactor) BulkAdd(vs *domain.UrlMaps) error {
	rs, err := i.repo.FindAll()
	if err != nil {
		return err
	}
	res, err := rs.AddAll(vs)
	if err != nil {
		return err
	}
	if err = i.repo.Update(res); err != nil {
		return err
	}
	color.Green("Successfully add all the new URL map.")
	return nil
}

func (i *interactor) Update(vs *domain.UrlMaps) error {
	if err := i.repo.Update(vs); err != nil {
		return err
	}
	i.presenter.UpdateSuccessfully()
	return nil
}

func (i *interactor) Clear() error {
	return i.Clear()
}

func (i *interactor) Set(vs *domain.UrlMaps) error {
	ok, err := i.repo.AskOne("Are you sure you want to overwrite the existing data?")
	if err != nil {
		return err
	}
	if !ok {
		i.presenter.Cancelled()
		return nil
	}
	if err := i.repo.Update(vs); err != nil {
		return err
	}
	i.presenter.SetSuccessfully()
	return nil
}

func (i *interactor) Open(t domain.Title) error {
	target, err := i.repo.Find(t)
	if err != nil {
		return err
	}
	return i.repo.OpenUrl(target.URL())
}

func (i *interactor) Delete() error {
	vs, err := i.repo.FindAll()
	if err != nil {
		return err
	}
	o, err := i.repo.SelectOne(vs)
	if err != nil {
		return err
	}
	t, idx, err := vs.GetItemFromLabel(o)
	if err != nil {
		return err
	}
	if err = i.repo.Update(vs.Delete(idx)); err != nil {
		return err
	}
	i.presenter.DeleteSuccessfully(t)
	return nil
}

func (i *interactor) Select(query string) error {
	vs, err := i.repo.Search(query)
	if err != nil {
		return err
	}
	if vs.IsEmpty() {
		i.presenter.NotFoundRecords()
		return nil
	}
	o, err := i.repo.SelectOne(vs)
	if err != nil {
		return err
	}
	s, _, err := vs.GetItemFromLabel(o)
	if err != nil {
		return err
	}
	if err = i.repo.OpenUrl(s.URL()); err != nil {
		return err
	}
	color.Green("Selected %s\n", o)
	return nil
}
