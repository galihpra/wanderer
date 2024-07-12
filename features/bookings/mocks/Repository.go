// Code generated by mockery v2.37.1. DO NOT EDIT.

package mocks

import (
	context "context"
	bookings "wanderer/features/bookings"

	echo "github.com/labstack/echo/v4"

	filters "wanderer/helpers/filters"

	mock "github.com/stretchr/testify/mock"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// ChangePaymentMethod provides a mock function with given fields: ctx, code, data
func (_m *Repository) ChangePaymentMethod(ctx context.Context, code int, data bookings.Booking) (*bookings.Payment, error) {
	ret := _m.Called(ctx, code, data)

	var r0 *bookings.Payment
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int, bookings.Booking) (*bookings.Payment, error)); ok {
		return rf(ctx, code, data)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int, bookings.Booking) *bookings.Payment); ok {
		r0 = rf(ctx, code, data)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*bookings.Payment)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int, bookings.Booking) error); ok {
		r1 = rf(ctx, code, data)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Create provides a mock function with given fields: ctx, data
func (_m *Repository) Create(ctx context.Context, data bookings.Booking) (*bookings.Booking, error) {
	ret := _m.Called(ctx, data)

	var r0 *bookings.Booking
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, bookings.Booking) (*bookings.Booking, error)); ok {
		return rf(ctx, data)
	}
	if rf, ok := ret.Get(0).(func(context.Context, bookings.Booking) *bookings.Booking); ok {
		r0 = rf(ctx, data)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*bookings.Booking)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, bookings.Booking) error); ok {
		r1 = rf(ctx, data)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Export provides a mock function with given fields:
func (_m *Repository) Export() ([]bookings.Booking, error) {
	ret := _m.Called()

	var r0 []bookings.Booking
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]bookings.Booking, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []bookings.Booking); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]bookings.Booking)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ExportFileCsv provides a mock function with given fields: c, data
func (_m *Repository) ExportFileCsv(c echo.Context, data []bookings.Booking) error {
	ret := _m.Called(c, data)

	var r0 error
	if rf, ok := ret.Get(0).(func(echo.Context, []bookings.Booking) error); ok {
		r0 = rf(c, data)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ExportFileExcel provides a mock function with given fields: c, data
func (_m *Repository) ExportFileExcel(c echo.Context, data []bookings.Booking) error {
	ret := _m.Called(c, data)

	var r0 error
	if rf, ok := ret.Get(0).(func(echo.Context, []bookings.Booking) error); ok {
		r0 = rf(c, data)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ExportFilePDF provides a mock function with given fields: c, data
func (_m *Repository) ExportFilePDF(c echo.Context, data []bookings.Booking) error {
	ret := _m.Called(c, data)

	var r0 error
	if rf, ok := ret.Get(0).(func(echo.Context, []bookings.Booking) error); ok {
		r0 = rf(c, data)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAll provides a mock function with given fields: ctx, flt
func (_m *Repository) GetAll(ctx context.Context, flt filters.Filter) ([]bookings.Booking, int, error) {
	ret := _m.Called(ctx, flt)

	var r0 []bookings.Booking
	var r1 int
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, filters.Filter) ([]bookings.Booking, int, error)); ok {
		return rf(ctx, flt)
	}
	if rf, ok := ret.Get(0).(func(context.Context, filters.Filter) []bookings.Booking); ok {
		r0 = rf(ctx, flt)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]bookings.Booking)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, filters.Filter) int); ok {
		r1 = rf(ctx, flt)
	} else {
		r1 = ret.Get(1).(int)
	}

	if rf, ok := ret.Get(2).(func(context.Context, filters.Filter) error); ok {
		r2 = rf(ctx, flt)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GetDetail provides a mock function with given fields: ctx, code
func (_m *Repository) GetDetail(ctx context.Context, code int) (*bookings.Booking, error) {
	ret := _m.Called(ctx, code)

	var r0 *bookings.Booking
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int) (*bookings.Booking, error)); ok {
		return rf(ctx, code)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int) *bookings.Booking); ok {
		r0 = rf(ctx, code)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*bookings.Booking)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int) error); ok {
		r1 = rf(ctx, code)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTourById provides a mock function with given fields: ctx, tourId
func (_m *Repository) GetTourById(ctx context.Context, tourId uint) (*bookings.Tour, error) {
	ret := _m.Called(ctx, tourId)

	var r0 *bookings.Tour
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uint) (*bookings.Tour, error)); ok {
		return rf(ctx, tourId)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uint) *bookings.Tour); ok {
		r0 = rf(ctx, tourId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*bookings.Tour)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uint) error); ok {
		r1 = rf(ctx, tourId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateBookingStatus provides a mock function with given fields: ctx, code, status
func (_m *Repository) UpdateBookingStatus(ctx context.Context, code int, status string) error {
	ret := _m.Called(ctx, code, status)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int, string) error); ok {
		r0 = rf(ctx, code, status)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdatePaymentStatus provides a mock function with given fields: ctx, code, bookingStatus, paymentStatus
func (_m *Repository) UpdatePaymentStatus(ctx context.Context, code int, bookingStatus string, paymentStatus string) error {
	ret := _m.Called(ctx, code, bookingStatus, paymentStatus)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int, string, string) error); ok {
		r0 = rf(ctx, code, bookingStatus, paymentStatus)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewRepository creates a new instance of Repository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *Repository {
	mock := &Repository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
