package server

import (
	"context"
	"time"

	"github.com/esc-chula/intania-vote/apps/api/model"
	"github.com/esc-chula/intania-vote/apps/api/service"
	grpcVote "github.com/esc-chula/intania-vote/libs/grpc-go/vote"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type voteServerImpl struct {
	grpcVote.UnimplementedVoteServiceServer
	svc service.VoteService
}

func NewVoteServer(s service.VoteService) grpcVote.VoteServiceServer {
	return voteServerImpl{
		svc: s,
	}
}

func RegisterVoteServiceServer(s grpc.ServiceRegistrar, server grpcVote.VoteServiceServer) {
	grpcVote.RegisterVoteServiceServer(s, server)
}

func (s voteServerImpl) CreateVote(ctx context.Context, req *grpcVote.CreateVoteRequest) (*grpcVote.CreateVoteResponse, error) {
	oidcId := req.GetOidcId()
	if oidcId == "" {
		return nil, status.Error(codes.FailedPrecondition, "missing oidcId")
	}
	vote := req.GetVote()
	if vote == nil {
		return nil, status.Error(codes.FailedPrecondition, "missing vote")
	}
	choices := req.GetChoices()
	if len(choices) == 0 {
		return nil, status.Error(codes.FailedPrecondition, "missing choices")
	}

	voteImage := vote.GetImage()
	startAt, err := time.Parse(time.RFC3339, vote.GetStartAt())
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid startAt format")
	}
	endAt, err := time.Parse(time.RFC3339, vote.GetEndAt())
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid endAt format")
	}
	if startAt.After(endAt) {
		return nil, status.Error(codes.InvalidArgument, "startAt must be before endAt")
	}

	choicesList := make([]*model.Choice, len(choices))
	for i, choice := range choices {
		choiceNumber := choice.GetNumber()
		choiceInformation := choice.GetInformation()
		choiceImage := choice.GetImage()

		choicesList[i] = &model.Choice{
			Number:      &choiceNumber,
			Name:        choice.GetName(),
			Description: choice.GetDescription(),
			Information: &choiceInformation,
			Image:       &choiceImage,
		}
	}

	_, err = s.svc.CreateVote(ctx, oidcId, &model.Vote{
		Name:               vote.GetName(),
		Description:        vote.GetDescription(),
		Slug:               vote.GetSlug(),
		Image:              &voteImage,
		Owner:              model.Owner(vote.GetOwner().String()),
		EligibleStudentId:  vote.GetEligibleStudentId(),
		EligibleDepartment: vote.GetEligibleDepartment(),
		EligibleYear:       vote.GetEligibleYear(),
		IsPrivate:          vote.GetIsPrivate(),
		IsRealTime:         vote.GetIsRealTime(),
		IsAllowMultiple:    vote.GetIsAllowMultiple(),
		StartAt:            startAt,
		EndAt:              endAt,
	}, choicesList)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to create vote")
	}

	return &grpcVote.CreateVoteResponse{}, nil
}
