package auction

import (
	"context"
	"fmt"

	"os"
	"time"

	"github.com/rafaelcamelo31/graduate-go-course/projects/auction/configuration/logger"
	"github.com/rafaelcamelo31/graduate-go-course/projects/auction/internal/entity/auction_entity"
	"github.com/rafaelcamelo31/graduate-go-course/projects/auction/internal/internal_error"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

type AuctionEntityMongo struct {
	Id          string                          `bson:"_id"`
	ProductName string                          `bson:"product_name"`
	Category    string                          `bson:"category"`
	Description string                          `bson:"description"`
	Condition   auction_entity.ProductCondition `bson:"condition"`
	Status      auction_entity.AuctionStatus    `bson:"status"`
	Timestamp   int64                           `bson:"timestamp"`
}

type AuctionRepository struct {
	Collection *mongo.Collection
}

func NewAuctionRepository(database *mongo.Database) *AuctionRepository {
	return &AuctionRepository{
		Collection: database.Collection("auctions"),
	}
}

func (ar *AuctionRepository) CreateAuction(
	ctx context.Context,
	auctionEntity *auction_entity.Auction) *internal_error.InternalError {
	logger.Info("Starting auction creation", zap.String("auctionId", auctionEntity.Id))

	// Valida os campos obrigatórios
	if auctionEntity.ProductName == "" || auctionEntity.Category == "" {
		logger.Warn("Validation failed: missing product name or category")
		return internal_error.NewBadRequestError("Product name and category are required")
	}

	// Cria o documento MongoDB
	auctionEntityMongo := &AuctionEntityMongo{
		Id:          auctionEntity.Id,
		ProductName: auctionEntity.ProductName,
		Category:    auctionEntity.Category,
		Description: auctionEntity.Description,
		Condition:   auctionEntity.Condition,
		Status:      auctionEntity.Status,
		Timestamp:   auctionEntity.Timestamp.Unix(),
	}
	_, err := ar.Collection.InsertOne(ctx, auctionEntityMongo)
	if err != nil {
		logger.Error("Error trying to insert auction", err)
		return internal_error.NewInternalServerError("Error trying to insert auction")
	}

	logger.Info("Auction successfully created, starting close routine", zap.String("auctionId", auctionEntity.Id))

	// Inicia a goroutine para fechamento automático
	go CloseAuctionRoutine(ctx, auctionEntity.Timestamp.Add(getAuctionDuration()), auctionEntity.Id, ar)

	return nil
}

func (ar *AuctionRepository) CloseAuction(
	ctx context.Context,
	auctionId string) error {
	_, err := ar.Collection.UpdateOne(ctx, bson.M{"_id": auctionId}, bson.M{"$set": bson.M{"status": auction_entity.Completed}})
	if err != nil {
		logger.Error("Error trying to close auction", err)
		return fmt.Errorf("error closing auction with ID %s: %w", auctionId, err)
	}

	logger.Info("Auction successfully closed", zap.String("auctionId", auctionId))
	return nil
}

type Repository interface {
	CloseAuction(ctx context.Context, auctionId string) error
}

func CloseAuctionRoutine(ctx context.Context, closeTime time.Time, auctionId string, repository Repository) {
	logger.Info("Starting close auction routine", zap.String("auctionId", auctionId), zap.Time("closeTime", closeTime))

	select {
	case <-time.After(time.Until(closeTime)):
		err := repository.CloseAuction(ctx, auctionId)
		if err != nil {
			logger.Error("Error trying to close auction", err)
			return
		}
		logger.Info("Auction closed successfully", zap.String("auctionId", auctionId))
	case <-ctx.Done():
		logger.Warn("Context cancelled, auction not closed", zap.String("auctionId", auctionId))
	}
}

func getAuctionDuration() time.Duration {
	auctionDuration := os.Getenv("AUCTION_DURATION")
	duration, err := time.ParseDuration(auctionDuration)
	if err != nil {
		logger.Warn("Invalid AUCTION_DURATION, using default value (5m)", zap.Error(err))
		return time.Minute * 5
	}

	logger.Info("Auction duration successfully retrieved", zap.String("duration", auctionDuration))
	return duration
}
