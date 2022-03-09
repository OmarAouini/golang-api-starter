package service

import (
	"errors"
	"reflect"
	"testing"

	"github.com/OmarAouini/golang-api-starter/entities"
	"github.com/OmarAouini/golang-api-starter/mocks"
	"github.com/OmarAouini/golang-api-starter/store"
)

func TestCompanyService_All(t *testing.T) {

	mockStore := new(mocks.CompanyStore)

	type fields struct {
		Store store.CompanyStore
	}
	tests := []struct {
		name    string
		fields  fields
		want    *[]entities.Company
		prepare func(m *mocks.CompanyStore)
		wantErr bool
	}{
		{
			name:   "should return empty list",
			fields: fields{Store: mockStore},
			want:   &[]entities.Company{},
			prepare: func(m *mocks.CompanyStore) {
				m.On("All").Return(&[]entities.Company{}, nil).Once()
			},
			wantErr: false,
		},
		{
			name:   "should return one",
			fields: fields{Store: mockStore},
			want:   &[]entities.Company{{ID: 1}},
			prepare: func(m *mocks.CompanyStore) {
				m.On("All").Return(&[]entities.Company{{ID: 1}}, nil).Once()
			},
			wantErr: false,
		},
		{
			name:   "should return error",
			fields: fields{Store: mockStore},
			want:   nil,
			prepare: func(m *mocks.CompanyStore) {
				m.On("All").Return(nil, errors.New("")).Once()
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			//prepare stubbing
			if tt.prepare != nil {
				tt.prepare(mockStore)
			}

			s := &CompanyService{
				Store: tt.fields.Store,
			}
			got, err := s.All()
			if (err != nil) != tt.wantErr {
				t.Errorf("CompanyService.All() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CompanyService.All() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCompanyService_Get(t *testing.T) {
	type fields struct {
		Store store.CompanyStore
	}
	type args struct {
		id int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *entities.Company
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &CompanyService{
				Store: tt.fields.Store,
			}
			got, err := s.Get(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("CompanyService.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CompanyService.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCompanyService_GetByName(t *testing.T) {
	type fields struct {
		Store store.CompanyStore
	}
	type args struct {
		name string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *entities.Company
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &CompanyService{
				Store: tt.fields.Store,
			}
			got, err := s.GetByName(tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("CompanyService.GetByName() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CompanyService.GetByName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCompanyService_Create(t *testing.T) {
	type fields struct {
		Store store.CompanyStore
	}
	type args struct {
		comp *entities.Company
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &CompanyService{
				Store: tt.fields.Store,
			}
			if err := s.Create(tt.args.comp); (err != nil) != tt.wantErr {
				t.Errorf("CompanyService.Create() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCompanyService_Delete(t *testing.T) {
	type fields struct {
		Store store.CompanyStore
	}
	type args struct {
		id int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &CompanyService{
				Store: tt.fields.Store,
			}
			if err := s.Delete(tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("CompanyService.Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
