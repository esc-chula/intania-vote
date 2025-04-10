package server

import (
	"context"

	"github.com/esc-chula/intania-vote/apps/api/service"
	"github.com/esc-chula/intania-vote/apps/api/zk"
	grpcBallot "github.com/esc-chula/intania-vote/libs/grpc-go/ballot"
	grpcChoice "github.com/esc-chula/intania-vote/libs/grpc-go/choice"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ballotServerImpl struct {
	grpcBallot.UnimplementedBallotServiceServer
	svc service.BallotService
}

func NewBallotServer(s service.BallotService) grpcBallot.BallotServiceServer {
	return ballotServerImpl{
		svc: s,
	}
}

func RegisterBallotServiceServer(s grpc.ServiceRegistrar, server grpcBallot.BallotServiceServer) {
	grpcBallot.RegisterBallotServiceServer(s, server)
}

func (s ballotServerImpl) CreateBallotProof(ctx context.Context, req *grpcBallot.CreateBallotProofRequest) (*grpcBallot.CreateBallotProofResponse, error) {
	oidcId := req.GetOidcId()
	if oidcId == "" {
		return nil, status.Error(codes.FailedPrecondition, "missing oidcId")
	}
	choiceId := req.GetChoiceId()
	if choiceId == 0 {
		return nil, status.Error(codes.FailedPrecondition, "missing choiceId")
	}

	proofData, err := s.svc.CreateBallotProof(ctx, oidcId, uint(choiceId))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create proof: %v", err)
	}

	return &grpcBallot.CreateBallotProofResponse{
		Proof: &grpcBallot.Proof{
			Commitment:     proofData.Commitment,
			BlindingFactor: proofData.BlindingFactor,
			Challenge:      proofData.Challenge,
			Response:       proofData.Response,
			Nullifier:      proofData.Nullifier,
			Receipt:        proofData.Receipt,
		},
	}, nil
}

func (s ballotServerImpl) CreateBallot(ctx context.Context, req *grpcBallot.CreateBallotRequest) (*grpcBallot.CreateBallotResponse, error) {
	oidcId := req.GetOidcId()
	if oidcId == "" {
		return nil, status.Error(codes.FailedPrecondition, "missing oidcId")
	}
	voteSlug := req.GetVoteSlug()
	if voteSlug == "" {
		return nil, status.Error(codes.FailedPrecondition, "missing voteSlug")
	}
	proof := req.GetProof()
	if proof == nil {
		return nil, status.Error(codes.FailedPrecondition, "missing proof")
	}

	_, receipt, err := s.svc.CreateBallot(ctx, oidcId, voteSlug, zk.Proof{
		Commitment:     proof.GetCommitment(),
		BlindingFactor: proof.GetBlindingFactor(),
		Challenge:      proof.GetChallenge(),
		Response:       proof.GetResponse(),
		Nullifier:      proof.GetNullifier(),
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create ballot: %v", err)
	}

	return &grpcBallot.CreateBallotResponse{
		BallotKey: *receipt,
	}, nil
}

func (s ballotServerImpl) VerifyBallot(ctx context.Context, req *grpcBallot.VerifyBallotRequest) (*grpcBallot.VerifyBallotResponse, error) {
	oidcId := req.GetOidcId()
	if oidcId == "" {
		return nil, status.Error(codes.FailedPrecondition, "missing oidcId")
	}
	ballotKey := req.GetBallotKey()
	if ballotKey == "" {
		return nil, status.Error(codes.FailedPrecondition, "missing voteSlug")
	}

	choice, timestamp, err := s.svc.VerifyBallot(ctx, oidcId, ballotKey)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to verify ballot: %v", err)
	}

	return &grpcBallot.VerifyBallotResponse{
		IsValid: true,
		Choice: &grpcChoice.Choice{
			Number:      choice.Number,
			Name:        choice.Name,
			Description: choice.Description,
			Information: choice.Information,
			Image:       choice.Image,
		},
		Timestamp: timestamp.String(),
	}, nil
}
