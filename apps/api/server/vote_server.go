package server

import (
	"context"
	"time"

	"github.com/esc-chula/intania-vote/apps/api/model"
	"github.com/esc-chula/intania-vote/apps/api/service"
	grpcChoice "github.com/esc-chula/intania-vote/libs/grpc-go/choice"
	grpcVote "github.com/esc-chula/intania-vote/libs/grpc-go/vote"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type voteServerImpl struct {
	grpcVote.UnimplementedVoteServiceServer
	svc       service.VoteService
	ballotSvc service.BallotService
}

func NewVoteServer(s service.VoteService, ballotS service.BallotService) grpcVote.VoteServiceServer {
	return voteServerImpl{
		svc:       s,
		ballotSvc: ballotS,
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
			Number:      choiceNumber,
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

func (s voteServerImpl) GetVoteById(ctx context.Context, req *grpcVote.GetVoteByIdRequest) (*grpcVote.GetVoteByIdResponse, error) {
	id := req.GetId()
	if id == 0 {
		return nil, status.Error(codes.FailedPrecondition, "missing id")
	}

	vote, err := s.svc.GetVoteById(ctx, uint(id))
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to get vote")
	}
	if vote == nil {
		return nil, status.Error(codes.NotFound, "vote not found")
	}

	choices := make([]*grpcChoice.Choice, len(vote.Choices))
	for i, choice := range vote.Choices {
		var ballotCounter *uint32
		if vote.EndAt.Before(time.Now()) {
			ballotCounter = &choice.BallotCounter
		} else {
			ballotCounter = nil
		}

		choices[i] = &grpcChoice.Choice{
			Number:        choice.Number,
			Name:          choice.Name,
			Description:   choice.Description,
			Information:   choice.Information,
			Image:         choice.Image,
			BallotCounter: ballotCounter,
		}
	}

	return &grpcVote.GetVoteByIdResponse{
		Vote: &grpcVote.Vote{
			Name:               vote.Name,
			Description:        vote.Description,
			Image:              vote.Image,
			Slug:               vote.Slug,
			Owner:              grpcVote.Owner(grpcVote.Owner_value[string(vote.Owner)]),
			EligibleStudentId:  vote.EligibleStudentId,
			EligibleDepartment: vote.EligibleDepartment,
			EligibleYear:       vote.EligibleYear,
			IsPrivate:          vote.IsPrivate,
			IsRealTime:         vote.IsRealTime,
			IsAllowMultiple:    vote.IsAllowMultiple,
			StartAt:            vote.StartAt.Format(time.RFC3339),
			EndAt:              vote.EndAt.Format(time.RFC3339),
		},
		Choices: choices,
	}, nil
}

func (s voteServerImpl) GetVoteBySlug(ctx context.Context, req *grpcVote.GetVoteBySlugRequest) (*grpcVote.GetVoteBySlugResponse, error) {
	slug := req.GetSlug()
	if slug == "" {
		return nil, status.Error(codes.FailedPrecondition, "missing slug")
	}

	vote, err := s.svc.GetVoteBySlug(ctx, slug)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to get vote")
	}
	if vote == nil {
		return nil, status.Error(codes.NotFound, "vote not found")
	}

	choices := make([]*grpcChoice.Choice, len(vote.Choices))
	for i, choice := range vote.Choices {
		var ballotCounter *uint32
		if vote.EndAt.Before(time.Now()) {
			ballotCounter = &choice.BallotCounter
		} else {
			ballotCounter = nil
		}

		choices[i] = &grpcChoice.Choice{
			Number:        choice.Number,
			Name:          choice.Name,
			Description:   choice.Description,
			Information:   choice.Information,
			Image:         choice.Image,
			BallotCounter: ballotCounter,
		}
	}

	return &grpcVote.GetVoteBySlugResponse{
		Vote: &grpcVote.Vote{
			Name:               vote.Name,
			Description:        vote.Description,
			Image:              vote.Image,
			Slug:               vote.Slug,
			Owner:              grpcVote.Owner(grpcVote.Owner_value[string(vote.Owner)]),
			EligibleStudentId:  vote.EligibleStudentId,
			EligibleDepartment: vote.EligibleDepartment,
			EligibleYear:       vote.EligibleYear,
			IsPrivate:          vote.IsPrivate,
			IsRealTime:         vote.IsRealTime,
			IsAllowMultiple:    vote.IsAllowMultiple,
			StartAt:            vote.StartAt.Format(time.RFC3339),
			EndAt:              vote.EndAt.Format(time.RFC3339),
		},
		Choices: choices,
	}, nil
}

func (s voteServerImpl) GetVotes(ctx context.Context, req *grpcVote.GetVotesRequest) (*grpcVote.GetVotesResponse, error) {
	votes, err := s.svc.GetVotes(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to get votes")
	}
	if len(votes) == 0 {
		return &grpcVote.GetVotesResponse{}, nil
	}

	voteList := make([]*grpcVote.Votes, len(votes))
	for i, vote := range votes {

		choices := make([]*grpcChoice.Choice, len(vote.Choices))
		for j, choice := range vote.Choices {
			var ballotCounter *uint32
			if vote.EndAt.Before(time.Now()) {
				ballotCounter = &choice.BallotCounter
			} else {
				ballotCounter = nil
			}

			choices[j] = &grpcChoice.Choice{
				Number:        choice.Number,
				Name:          choice.Name,
				Description:   choice.Description,
				Information:   choice.Information,
				Image:         choice.Image,
				BallotCounter: ballotCounter,
			}
		}

		ballotsCount, err := s.ballotSvc.CountBallotsByVoteId(ctx, vote.Id)
		if err != nil {
			return nil, status.Error(codes.Internal, "failed to get ballots count")
		}

		voteList[i] = &grpcVote.Votes{
			Vote: &grpcVote.Vote{
				Name:               vote.Name,
				Description:        vote.Description,
				Image:              vote.Image,
				Slug:               vote.Slug,
				Owner:              grpcVote.Owner(grpcVote.Owner_value[string(vote.Owner)]),
				EligibleStudentId:  vote.EligibleStudentId,
				EligibleDepartment: vote.EligibleDepartment,
				EligibleYear:       vote.EligibleYear,
				IsPrivate:          vote.IsPrivate,
				IsRealTime:         vote.IsRealTime,
				IsAllowMultiple:    vote.IsAllowMultiple,
				StartAt:            vote.StartAt.Format(time.RFC3339),
				EndAt:              vote.EndAt.Format(time.RFC3339),
			},
			Choices:      choices,
			TotalBallots: uint32(ballotsCount),
		}
	}

	return &grpcVote.GetVotesResponse{
		Votes: voteList,
	}, nil
}

func (s voteServerImpl) GetVotesByUserEligibility(ctx context.Context, req *grpcVote.GetVotesByUserEligibilityRequest) (*grpcVote.GetVotesByUserEligibilityResponse, error) {
	oidcId := req.GetOidcId()
	if oidcId == "" {
		return nil, status.Error(codes.FailedPrecondition, "missing oidcId")
	}

	votes, err := s.svc.GetVotesByUserEligibility(ctx, oidcId)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to get votes")
	}
	if len(votes) == 0 {
		return &grpcVote.GetVotesByUserEligibilityResponse{}, nil
	}

	userBallots, err := s.ballotSvc.GetBallotsByOidcId(ctx, oidcId)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to get ballots")
	}

	voteList := []*grpcVote.Votes{}
	for _, vote := range votes {
		var isVoted bool
		for _, ballot := range userBallots {
			if ballot.VoteId == vote.Id {
				isVoted = true
				break
			}
		}
		if isVoted {
			continue
		}

		choices := make([]*grpcChoice.Choice, len(vote.Choices))
		for j, choice := range vote.Choices {
			choices[j] = &grpcChoice.Choice{
				Number:        choice.Number,
				Name:          choice.Name,
				Description:   choice.Description,
				Information:   choice.Information,
				Image:         choice.Image,
				BallotCounter: nil,
			}
		}

		// append vote to list
		voteList = append(voteList, &grpcVote.Votes{
			Vote: &grpcVote.Vote{
				Name:               vote.Name,
				Description:        vote.Description,
				Image:              vote.Image,
				Slug:               vote.Slug,
				Owner:              grpcVote.Owner(grpcVote.Owner_value[string(vote.Owner)]),
				EligibleStudentId:  vote.EligibleStudentId,
				EligibleDepartment: vote.EligibleDepartment,
				EligibleYear:       vote.EligibleYear,
				IsPrivate:          vote.IsPrivate,
				IsRealTime:         vote.IsRealTime,
				IsAllowMultiple:    vote.IsAllowMultiple,
				StartAt:            vote.StartAt.Format(time.RFC3339),
				EndAt:              vote.EndAt.Format(time.RFC3339),
			},
			Choices: choices,
		})
	}

	return &grpcVote.GetVotesByUserEligibilityResponse{
		Votes: voteList,
	}, nil
}

func (s voteServerImpl) HasUserVoted(ctx context.Context, req *grpcVote.HasUserVotedRequest) (*grpcVote.HasUserVotedResponse, error) {
	oidcId := req.GetOidcId()
	if oidcId == "" {
		return nil, status.Error(codes.FailedPrecondition, "missing oidcId")
	}
	slug := req.GetSlug()
	if slug == "" {
		return nil, status.Error(codes.FailedPrecondition, "missing slug")
	}

	vote, err := s.svc.GetVoteBySlug(ctx, slug)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to get vote")
	}
	if vote == nil {
		return nil, status.Error(codes.NotFound, "vote not found")
	}

	hasVoted, err := s.ballotSvc.HasUserVoted(ctx, oidcId, slug)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to get ballots")
	}

	if hasVoted {
		return &grpcVote.HasUserVotedResponse{
			HasVoted: true,
		}, nil
	}

	return &grpcVote.HasUserVotedResponse{
		HasVoted: false,
	}, nil
}
