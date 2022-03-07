package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"otp_service/graph/generated"
	"otp_service/graph/model"
	"otp_service/service"
)

func (r *mutationResolver) SendOtp(ctx context.Context, input model.SendOtp) (*model.Status, error) {
	ok, message := service.SendOtp(input.Channel, input.OtpType)
	var Status model.Status
	Status.Status = ok
	Status.Message = message
	return &Status, nil
}

func (r *mutationResolver) VerifyOtp(ctx context.Context, input model.VerifyOtp) (*model.Status, error) {
	ok, message := service.VerifyOtp(input.Channel, input.Otp, input.OtpType)
	var Status model.Status
	Status.Status = ok
	Status.Message = message
	return &Status, nil
}

func (r *mutationResolver) AddOtpType(ctx context.Context, input model.AddOtpType) (*model.Status, error) {
	ok, message := service.AddOtpType(input)
	var Status model.Status
	Status.Status = ok
	Status.Message = message
	return &Status, nil
}

func (r *mutationResolver) RemoveOtpType(ctx context.Context, input model.RemoveOtpType) (*model.Status, error) {
	ok, message := service.RemoveOtpType(input)
	var Status model.Status
	Status.Status = ok
	Status.Message = message
	return &Status, nil
}

func (r *queryResolver) GetOtpType(ctx context.Context) ([]*model.OtpType, error) {
	return service.GetOtpType(), nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
