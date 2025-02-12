package usecase

import (
	"context"
	"ticketing/payments/internal/common/exception"
	"ticketing/payments/internal/domain"
	"ticketing/payments/internal/infrastructure"
	"ticketing/payments/internal/model"
	"ticketing/payments/internal/publisher"
	"ticketing/payments/internal/repository"

	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type PaymentUsecase interface {
	Create(ctx context.Context, request *model.PaymentRequest) (*model.PaymentResponse, error)
}

type PaymentUsecaseImpl struct {
	PaymentRepository repository.PaymentRepository
	PaymentPublisher  publisher.PaymentPublisher
	PaymentGateway    infrastructure.PaymentGateway
	OrderRepository   repository.OrderRepository
	Logger            *logrus.Logger
	Validate          *validator.Validate
	Config            *viper.Viper
}

func NewPaymentUsecase(paymentRepo repository.PaymentRepository, paymentPublisher publisher.PaymentPublisher, paymentGateway infrastructure.PaymentGateway,
	orderRepo repository.OrderRepository, log *logrus.Logger, validate *validator.Validate, config *viper.Viper) PaymentUsecase {
	return &PaymentUsecaseImpl{
		PaymentRepository: paymentRepo,
		PaymentPublisher:  paymentPublisher,
		PaymentGateway:    paymentGateway,
		OrderRepository:   orderRepo,
		Logger:            log,
		Validate:          validate,
		Config:            config,
	}
}

func (uc *PaymentUsecaseImpl) Create(ctx context.Context, request *model.PaymentRequest) (*model.PaymentResponse, error) {
	if err := uc.Validate.Struct(request); err != nil {
		uc.Logger.WithError(err).Error("failed validating request body")
		return nil, err
	}

	order, err := uc.OrderRepository.FindById(ctx, request.OrderID)
	if err != nil {
		uc.Logger.WithError(err).Error("failed find order by id")
		return nil, exception.ErrOrderNotFound
	}

	if order.UserID != request.UserID {
		uc.Logger.WithError(err).Error("user unauthorized to fetch")
		return nil, exception.ErrUserUnauthorized
	}

	if order.Status == domain.Cancelled {
		uc.Logger.WithError(err).Error("order is cancelled")
		return nil, exception.ErrOrderNotFound
	}

	createPayment, err := uc.PaymentGateway.CreatePayment(order)
	if err != nil {
		uc.Logger.WithError(err).Error("failed create payment intent")
		return nil, exception.ErrPaymentFailed
	}

	payment := new(domain.Payment)
	payment.OrderID = request.OrderID
	payment.StripeID = createPayment.ClientSecret
	if err := uc.PaymentRepository.Create(ctx, payment); err != nil {
		uc.Logger.WithError(err).Error("failed create payment to database")
		return nil, exception.ErrInternalServerError
	}

	if err := uc.PaymentPublisher.Created(payment); err != nil {
		uc.Logger.WithError(err).Error("failed publish event PaymentCreated event")
		return nil, exception.ErrMessageNotPublished
	}

	response := new(model.PaymentResponse)
	response.ID = payment.StripeID
	return response, nil
}
