package store

import (
	"fmt"

	"github.com/medtune/beta-platform/pkg/store/model"
)

type bioAnalysisStore interface {
	// Pathology spec analysis level
	CreatePathologyAL(*model.PathologyAnalysisLevel) error
	GetPathologyAL(string) (*model.PathologyAnalysisLevel, error)
	GetPathologiesAL() (*[]model.PathologyAnalysisLevel, error)

	// Specs
	CreateSpec(string, string, float64, float64) error
	GetSpec(string) (*model.SpecAnalysisPool, error)
	GetSpecs() (*[]model.SpecAnalysisPool, error)
}

// CreatePathologyAL .
func (s *Store) CreatePathologyAL(m *model.PathologyAnalysisLevel) error {
	if k, err := s.Valid(m); err != nil || !k {
		return err
	}
	if _, err := s.Insert(m); err != nil {
		return err
	}
	return nil
}

// GetPathologyAL .
func (s *Store) GetPathologyAL(name string) (*model.PathologyAnalysisLevel, error) {
	pal := &model.PathologyAnalysisLevel{}
	has, err := s.Where("name = ?", name).Get(pal)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, fmt.Errorf("record doesnt exist")
	}
	return pal, nil
}

// GetPathologiesAL .
func (s *Store) GetPathologiesAL() (*[]model.PathologyAnalysisLevel, error) {
	var pal []model.PathologyAnalysisLevel
	err := s.Find(&pal)
	if err != nil {
		return nil, err
	}
	return &pal, nil
}

// CreateSpec .
func (s *Store) CreateSpec(name string, unit string, min float64, max float64) error {
	m := &model.SpecAnalysisPool{
		Name: name,
		Unit: unit,
		Min:  min,
		Max:  max,
	}
	if k, err := s.Valid(m); err != nil || !k {
		return err
	}
	if _, err := s.Insert(m); err != nil {
		return err
	}
	return nil
}

// GetSpec .
func (s *Store) GetSpec(name string) (*model.SpecAnalysisPool, error) {
	spec := &model.SpecAnalysisPool{}
	has, err := s.Where("name = ?", name).Get(spec)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, fmt.Errorf("record doesnt exist")
	}
	return spec, nil
}

// GetSpecs .
func (s *Store) GetSpecs() (*[]model.SpecAnalysisPool, error) {
	var specs []model.SpecAnalysisPool
	err := s.Find(&specs)
	if err != nil {
		return nil, err
	}
	return &specs, nil
}