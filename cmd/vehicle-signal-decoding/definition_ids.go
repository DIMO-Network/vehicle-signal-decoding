package main

import (
	"context"
	"github.com/DIMO-Network/shared/db"
	"github.com/DIMO-Network/vehicle-signal-decoding/internal/config"
	"github.com/DIMO-Network/vehicle-signal-decoding/internal/core/services"
	"github.com/DIMO-Network/vehicle-signal-decoding/internal/infrastructure/db/models"
	"github.com/rs/zerolog"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func SyncDefinitionIDs(ctx context.Context, logger zerolog.Logger, config *config.Settings, _ []string) {
	sqlDb := db.NewDbConnectionFromSettings(ctx, &config.DB, true)

	sqlDb.WaitForDB(logger)
	definitionsConn, err := grpc.NewClient(config.DefinitionsGRPCAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.Fatal().Err(err).Msg("Failed to create definitions connection")
	}
	deviceDefsvc := services.NewDeviceDefinitionsService(definitionsConn)

	all, err := models.TemplateDeviceDefinitions().All(ctx, sqlDb.DBS().Writer)
	if err != nil {
		logger.Fatal().Err(err).Msg("Failed to get all template device definitions")
	}
	for _, templateToDD := range all {
		if templateToDD.DefinitionID.IsZero() {
			dd, err := deviceDefsvc.GetDeviceDefinitionByID(ctx, templateToDD.DeviceDefinitionID)
			if err != nil {
				logger.Err(err).Msg("Failed to get device definition")
			}
			templateToDD.DefinitionID = null.StringFrom(dd.NameSlug)
			_, err = templateToDD.Update(ctx, sqlDb.DBS().Writer, boil.Infer())
			if err != nil {
				logger.Err(err).Msg("Failed to update vehicle to template")
			}
		}
	}
}
