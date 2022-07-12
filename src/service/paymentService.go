package service

import "github.com/tokopedia/test/bookmyshow/src/repo"

type PaymentService struct {
	paymentRepo repo.PaymentRepo
	bookingRepo repo.BookingRepo
}

func NewPaymentService(paymentRepo repo.PaymentRepo, bookingRepo repo.BookingRepo) PaymentService {
	return PaymentService{
		paymentRepo: paymentRepo,
		bookingRepo: bookingRepo,
	}
}

func (p PaymentService) PaymentFailed(bookingID, user string) {
	p.paymentRepo.ProcessPaymentFailed(p.bookingRepo.GetBooking(bookingID), user)
}

func (p PaymentService) PaymentSuccess(bookingID, user string) {
	p.bookingRepo.ConfirmBooking(p.bookingRepo.GetBooking(bookingID), user)
}
